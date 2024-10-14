package keeper_test

import (
	"errors"
	"testing"

	testutil "github.com/ExocoreNetwork/exocore/testutil/keeper"
	commontypes "github.com/ExocoreNetwork/exocore/x/appchain/common/types"
	"github.com/ExocoreNetwork/exocore/x/appchain/coordinator/types"
	epochstypes "github.com/ExocoreNetwork/exocore/x/epochs/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestKeeper_AddSubscriberChain(t *testing.T) {
	k, ctx, mocks := testutil.NewCoordinatorKeeper(t)

	tests := []struct {
		name    string
		req     *types.RegisterSubscriberChainRequest
		setup   func()
		wantErr bool
	}{
		{
			name: "success",
			req: &types.RegisterSubscriberChainRequest{
				ChainID:              "test-chain-1",
				EpochIdentifier:      "day",
				AssetIDs:             []string{"1", "2"},
				MinSelfDelegationUsd: uint64(1000),
				SubscriberParams: commontypes.SubscriberParams{
					UnbondingPeriod: 100,
				},
				FromAddress: "exo1...",
			},
			setup: func() {
				mocks.EpochsKeeper.EXPECT().GetEpochInfo(ctx, "day").Return(epochstypes.EpochInfo{
					Identifier:   "day",
					Duration:     86400,
					CurrentEpoch: 10,
				}, true)
				mocks.AVSKeeper.EXPECT().RegisterAVSWithChainID(ctx, gomock.Any()).Return(common.Address{}, nil)
			},
			wantErr: false,
		},
		{
			name:    "nil request",
			req:     nil,
			wantErr: true,
		},
		{
			name: "epoch not found",
			req: &types.RegisterSubscriberChainRequest{
				ChainID:         "test-chain-2",
				EpochIdentifier: "week",
			},
			setup: func() {
				mocks.EpochsKeeper.EXPECT().GetEpochInfo(ctx, "week").Return(epochstypes.EpochInfo{}, false)
			},
			wantErr: true,
		},
		{
			name: "avs registration fails",
			req: &types.RegisterSubscriberChainRequest{
				ChainID:              "test-chain-3",
				EpochIdentifier:      "day",
				AssetIDs:             []string{"1"},
				MinSelfDelegationUsd: uint64(1000),
				SubscriberParams: commontypes.SubscriberParams{
					UnbondingPeriod: 100,
				},
				FromAddress: "exo1...",
			},
			setup: func() {
				mocks.EpochsKeeper.EXPECT().GetEpochInfo(ctx, "day").Return(epochstypes.EpochInfo{
					Identifier:   "day",
					Duration:     86400,
					CurrentEpoch: 10,
				}, true)
				mocks.AVSKeeper.EXPECT().RegisterAVSWithChainID(ctx, gomock.Any()).Return(common.Address{}, errors.New("avs registration failed"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}

			resp, err := k.AddSubscriberChain(ctx, tt.req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)

				// Check if the pending subscriber chain was appended
				pendingChains := k.GetPendingSubChains(ctx, tt.req.EpochIdentifier, uint64(10))
				require.Len(t, pendingChains.List, 1)
				assert.Equal(t, *tt.req, pendingChains.List[0])
			}
		})
	}
}
