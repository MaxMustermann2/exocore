package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ExocoreNetwork/exocore/testutil/keeper"
	"github.com/ExocoreNetwork/exocore/x/appchain/coordinator/types"
)

func TestAppendGetConsAddrsToPrune(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)

	chainID := "test-chain"
	vscID := uint64(1)
	consAddr1 := sdk.ConsAddress([]byte("consAddr1"))
	consAddr2 := sdk.ConsAddress([]byte("consAddr2"))

	// Test appending and getting
	keeper.AppendConsAddrToPrune(ctx, chainID, vscID, consAddr1)
	keeper.AppendConsAddrToPrune(ctx, chainID, vscID, consAddr2)

	addrsToPrune := keeper.GetConsAddrsToPrune(ctx, chainID, vscID)
	require.Equal(t, 2, len(addrsToPrune.List))
	require.Contains(t, addrsToPrune.List, consAddr1.Bytes())
	require.Contains(t, addrsToPrune.List, consAddr2.Bytes())
}

func TestSetGetConsAddrsToPrune(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)

	chainID := "test-chain"
	vscID := uint64(1)
	consAddrs := types.ConsensusAddresses{
		List: [][]byte{
			sdk.ConsAddress([]byte("consAddr1")).Bytes(),
			sdk.ConsAddress([]byte("consAddr2")).Bytes(),
		},
	}

	// Test setting and getting
	keeper.SetConsAddrsToPrune(ctx, chainID, vscID, consAddrs)
	gotAddrs := keeper.GetConsAddrsToPrune(ctx, chainID, vscID)
	require.Equal(t, consAddrs, gotAddrs)

	// Test deleting when empty
	keeper.SetConsAddrsToPrune(ctx, chainID, vscID, types.ConsensusAddresses{})
	gotAddrs = keeper.GetConsAddrsToPrune(ctx, chainID, vscID)
	require.Empty(t, gotAddrs.List)
}

func TestSetGetDeleteMaturityVscIDForChainIDConsAddr(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)

	chainID := "test-chain"
	consAddr := sdk.ConsAddress([]byte("consAddr"))
	vscID := uint64(1)

	// Test setting and getting
	keeper.SetMaturityVscIDForChainIDConsAddr(ctx, chainID, consAddr, vscID)
	gotVscID := keeper.GetMaturityVscIDForChainIDConsAddr(ctx, chainID, consAddr)
	require.Equal(t, vscID, gotVscID)

	// Test deleting
	keeper.DeleteMaturityVscIDForChainIDConsAddr(ctx, chainID, consAddr)
	gotVscID = keeper.GetMaturityVscIDForChainIDConsAddr(ctx, chainID, consAddr)
	require.Equal(t, uint64(0), gotVscID)
}

func TestAppendGetUndelegationsToRelease(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)

	chainID := "test-chain"
	vscID := uint64(1)
	recordKey1 := []byte("recordKey1")
	recordKey2 := []byte("recordKey2")

	// Test appending and getting
	keeper.AppendUndelegationToRelease(ctx, chainID, vscID, recordKey1)
	keeper.AppendUndelegationToRelease(ctx, chainID, vscID, recordKey2)

	undelegationsToRelease := keeper.GetUndelegationsToRelease(ctx, chainID, vscID)
	require.Equal(t, 2, len(undelegationsToRelease.List))
	require.Contains(t, undelegationsToRelease.List, recordKey1)
	require.Contains(t, undelegationsToRelease.List, recordKey2)
}

func TestSetGetUndelegationsToRelease(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)

	chainID := "test-chain"
	vscID := uint64(1)
	undelegations := types.UndelegationRecordKeys{
		List: [][]byte{
			[]byte("recordKey1"),
			[]byte("recordKey2"),
		},
	}

	// Test setting and getting
	keeper.SetUndelegationsToRelease(ctx, chainID, vscID, undelegations)
	gotUndelegations := keeper.GetUndelegationsToRelease(ctx, chainID, vscID)
	require.Equal(t, undelegations, gotUndelegations)

	// Test deleting when empty
	keeper.SetUndelegationsToRelease(ctx, chainID, vscID, types.UndelegationRecordKeys{})
	gotUndelegations = keeper.GetUndelegationsToRelease(ctx, chainID, vscID)
	require.Empty(t, gotUndelegations.List)
}
