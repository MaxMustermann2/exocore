package keeper

import (
	"fmt"

	exocoretypes "github.com/ExocoreNetwork/exocore/types"
	avstypes "github.com/ExocoreNetwork/exocore/x/avs/types"
	"github.com/ExocoreNetwork/exocore/x/dogfood/types"
	abci "github.com/cometbft/cometbft/abci/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(
	ctx sdk.Context,
	genState types.GenesisState,
) []abci.ValidatorUpdate {
	// the `params` validator is not super useful to validate state level information
	// so, it must be done here. by extension, the `InitGenesis` of the epochs module
	// should be called before that of this module.
	epochID := genState.Params.EpochIdentifier
	epochInfo, found := k.epochsKeeper.GetEpochInfo(ctx, epochID)
	if !found {
		// the panic is suitable here because it is being done at genesis, when the node
		// is not running. it means that the genesis file is malformed.
		panic(fmt.Sprintf("epoch info not found %s", epochID))
	}
	// the staking assets are validated during AVS registration so we skip it here
	k.SetParams(ctx, genState.Params)
	// create the AVS
	var avsAddr common.Address
	var err error
	// the avs module will remove the revision by itself
	if avsAddr, err = k.avsKeeper.RegisterAVSWithChainID(ctx, &avstypes.AVSRegisterOrDeregisterParams{
		AvsName:           ctx.ChainID(),
		AssetID:           genState.Params.AssetIDs,
		UnbondingPeriod:   uint64(genState.Params.EpochsUntilUnbonded),
		MinSelfDelegation: genState.Params.MinSelfDelegation.Uint64(),
		EpochIdentifier:   epochID,
		ChainID:           ctx.ChainID(),
	}); err != nil {
		panic(fmt.Errorf("could not create the dogfood AVS: %s", err))
	}
	avsAddrString := avsAddr.String()
	ctx.Logger().With(types.ModuleName).Info(fmt.Sprintf("created dogfood avs %s %s", avsAddrString, ctx.ChainID()))
	// create the validators
	out := make([]exocoretypes.WrappedConsKeyWithPower, 0, len(genState.ValSet))
	for _, val := range genState.ValSet {
		// we have already checked in gs.Validate() that wrappedKey is not nil
		wrappedKey := exocoretypes.NewWrappedConsKeyFromHex(val.PublicKey)
		// #nosec G703 // already validated
		operatorAddr, _ := sdk.AccAddressFromBech32(val.OperatorAccAddr)
		// OptIntoAVS checks that the operator exists and will error if it does not.
		if err := k.operatorKeeper.OptInWithConsKey(
			ctx, operatorAddr, avsAddrString, wrappedKey,
		); err != nil {
			panic(fmt.Errorf("failed to opt into avs: %s", err))
		}
		out = append(out, exocoretypes.WrappedConsKeyWithPower{
			Power: val.Power,
			Key:   wrappedKey,
		})
	}
	for i := range genState.OptOutExpiries {
		obj := genState.OptOutExpiries[i]
		epoch := obj.Epoch
		if epoch < epochInfo.CurrentEpoch {
			panic(fmt.Sprintf("epoch %d is in the past", epoch))
		}
		for _, addr := range obj.OperatorAccAddrs {
			// #nosec G703 // already validated
			operatorAddr, _ := sdk.AccAddressFromBech32(addr)
			k.AppendOptOutToFinish(ctx, epoch, operatorAddr)
			k.SetOperatorOptOutFinishEpoch(ctx, operatorAddr, epoch)
		}
	}
	for i := range genState.ConsensusAddrsToPrune {
		obj := genState.ConsensusAddrsToPrune[i]
		epoch := obj.Epoch
		if epoch < epochInfo.CurrentEpoch {
			panic(fmt.Sprintf("epoch %d is in the past", epoch))
		}
		for _, addr := range obj.ConsAddrs {
			// #nosec G703 // already validated
			accAddr, _ := sdk.ConsAddressFromBech32(addr)
			k.AppendConsensusAddrToPrune(ctx, epoch, accAddr)
		}
	}
	for i := range genState.UndelegationMaturities {
		obj := genState.UndelegationMaturities[i]
		epoch := obj.Epoch
		if epoch < epochInfo.CurrentEpoch {
			panic(fmt.Sprintf("epoch %d is in the past", epoch))
		}
		for _, recordKey := range obj.UndelegationRecordKeys {
			// #nosec G703 // already validated
			recordKeyBytes, _ := hexutil.Decode(recordKey)
			k.AppendUndelegationToMature(ctx, epoch, recordKeyBytes)
			k.SetUndelegationMaturityEpoch(ctx, recordKeyBytes, epoch)
		}
	}
	// ApplyValidatorChanges only gets changes and hence the vote power must be set here.
	k.SetLastTotalPower(ctx, genState.LastTotalPower)

	// ApplyValidatorChanges will sort it internally
	return k.ApplyValidatorChanges(
		ctx, out,
	)
}

// ExportGenesis returns the module's exported genesis
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetDogfoodParams(ctx)
	validators := []types.GenesisValidator{}
	k.IterateBondedValidatorsByPower(ctx, func(_ int64, val stakingtypes.ValidatorI) bool {
		// #nosec G703 // already validated
		pubKey, _ := val.ConsPubKey()
		// #nosec G703 // already validated
		convKey, _ := cryptocodec.ToTmPubKeyInterface(pubKey)
		validators = append(validators,
			types.GenesisValidator{
				PublicKey: hexutil.Encode(convKey.Bytes()),
				Power:     val.GetConsensusPower(sdk.DefaultPowerReduction),
			},
		)
		return false /* stop */
	})
	return types.NewGenesis(
		k.GetDogfoodParams(ctx),
		validators,
		k.GetAllOptOutsToFinish(ctx),
		k.GetAllConsAddrsToPrune(ctx),
		k.GetAllUndelegationsToMature(ctx),
		k.GetLastTotalPower(ctx),
	)
}
