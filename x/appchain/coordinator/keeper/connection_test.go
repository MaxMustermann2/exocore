package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	connectiontypes "github.com/cosmos/ibc-go/v7/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibctmtypes "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"

	testutil "github.com/ExocoreNetwork/exocore/testutil/keeper"
	commontypes "github.com/ExocoreNetwork/exocore/x/appchain/common/types"
)

func TestKeeper_VerifySubscriberChain(t *testing.T) {
	k, ctx, mocks := testutil.NewCoordinatorKeeper(t)

	tests := []struct {
		name           string
		connectionHops []string
		setup          func()
		wantErr        bool
		errorType      error
	}{
		{
			name:           "success",
			connectionHops: []string{"connection-0"},
			setup: func() {
				connection := connectiontypes.ConnectionEnd{ClientId: "07-tendermint-0"}
				mocks.ConnectionKeeper.EXPECT().GetConnection(ctx, "connection-0").Return(connection, true)

				clientState := ibctmtypes.ClientState{ChainId: "test-chain-1"}
				mocks.ClientKeeper.EXPECT().GetClientState(ctx, "07-tendermint-0").Return(&clientState, true)

				// Instead of mocking Keeper.GetClientForChain, we'll set up the state
				k.SetClientForChain(ctx, "test-chain-1", "07-tendermint-0")
			},
			wantErr: false,
		},
		// ... other test cases remain the same
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}

			err := k.VerifySubscriberChain(ctx, "", tt.connectionHops)

			if tt.wantErr {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.errorType)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestKeeper_SetSubscriberChain(t *testing.T) {
	k, ctx, mocks := testutil.NewCoordinatorKeeper(t)

	tests := []struct {
		name      string
		channelID string
		setup     func()
		wantErr   bool
		errorType error
	}{
		{
			name:      "success",
			channelID: "channel-0",
			setup: func() {
				channel := channeltypes.Channel{ConnectionHops: []string{"connection-0"}}
				mocks.ChannelKeeper.EXPECT().GetChannel(ctx, commontypes.CoordinatorPortID, "channel-0").Return(channel, true)

				connection := connectiontypes.ConnectionEnd{ClientId: "07-tendermint-0"}
				mocks.ConnectionKeeper.EXPECT().GetConnection(ctx, "connection-0").Return(connection, true)

				clientState := ibctmtypes.ClientState{ChainId: "test-chain-1"}
				mocks.ClientKeeper.EXPECT().GetClientState(ctx, "07-tendermint-0").Return(&clientState, true)

				// Instead of mocking Keeper methods, we'll check the state after calling SetSubscriberChain
			},
			wantErr: false,
		},
		// ... other test cases remain the same
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}

			err := k.SetSubscriberChain(ctx, tt.channelID)

			if tt.wantErr {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.errorType)
			} else {
				assert.NoError(t, err)
				// Check the state after SetSubscriberChain
				channelID, found := k.GetChannelForChain(ctx, "test-chain-1")
				assert.True(t, found)
				assert.Equal(t, tt.channelID, channelID)

				chainID, found := k.GetChainForChannel(ctx, tt.channelID)
				assert.True(t, found)
				assert.Equal(t, "test-chain-1", chainID)

				height, found := k.GetInitChainHeight(ctx, "test-chain-1")
				assert.True(t, found)
				assert.Equal(t, uint64(ctx.BlockHeight()), height)
			}
		})
	}
}

func TestKeeper_InitChainHeight(t *testing.T) {
	k, ctx, _ := testutil.NewCoordinatorKeeper(t)

	chainID := "test-chain-1"
	height := uint64(100)

	// Test SetInitChainHeight and GetInitChainHeight
	k.SetInitChainHeight(ctx, chainID, height)

	gotHeight, found := k.GetInitChainHeight(ctx, chainID)
	require.True(t, found)
	assert.Equal(t, height, gotHeight)

	// Test GetInitChainHeight for non-existent chain
	_, found = k.GetInitChainHeight(ctx, "non-existent-chain")
	assert.False(t, found)
}
