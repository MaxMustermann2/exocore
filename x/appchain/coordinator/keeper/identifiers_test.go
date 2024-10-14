package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ExocoreNetwork/exocore/testutil/keeper"
)

func TestSetGetClientForChain(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)

	chainID := "test-chain"
	clientID := "test-client"

	// Test setting and getting
	keeper.SetClientForChain(ctx, chainID, clientID)
	gotClientID, found := keeper.GetClientForChain(ctx, chainID)

	require.True(t, found)
	require.Equal(t, clientID, gotClientID)

	// Test getting non-existent client
	_, found = keeper.GetClientForChain(ctx, "non-existent-chain")
	require.False(t, found)
}

func TestDeleteClientForChain(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)

	chainID := "test-chain"
	clientID := "test-client"

	keeper.SetClientForChain(ctx, chainID, clientID)
	keeper.DeleteClientForChain(ctx, chainID)

	_, found := keeper.GetClientForChain(ctx, chainID)
	require.False(t, found)
}

func TestGetAllChainsWithClients(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)

	chainIDs := []string{"chain1", "chain2", "chain3"}
	for _, chainID := range chainIDs {
		keeper.SetClientForChain(ctx, chainID, "client-"+chainID)
	}

	gotChainIDs := keeper.GetAllChainsWithClients(ctx)
	require.ElementsMatch(t, chainIDs, gotChainIDs)
}

func TestSetGetChannelForChain(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)

	chainID := "test-chain"
	channelID := "test-channel"

	keeper.SetChannelForChain(ctx, chainID, channelID)
	gotChannelID, found := keeper.GetChannelForChain(ctx, chainID)

	require.True(t, found)
	require.Equal(t, channelID, gotChannelID)

	_, found = keeper.GetChannelForChain(ctx, "non-existent-chain")
	require.False(t, found)
}

func TestGetAllChainsWithChannels(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)

	chainIDs := []string{"chain1", "chain2", "chain3"}
	for _, chainID := range chainIDs {
		keeper.SetChannelForChain(ctx, chainID, "channel-"+chainID)
	}

	gotChainIDs := keeper.GetAllChainsWithChannels(ctx)
	require.ElementsMatch(t, chainIDs, gotChainIDs)
}

func TestSetGetChainForChannel(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)

	channelID := "test-channel"
	chainID := "test-chain"

	keeper.SetChainForChannel(ctx, channelID, chainID)
	gotChainID, found := keeper.GetChainForChannel(ctx, channelID)

	require.True(t, found)
	require.Equal(t, chainID, gotChainID)

	_, found = keeper.GetChainForChannel(ctx, "non-existent-channel")
	require.False(t, found)
}
