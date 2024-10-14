package keeper_test

import (
	"testing"

	testutil "github.com/ExocoreNetwork/exocore/testutil/keeper"
	"github.com/stretchr/testify/assert"
)

func TestKeeper_MapHeightToChainVscID(t *testing.T) {
	keeper, ctx, _ := testutil.NewCoordinatorKeeper(t)

	chainID := "test-chain"
	vscID := uint64(1)
	height := uint64(100)

	keeper.MapHeightToChainVscID(ctx, chainID, vscID, height)

	// Verify the mapping was stored correctly
	storedHeight := keeper.GetHeightForChainVscID(ctx, chainID, vscID)
	assert.Equal(t, height, storedHeight)
}

func TestKeeper_GetHeightForChainVscID(t *testing.T) {
	keeper, ctx, _ := testutil.NewCoordinatorKeeper(t)

	chainID := "test-chain"
	vscID := uint64(1)
	height := uint64(100)

	// Test when the mapping doesn't exist
	nonExistentHeight := keeper.GetHeightForChainVscID(ctx, chainID, vscID)
	assert.Equal(t, uint64(0), nonExistentHeight)

	// Set a mapping
	keeper.MapHeightToChainVscID(ctx, chainID, vscID, height)

	// Test when the mapping exists
	existingHeight := keeper.GetHeightForChainVscID(ctx, chainID, vscID)
	assert.Equal(t, height, existingHeight)
}

func TestKeeper_SetVscIDForChain(t *testing.T) {
	keeper, ctx, _ := testutil.NewCoordinatorKeeper(t)

	chainID := "test-chain"
	vscID := uint64(1)

	keeper.SetVscIDForChain(ctx, chainID, vscID)

	// Verify the vscID was stored correctly
	storedVscID := keeper.GetVscIDForChain(ctx, chainID)
	assert.Equal(t, vscID, storedVscID)
}

func TestKeeper_GetVscIDForChain(t *testing.T) {
	keeper, ctx, _ := testutil.NewCoordinatorKeeper(t)

	chainID := "test-chain"
	vscID := uint64(1)

	// Test when the vscID doesn't exist
	nonExistentVscID := keeper.GetVscIDForChain(ctx, chainID)
	assert.Equal(t, uint64(0), nonExistentVscID)

	// Set a vscID
	keeper.SetVscIDForChain(ctx, chainID, vscID)

	// Test when the vscID exists
	existingVscID := keeper.GetVscIDForChain(ctx, chainID)
	assert.Equal(t, vscID, existingVscID)
}

func TestKeeper_IncrementVscIDForChain(t *testing.T) {
	keeper, ctx, _ := testutil.NewCoordinatorKeeper(t)

	chainID := "test-chain"

	// Test initial increment (should start from 0)
	newVscID := keeper.IncrementVscIDForChain(ctx, chainID)
	assert.Equal(t, uint64(1), newVscID)

	// Test subsequent increment
	newVscID = keeper.IncrementVscIDForChain(ctx, chainID)
	assert.Equal(t, uint64(2), newVscID)

	// Verify the stored vscID
	storedVscID := keeper.GetVscIDForChain(ctx, chainID)
	assert.Equal(t, uint64(2), storedVscID)
}
