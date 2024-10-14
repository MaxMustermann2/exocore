package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	testkeeper "github.com/ExocoreNetwork/exocore/testutil/keeper"
	keepermod "github.com/ExocoreNetwork/exocore/x/appchain/coordinator/keeper"
)

func TestEpochsHooks(t *testing.T) {
	keeper, _, _ := testkeeper.NewCoordinatorKeeper(t)
	hooks := keeper.EpochsHooks()

	require.NotNil(t, hooks)
	require.IsType(t, keepermod.EpochsHooksWrapper{}, hooks)
}

func TestBeforeEpochStart(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)
	hooks := keeper.EpochsHooks()

	// BeforeEpochStart should do nothing, so we just ensure it doesn't panic
	require.NotPanics(t, func() {
		hooks.BeforeEpochStart(ctx, "test-epoch", 1)
	})
}

func TestAfterEpochEnd(t *testing.T) {
	keeper, ctx, mocks := testkeeper.NewCoordinatorKeeper(t)
	hooks := keeper.EpochsHooks()

	identifier := "test-epoch"
	epoch := int64(1)

	// Setup expectations for mocked methods
	mocks.AVSKeeper.EXPECT().GetEpochEndChainIDs(gomock.Any(), gomock.Any(), gomock.Any())
	mocks.OperatorKeeper.EXPECT().GetChainIDsForOperator(gomock.Any(), gomock.Any()).AnyTimes()
	mocks.OperatorKeeper.EXPECT().GetOperatorConsKeyForChainID(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	mocks.OperatorKeeper.EXPECT().IsOperatorRemovingKeyFromChainID(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	mocks.DelegationKeeper.EXPECT().IncrementUndelegationHoldCount(gomock.Any(), gomock.Any()).AnyTimes()

	// Expect ActivateScheduledChains to be called
	mocks.AVSKeeper.EXPECT().IsAVSByChainID(gomock.Any(), gomock.Any()).AnyTimes()
	mocks.AVSKeeper.EXPECT().DeleteAVSInfo(gomock.Any(), gomock.Any()).AnyTimes()

	// Expect RemoveTimedoutSubscribers to be called
	mocks.ClientKeeper.EXPECT().GetClientState(gomock.Any(), gomock.Any()).AnyTimes()
	// mocks.ClientKeeper.EXPECT().GetLatestHeight(gomock.Any(), gomock.Any()).AnyTimes()

	// Expect QueueValidatorUpdatesForEpochID to be called
	mocks.OperatorKeeper.EXPECT().GetActiveOperatorsForChainID(gomock.Any(), gomock.Any()).AnyTimes()
	mocks.OperatorKeeper.EXPECT().GetVotePowerForChainID(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

	// Expect SendQueuedValidatorUpdates to be called
	mocks.ChannelKeeper.EXPECT().GetNextSequenceSend(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	mocks.ChannelKeeper.EXPECT().SendPacket(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

	// Call AfterEpochEnd
	require.NotPanics(t, func() {
		hooks.AfterEpochEnd(ctx, identifier, epoch)
	})

}
