package keeper

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"github.com/ExocoreNetwork/exocore/utils"
	commontypes "github.com/ExocoreNetwork/exocore/x/appchain/common/types"
	"github.com/ExocoreNetwork/exocore/x/appchain/coordinator/types"
	abci "github.com/cometbft/cometbft/abci/types"
	tmtypes "github.com/cometbft/cometbft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/v7/modules/core/23-commitment/types"
	ibctmtypes "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
)

// ActivateScheduledChains activates the scheduled chains for the given epoch.
func (k Keeper) ActivateScheduledChains(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	executable := k.GetPendingSubChains(ctx, epochIdentifier, uint64(epochNumber))
	for _, subscriber := range executable.List {
		cctx, writeFn, err := k.CreateClientForSubscriberInCachedCtx(ctx, subscriber)
		if err != nil {
			k.Logger(ctx).Error(
				"error creating client for subscriber",
				"chainID", subscriber.ChainID,
				"error", err,
			)
			_, addr := k.avsKeeper.IsAVSByChainID(ctx, subscriber.ChainID)
			if err := k.avsKeeper.DeleteAVSInfo(ctx, addr); err != nil {
				// should never happen
				k.Logger(ctx).Error(
					"subscriber AVS not deleted",
					"chainID", subscriber,
					"error", err,
				)
			}
			continue
		}
		// copy over the events from the cached ctx
		ctx.EventManager().EmitEvents(cctx.EventManager().Events())
		writeFn()
		k.Logger(ctx).Info(
			"subscriber chain started",
			"chainID", subscriber,
			// we start at the current block and do not allow scheduling. this is the same
			// as any other AVS.
			"spawn time", ctx.BlockTime().UTC(),
		)
	}
	// clear pending queue, including those that errored out.
	k.ClearPendingSubChains(ctx, epochIdentifier, uint64(epochNumber))
}

// CreateClientForSubscriberInCachedCtx is a wrapper function around CreateClientForSubscriber.
func (k Keeper) CreateClientForSubscriberInCachedCtx(
	ctx sdk.Context,
	req types.RegisterSubscriberChainRequest,
) (cctx sdk.Context, writeCache func(), err error) {
	cctx, writeCache = ctx.CacheContext()
	err = k.CreateClientForSubscriber(cctx, req)
	return
}

// CreateClientForSubscriber creates a new client for the subscriber chain.
func (k Keeper) CreateClientForSubscriber(
	ctx sdk.Context,
	req types.RegisterSubscriberChainRequest,
) error {
	chainID := req.ChainID
	k.Logger(ctx).Info("Creating client for subscriber chain", "chainID", chainID)
	subscriberParams := req.SubscriberParams
	// we always deploy a new client for the subscriber chain for our module
	// technically, the below can never happen but it is guarded in ICS-20 and therefore, here.
	if _, found := k.GetClientForChain(ctx, chainID); found { // IBC related
		// client already exists
		return types.ErrDuplicateSubChain.Wrapf("chainID: %s", chainID)
	}
	// create the client
	coordinatorParams := k.GetParams(ctx)
	clientState := coordinatorParams.TemplateClient
	clientState.ChainId = chainID
	// TODO(mm): Make this configurable for switchover use case
	clientState.LatestHeight = clienttypes.Height{
		RevisionNumber: clienttypes.ParseChainID(chainID), // IBC related
		RevisionHeight: 1,
	}
	subscriberUnbondingPeriod := subscriberParams.UnbondingPeriod
	trustPeriod, err := commontypes.CalculateTrustPeriod(
		// note that this is the client that will live on Exocore
		// and we use the counterparty's unbonding period as a base
		subscriberUnbondingPeriod, coordinatorParams.TrustingPeriodFraction,
	)
	if err != nil {
		return err
	}
	clientState.TrustingPeriod = trustPeriod
	clientState.UnbondingPeriod = subscriberUnbondingPeriod

	// create the consensus state for the subscriber
	// do remember that the state we are storing here is just for the subscriber module
	// on the app chain, and not any other module. so it can have balances outside of this.
	subscriberGenesis, validatorSetHash, err := k.MakeSubscriberGenesis(ctx, req)
	if err != nil {
		return err
	}
	// this state can be pruned after the initial handshake occurs.
	k.SetSubscriberGenesis(ctx, chainID, subscriberGenesis)
	consensusState := ibctmtypes.NewConsensusState(
		ctx.BlockTime(),
		commitmenttypes.NewMerkleRoot([]byte(ibctmtypes.SentinelRoot)),
		validatorSetHash, // use the hash of the updated initial valset
	)

	clientID, err := k.clientKeeper.CreateClient(ctx, clientState, consensusState)
	if err != nil {
		return err
	}
	k.SetClientForChain(ctx, chainID, clientID)

	epochInfo, _ := k.epochsKeeper.GetEpochInfo(ctx, req.EpochIdentifier)
	initTimeoutPeriod := coordinatorParams.InitTimeoutPeriod
	// the CurrentEpoch below is the one that has already ended, so we increment it by 1.
	// assume we start with a value of 2 and we are giving 4 full epochs for initialization.
	// so when epoch 6 ends, the timeout ends.
	initTimeoutPeriod.EpochNumber += uint64(epochInfo.CurrentEpoch) + 1
	// lookup from timeout to chainID
	k.AppendChainToInitTimeout(ctx, initTimeoutPeriod, chainID)
	// reverse lookup from chainID to timeout
	k.SetChainInitTimeout(ctx, chainID, initTimeoutPeriod)

	// if the chain doesn't initialize in time, the following items will be cleared.
	// same for timeout or error.
	k.SetSubSlashFractionDowntime(ctx, chainID, subscriberParams.SlashFractionDowntime)
	k.SetSubSlashFractionDoubleSign(ctx, chainID, subscriberParams.SlashFractionDoubleSign)
	k.SetSubDowntimeJailDuration(ctx, chainID, subscriberParams.DowntimeJailDuration)
	k.SetMaxValidatorsForChain(ctx, chainID, req.MaxValidators)

	k.Logger(ctx).Info(
		"subscriber chain registered (client created)",
		"chainId", chainID,
		"clientId", clientID,
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSubscriberClientCreated,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(commontypes.AttributeChainID, chainID),
			sdk.NewAttribute(clienttypes.AttributeKeyClientID, clientID),
			sdk.NewAttribute(types.AttributeInitialHeight, clientState.LatestHeight.String()),
			sdk.NewAttribute(
				types.AttributeInitializationTimeout,
				initTimeoutPeriod.String(),
			),
			sdk.NewAttribute(
				types.AttributeTrustingPeriod,
				clientState.TrustingPeriod.String(),
			),
			sdk.NewAttribute(
				types.AttributeUnbondingPeriod,
				clientState.UnbondingPeriod.String(),
			),
		),
	)

	return nil
}

func (k Keeper) MakeSubscriberGenesis(
	ctx sdk.Context, req types.RegisterSubscriberChainRequest,
) (*commontypes.SubscriberGenesisState, []byte, error) {
	params := k.GetParams(ctx)
	chainID := req.ChainID
	k.Logger(ctx).Info("Creating genesis state for subscriber chain", "chainID", chainID)
	coordinatorUnbondingPeriod := k.stakingKeeper.UnbondingTime(ctx)
	// client state
	clientState := params.TemplateClient
	clientState.ChainId = chainID
	clientState.LatestHeight = clienttypes.GetSelfHeight(ctx)
	trustPeriod, err := commontypes.CalculateTrustPeriod(
		coordinatorUnbondingPeriod, params.TrustingPeriodFraction,
	)
	if err != nil {
		return nil, nil, errorsmod.Wrapf(
			sdkerrors.ErrInvalidHeight,
			"error %s calculating self trusting period for chain %s",
			err, chainID,
		)
	}
	clientState.TrustingPeriod = trustPeriod
	clientState.UnbondingPeriod = coordinatorUnbondingPeriod
	// consensus state
	consState, err := k.clientKeeper.GetSelfConsensusState(ctx, clientState.LatestHeight)
	if err != nil {
		return nil, nil, errorsmod.Wrapf(
			clienttypes.ErrConsensusStateNotFound,
			"error %s getting self consensus state for chain %s",
			err, chainID,
		)
	}
	operators, keys := k.operatorKeeper.GetActiveOperatorsForChainID(ctx, chainID)
	powers, err := k.operatorKeeper.GetVotePowerForChainID(
		ctx, operators, chainID,
	)
	if err != nil {
		k.Logger(ctx).Error("error getting vote power for chain", "error", err)
		return nil, nil, err
	}
	operators, keys, powers = utils.SortByPower(operators, keys, powers)
	maxVals := req.MaxValidators
	validatorUpdates := make([]abci.ValidatorUpdate, 0, maxVals)
	for i := range operators {
		if i >= int(maxVals) {
			break
		}
		power := powers[i]
		if power < 1 {
			break
		}
		wrappedKey := keys[i]
		validator, err := commontypes.NewSubscriberValidator(
			wrappedKey.ToConsAddr(), power, wrappedKey.ToSdkKey(),
		)
		if err != nil {
			// cannot happen, but just in case add this check.
			// simply skip the validator if it does.
			continue
		}
		validatorUpdates = append(validatorUpdates, abci.ValidatorUpdate{
			PubKey: *wrappedKey.ToTmProtoKey(),
			Power:  power,
		})
		// at the time of genesis, all validators are stored since there is no
		// existing validator set to take a diff with.
		k.SetSubscriberValidatorForChain(ctx, chainID, validator)
	}
	if len(validatorUpdates) == 0 {
		return nil, nil, errorsmod.Wrapf(
			types.ErrNoOperators, "no operators with stake found for chainID: %s", chainID,
		)
	}
	// sorted by power (with accAddr as tie breaker) and hence deterministic
	updatesAsValSet, err := tmtypes.PB2TM.ValidatorUpdates(validatorUpdates)
	if err != nil {
		return nil, nil, fmt.Errorf(
			"could not create validator set from validator updates: %s", err,
		)
	}
	return &commontypes.SubscriberGenesisState{
		Params: req.SubscriberParams,
		Coordinator: commontypes.CoordinatorInfo{
			ClientState:    clientState,
			ConsensusState: consState.(*ibctmtypes.ConsensusState),
			InitialValSet:  validatorUpdates,
		},
	}, tmtypes.NewValidatorSet(updatesAsValSet).Hash(), nil
}

// SetSubscriberGenesis sets the genesis state for the subscriber chain.
func (k Keeper) SetSubscriberGenesis(
	ctx sdk.Context,
	chainID string,
	genesis *commontypes.SubscriberGenesisState,
) {
	store := ctx.KVStore(k.storeKey)
	key := types.SubscriberGenesisKey(chainID)
	store.Set(key, k.cdc.MustMarshal(genesis))
}

// GetSubscriberGenesis gets the genesis state for the subscriber chain.
func (k Keeper) GetSubscriberGenesis(
	ctx sdk.Context,
	chainID string,
) (genesis commontypes.SubscriberGenesisState) {
	store := ctx.KVStore(k.storeKey)
	key := types.SubscriberGenesisKey(chainID)
	k.cdc.MustUnmarshal(store.Get(key), &genesis)
	return genesis
}

// DeleteSubscriberGenesis deletes the genesis state for the subscriber chain.
// It is a pruning function.
func (k Keeper) DeleteSubscriberGenesis(
	ctx sdk.Context,
	chainID string,
) {
	store := ctx.KVStore(k.storeKey)
	key := types.SubscriberGenesisKey(chainID)
	store.Delete(key)
}
