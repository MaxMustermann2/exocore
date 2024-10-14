package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	testkeeper "github.com/ExocoreNetwork/exocore/testutil/keeper"
	testutiltx "github.com/ExocoreNetwork/exocore/testutil/tx"
	commontypes "github.com/ExocoreNetwork/exocore/x/appchain/common/types"
	keepermod "github.com/ExocoreNetwork/exocore/x/appchain/coordinator/keeper"
)

func TestOperatorHooks(t *testing.T) {
	keeper, _, _ := testkeeper.NewCoordinatorKeeper(t)
	hooks := keeper.OperatorHooks()

	require.NotNil(t, hooks)
	require.IsType(t, keepermod.OperatorHooksWrapper{}, hooks)
}

func TestAfterOperatorKeySet(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)
	hooks := keeper.OperatorHooks()

	// AfterOperatorKeySet should do nothing, so we just ensure it doesn't panic
	require.NotPanics(t, func() {
		hooks.AfterOperatorKeySet(ctx, sdk.AccAddress{}, "test-chain", testutiltx.GenerateConsensusKey())
	})
}

func TestAfterOperatorKeyReplaced(t *testing.T) {
	keeper, ctx, mocks := testkeeper.NewCoordinatorKeeper(t)
	hooks := keeper.OperatorHooks()

	operator := sdk.AccAddress("testoperator")
	chainID := "test-chain"
	oldKey := testutiltx.GenerateConsensusKey()
	newKey := testutiltx.GenerateConsensusKey()

	t.Run("Existing validator", func(t *testing.T) {
		// Setup: Add a validator for the chain
		validator, err := commontypes.NewSubscriberValidator(oldKey.ToConsAddr(), 100, oldKey.ToSdkKey())
		require.NoError(t, err)
		keeper.SetSubscriberValidatorForChain(ctx, chainID, validator)

		// Set expectations
		mocks.OperatorKeeper.EXPECT().DeleteOperatorAddressForChainIDAndConsAddr(gomock.Any(), chainID, oldKey.ToConsAddr()).Times(0)

		// Call the hook
		hooks.AfterOperatorKeyReplaced(ctx, operator, oldKey, newKey, chainID)

		// Verify
		nextVscID := keeper.GetVscIDForChain(ctx, chainID) + 1
		consAddrsToPrune := keeper.GetConsAddrsToPrune(ctx, chainID, nextVscID)
		require.Contains(t, consAddrsToPrune.List, oldKey.ToConsAddr().Bytes())

		maturityVscID := keeper.GetMaturityVscIDForChainIDConsAddr(ctx, chainID, oldKey.ToConsAddr())
		require.Equal(t, nextVscID, maturityVscID)
	})

	t.Run("Non-existing validator", func(t *testing.T) {
		nonExistentKey := testutiltx.GenerateConsensusKey()

		// Set expectations
		mocks.OperatorKeeper.EXPECT().DeleteOperatorAddressForChainIDAndConsAddr(gomock.Any(), chainID, nonExistentKey.ToConsAddr()).Times(1)

		// Call the hook
		hooks.AfterOperatorKeyReplaced(ctx, operator, nonExistentKey, newKey, chainID)

		// Verify
		nextVscID := keeper.GetVscIDForChain(ctx, chainID) + 1
		consAddrsToPrune := keeper.GetConsAddrsToPrune(ctx, chainID, nextVscID)
		require.NotContains(t, consAddrsToPrune.List, nonExistentKey.ToConsAddr())
	})
}

func TestAfterOperatorKeyRemovalInitiated(t *testing.T) {
	keeper, ctx, mocks := testkeeper.NewCoordinatorKeeper(t)
	hooks := keeper.OperatorHooks()

	operator := sdk.AccAddress("testoperator")
	chainID := "test-chain"
	key := testutiltx.GenerateConsensusKey()

	t.Run("Existing validator", func(t *testing.T) {
		// Setup: Add a validator for the chain
		validator, err := commontypes.NewSubscriberValidator(key.ToConsAddr(), 100, key.ToSdkKey())
		require.NoError(t, err)
		keeper.SetSubscriberValidatorForChain(ctx, chainID, validator)

		// Set expectations
		mocks.OperatorKeeper.EXPECT().DeleteOperatorAddressForChainIDAndConsAddr(gomock.Any(), chainID, key.ToConsAddr()).Times(0)

		// Call the hook
		hooks.AfterOperatorKeyRemovalInitiated(ctx, operator, chainID, key)

		// Verify
		nextVscID := keeper.GetVscIDForChain(ctx, chainID) + 1
		consAddrsToPrune := keeper.GetConsAddrsToPrune(ctx, chainID, nextVscID)
		require.Contains(t, consAddrsToPrune.List, key.ToConsAddr().Bytes())

		maturityVscID := keeper.GetMaturityVscIDForChainIDConsAddr(ctx, chainID, key.ToConsAddr())
		require.Equal(t, nextVscID, maturityVscID)
	})

	t.Run("Non-existing validator", func(t *testing.T) {
		nonExistentKey := testutiltx.GenerateConsensusKey()

		// Set expectations
		mocks.OperatorKeeper.EXPECT().DeleteOperatorAddressForChainIDAndConsAddr(gomock.Any(), chainID, nonExistentKey.ToConsAddr()).Times(1)

		// Call the hook
		hooks.AfterOperatorKeyRemovalInitiated(ctx, operator, chainID, nonExistentKey)

		// Verify
		nextVscID := keeper.GetVscIDForChain(ctx, chainID) + 1
		consAddrsToPrune := keeper.GetConsAddrsToPrune(ctx, chainID, nextVscID)
		require.NotContains(t, consAddrsToPrune.List, nonExistentKey.ToConsAddr())
	})
}
