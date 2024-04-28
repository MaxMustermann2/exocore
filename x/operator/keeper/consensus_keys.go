package keeper

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"github.com/cometbft/cometbft/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	delegationtypes "github.com/ExocoreNetwork/exocore/x/delegation/types"
	"github.com/ExocoreNetwork/exocore/x/operator/types"

	tmprotocrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
)

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetOperatorConsKeyForChainID sets the (consensus) public key for the given operator address
// and chain id. By doing this, an operator is consenting to be an operator on the given chain.
// If a key already exists, it will be overwritten and the change in voting power will flow
// through to the validator set.
func (k *Keeper) SetOperatorConsKeyForChainID(
	ctx sdk.Context,
	opAccAddr sdk.AccAddress,
	chainID string,
	// should be tm-ed25519
	consKey *tmprotocrypto.PublicKey,
) error {
	return k.setOperatorConsKeyForChainID(ctx, opAccAddr, chainID, consKey, false)
}

// setOperatorConsKeyForChainID is the private version of SetOperatorConsKeyForChainID.
// it is used with a boolean flag to indicate that the call is from genesis.
// if so, operator freeze status is not checked and hooks are not called.
func (k *Keeper) setOperatorConsKeyForChainID(
	ctx sdk.Context,
	opAccAddr sdk.AccAddress,
	chainID string,
	consKey *tmprotocrypto.PublicKey,
	genesis bool,
) error {
	// check if the chain id is valid
	// TODO: when appchain AVS is implemented, move the check there.
	if chainID != ctx.ChainID() {
		return types.ErrUnknownChainID
	}
	// check if we are an operator
	if !k.IsOperator(ctx, opAccAddr) {
		return delegationtypes.ErrOperatorNotExist
	}
	// check for slashing
	if !genesis && k.slashKeeper.IsOperatorFrozen(ctx, opAccAddr) {
		return delegationtypes.ErrOperatorIsFrozen
	}
	// check if the operator is opted in. it is a 2-step process.
	// 1. opt into chainID
	// 2. set key for chainID.
	if !k.IsOptedIn(ctx, opAccAddr.String(), chainID) {
		return types.ErrNotOptedIn
	}
	// if removing key, do not allow key replacement
	if k.IsOperatorRemovingKeyFromChainID(ctx, opAccAddr, chainID) {
		return types.ErrAlreadyRemovingKey
	}
	// convert to bytes
	bz := k.cdc.MustMarshal(consKey)
	// convert to address for reverse lookup
	consAddr, err := types.TMCryptoPublicKeyToConsAddr(consKey)
	if err != nil {
		return errorsmod.Wrap(
			err, "SetOperatorConsKeyForChainID: cannot convert pub key to consensus address",
		)
	}
	// check if the key is already in use by another operator
	// even if it is the same operator, do not allow calling this function twice.
	keyInUse, _ := k.GetOperatorAddressForChainIDAndConsAddr(ctx, chainID, consAddr)
	if keyInUse {
		return types.ErrConsKeyAlreadyInUse
	}
	// check that such a key is already set. if yes, we will consider it as key replacement.
	found, prevKey := k.getOperatorConsKeyForChainID(ctx, opAccAddr, chainID)
	if err != nil {
		// this should not happen
		panic(err)
	}
	var alreadyRecorded bool
	if found {
		// ultimately performs bytes.Equal
		if prevKey.Equal(consKey) {
			// no-op
			return nil
		}
		// if this key is different, we will set the vote power of the old key to 0
		// in the validator update. but, we must only do so once in a block, since the
		// first existing key is the one to replace with 0 vote power and not any others.
		alreadyRecorded, _ = k.getOperatorPrevConsKeyForChainID(ctx, opAccAddr, chainID)
		if !alreadyRecorded {
			if err := k.setOperatorPrevConsKeyForChainID(
				ctx, opAccAddr, chainID, prevKey,
			); err != nil {
				// this should not happen
				panic(err)
			}
		}
	}
	k.setOperatorConsKeyForChainIDUnchecked(ctx, opAccAddr, consAddr, chainID, bz)
	if !genesis {
		if found {
			if !alreadyRecorded {
				k.Hooks().AfterOperatorKeyReplacement(ctx, opAccAddr, prevKey, consKey, chainID)
			}
		} else {
			k.Hooks().AfterOperatorOptIn(ctx, opAccAddr, chainID, consKey)
		}
	}
	return nil
}

// setOperatorConsKeyForChainIDUnchecked is the internal private version. It performs
// no error checking of the input. The caller must do the error checking
// and then call this function.
func (k Keeper) setOperatorConsKeyForChainIDUnchecked(
	ctx sdk.Context, opAccAddr sdk.AccAddress, consAddr sdk.ConsAddress,
	chainID string, bz []byte,
) {
	store := ctx.KVStore(k.storeKey)
	// forward lookup
	// given operator address and chain id, find the consensus key,
	// since it is sorted by operator address, it helps for faster indexing by operator
	// for example, when an operator is delegated to, we can find all impacted
	// chain ids and their respective consensus keys
	store.Set(types.KeyForOperatorAndChainIDToConsKey(opAccAddr, chainID), bz)
	// reverse lookups
	// 1. given chain id and operator address, find the consensus key,
	// at initial onboarding of an app chain, it will allow us to find all
	// operators that have opted in and their consensus keys
	store.Set(types.KeyForChainIDAndOperatorToConsKey(chainID, opAccAddr), bz)
	// 2. given a chain id and a consensus addr, find the operator address,
	// the slashing module asks for an operator to be slashed by their consensus
	// address, so this will allow us to find the operator address to slash.
	// however, we do not want to retain this information forever, so we will
	// prune it once the validator set update id matures (if key replacement).
	// this pruning will be triggered by the app chain module and will not be
	// recorded here.
	store.Set(types.KeyForChainIDAndConsKeyToOperator(chainID, consAddr), opAccAddr.Bytes())
}

// setOperatorPrevConsKeyForChainID sets the previous (consensus) public key for the given
// operator address and chain id. This is used to track the previous key when a key is replaced.
// It is internal-only because such a key must only be set upon key replacement. So it does
// not perform any meaningful error checking of the input.
func (k *Keeper) setOperatorPrevConsKeyForChainID(
	ctx sdk.Context,
	opAccAddr sdk.AccAddress,
	chainID string,
	prevKey *tmprotocrypto.PublicKey,
) error {
	bz, err := prevKey.Marshal()
	if err != nil {
		return errorsmod.Wrap(
			err,
			"SetOperatorPrevConsKeyForChainID: error occurred when marshal public key",
		)
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyForOperatorAndChainIDToPrevConsKey(opAccAddr, chainID), bz)
	return nil
}

// GetOperatorPrevConsKeyForChainID gets the previous (consensus) public key for the given
// operator address and chain id. When such a key is returned, callers should set its vote power
// to 0 in the validator update.
func (k *Keeper) GetOperatorPrevConsKeyForChainID(
	ctx sdk.Context, opAccAddr sdk.AccAddress, chainID string,
) (bool, *tmprotocrypto.PublicKey, error) {
	// check if we are an operator
	if !k.IsOperator(ctx, opAccAddr) {
		return false, nil, delegationtypes.ErrOperatorNotExist
	}
	// check if the chain exists as an AVS
	if chainID != ctx.ChainID() {
		return false, nil, types.ErrUnknownChainID
	}
	found, key := k.getOperatorPrevConsKeyForChainID(ctx, opAccAddr, chainID)
	return found, key, nil
}

// getOperatorPrevConsKeyForChainID is the internal version of GetOperatorPrevConsKeyForChainID.
// It performs no error checking of the input.
func (k *Keeper) getOperatorPrevConsKeyForChainID(
	ctx sdk.Context,
	opAccAddr sdk.AccAddress,
	chainID string,
) (bool, *tmprotocrypto.PublicKey) {
	store := ctx.KVStore(k.storeKey)
	res := store.Get(types.KeyForOperatorAndChainIDToPrevConsKey(opAccAddr, chainID))
	if res == nil {
		return false, nil
	}
	key := &tmprotocrypto.PublicKey{}
	k.cdc.MustUnmarshal(res, key)
	return true, key
}

// GetOperatorConsKeyForChainID gets the (consensus) public key for the given operator address
// and chain id. This should be exposed via the query surface.
func (k *Keeper) GetOperatorConsKeyForChainID(
	ctx sdk.Context,
	opAccAddr sdk.AccAddress,
	chainID string,
) (bool, *tmprotocrypto.PublicKey, error) {
	// check if we are an operator
	if !k.IsOperator(ctx, opAccAddr) {
		return false, nil, delegationtypes.ErrOperatorNotExist
	}
	if chainID != ctx.ChainID() {
		return false, nil, types.ErrUnknownChainID
	}
	found, key := k.getOperatorConsKeyForChainID(ctx, opAccAddr, chainID)
	return found, key, nil
}

// getOperatorConsKeyForChainID is the internal version of GetOperatorConsKeyForChainID. It
// performs no error checking of the input.
func (k *Keeper) getOperatorConsKeyForChainID(
	ctx sdk.Context,
	opAccAddr sdk.AccAddress,
	chainID string,
) (bool, *tmprotocrypto.PublicKey) {
	store := ctx.KVStore(k.storeKey)
	res := store.Get(types.KeyForOperatorAndChainIDToConsKey(opAccAddr, chainID))
	if res == nil {
		return false, nil
	}
	key := &tmprotocrypto.PublicKey{}
	k.cdc.MustUnmarshal(res, key)
	return true, key
}

// GetOperatorsForChainID returns a list of {operatorAddr, pubKey} for the given
// chainID. This is used to create or update the validator set. It skips
// jailed or frozen operators.
func (k *Keeper) GetOperatorsForChainID(
	ctx sdk.Context, chainID string,
) ([]sdk.AccAddress, []*tmprotocrypto.PublicKey) {
	if ctx.ChainID() != chainID {
		return nil, nil
	}
	// prefix is the byte prefix and then chainID with length
	prefix := types.ChainIDAndAddrKey(
		types.BytePrefixForChainIDAndOperatorToConsKey,
		chainID, nil,
	)
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(
		store, prefix,
	)
	defer iterator.Close()
	var addrs []sdk.AccAddress
	var pubKeys []*tmprotocrypto.PublicKey
	for ; iterator.Valid(); iterator.Next() {
		// this key is of the format prefix | len | chainID | addr
		// and our prefix is of the format prefix | len | chainID
		// so just drop it and convert to sdk.AccAddress
		addr := iterator.Key()[len(prefix):]
		res := iterator.Value()
		ret := &tmprotocrypto.PublicKey{}
		k.cdc.MustUnmarshal(res, ret)
		addrs = append(addrs, addr)
		pubKeys = append(pubKeys, ret)
	}
	return addrs, pubKeys
}

func (k *Keeper) GetOperatorAddressForChainIDAndConsAddr(
	ctx sdk.Context, chainID string, consAddr sdk.ConsAddress,
) (found bool, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	res := store.Get(types.KeyForChainIDAndConsKeyToOperator(chainID, consAddr))
	if res == nil {
		return
	}
	found = true
	addr = sdk.AccAddress(res)
	return found, addr
}

// DeleteOperatorAddressForChainIDAndConsAddr is a pruning method used to delete the
// mapping from chain id and consensus address to operator address. This mapping is used
// to obtain the operator address from its consensus public key, which is sent to the
// coordinator chain by a subscriber chain for slashing.
func (k *Keeper) DeleteOperatorAddressForChainIDAndConsAddr(
	ctx sdk.Context, chainID string, consAddr sdk.ConsAddress,
) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyForChainIDAndConsKeyToOperator(chainID, consAddr))
}

// SetHooks stores the given hooks implementations.
// Note that the Keeper is changed into a pointer to prevent an ineffective assignment.
func (k *Keeper) SetHooks(hooks types.OperatorHooks) {
	if hooks == nil {
		panic("cannot set nil hooks")
	}
	if k.hooks != nil {
		panic("cannot set hooks twice")
	}
	k.hooks = hooks
}

func (k *Keeper) Hooks() types.OperatorHooks {
	if k.hooks == nil {
		// return a no-op implementation if no hooks are set to prevent calling nil functions
		return types.MultiOperatorHooks{}
	}
	return k.hooks
}

// InitiateOperatorKeyRemovalForChainID initiates an operator removing their key from the
// chain id. It validates whether the operator is registered, and that it is not frozen, and
// that the chain is present within the system. It also checks if the operator is already
// removing the key.
// The operator must first call this method, and then call `OptOut` to complete the process.
func (k *Keeper) InitiateOperatorKeyRemovalForChainID(
	ctx sdk.Context, opAccAddr sdk.AccAddress, chainID string,
) error {
	// check if we are an operator
	if !k.IsOperator(ctx, opAccAddr) {
		return delegationtypes.ErrOperatorNotExist
	}
	// check for slashing
	if k.slashKeeper.IsOperatorFrozen(ctx, opAccAddr) {
		return delegationtypes.ErrOperatorIsFrozen
	}
	// check if the chain id is valid
	if chainID != ctx.ChainID() {
		return types.ErrUnknownChainID
	}
	found, key := k.getOperatorConsKeyForChainID(ctx, opAccAddr, chainID)
	if !found {
		return types.ErrNotOptedIn
	}
	if k.IsOperatorRemovingKeyFromChainID(ctx, opAccAddr, chainID) {
		return types.ErrAlreadyRemovingKey
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyForOperatorKeyRemovalForChainID(opAccAddr, chainID), []byte{})
	k.Hooks().AfterOperatorOptOutInitiated(ctx, opAccAddr, chainID, key)
	return nil
}

// IsOperatorRemovingKeyFromChainID returns true if the operator is removing the consensus
// key from the given chain id.
func (k *Keeper) IsOperatorRemovingKeyFromChainID(
	ctx sdk.Context, opAccAddr sdk.AccAddress, chainID string,
) bool {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyForOperatorKeyRemovalForChainID(opAccAddr, chainID))
	return bz != nil
}

// CompleteOperatorRemovingKeyFromChainID completes the operator key removal from the given
// chain id.
func (k *Keeper) CompleteKeyRemovalForChainID(
	ctx sdk.Context, opAccAddr sdk.AccAddress, chainID string,
) {
	if k.IsOperatorRemovingKeyFromChainID(ctx, opAccAddr, chainID) {
		store := ctx.KVStore(k.storeKey)
		store.Delete(types.KeyForOperatorKeyRemovalForChainID(opAccAddr, chainID))
	}
}

func (k *Keeper) GetActiveOperatorsForChainID(
	ctx sdk.Context, chainID string,
) ([]sdk.AccAddress, []*tmprotocrypto.PublicKey) {
	operatorsAddr, pks := k.GetOperatorsForChainID(ctx, chainID)
	avsAddr, err := k.avsKeeper.GetAvsAddrByChainID(ctx, chainID)
	if err != nil {
		k.Logger(ctx).Error(err.Error(), chainID)
		return nil, nil
	}

	activeOperator := make([]sdk.AccAddress, 0)
	activePks := make([]*tmprotocrypto.PublicKey, 0)
	// check if the operator is active
	for i, operator := range operatorsAddr {
		if k.IsActive(ctx, operator.String(), avsAddr) {
			activeOperator = append(activeOperator, operator)
			activePks = append(activePks, pks[i])
		}
	}
	return activeOperator, activePks
}

func (k *Keeper) ValidatorByConsAddrForChainID(
	ctx sdk.Context, consAddr sdk.ConsAddress, chainID string,
) stakingtypes.ValidatorI {
	found, operatorAddr := k.GetOperatorAddressForChainIDAndConsAddr(
		ctx, chainID, consAddr,
	)
	if !found {
		return stakingtypes.Validator{}
	}
	jailed := k.IsOperatorJailedForChainID(ctx, consAddr, chainID)
	return stakingtypes.Validator{
		Jailed:          jailed,
		OperatorAddress: sdk.ValAddress(operatorAddr).String(),
	}
}
