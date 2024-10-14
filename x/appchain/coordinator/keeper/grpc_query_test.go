package keeper_test

import (
	"testing"

	testutil "github.com/ExocoreNetwork/exocore/testutil/keeper"
	commontypes "github.com/ExocoreNetwork/exocore/x/appchain/common/types"
	"github.com/ExocoreNetwork/exocore/x/appchain/coordinator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestQueryParams(t *testing.T) {
	keeper, ctx, _ := testutil.NewCoordinatorKeeper(t)

	t.Run("Valid request", func(t *testing.T) {
		expectedParams := types.DefaultParams()
		keeper.SetParams(ctx, expectedParams)

		response, err := keeper.QueryParams(sdk.WrapSDKContext(ctx), &types.QueryParamsRequest{})
		require.NoError(t, err)
		assert.Equal(t, expectedParams, response.Params)
	})

	t.Run("Nil request", func(t *testing.T) {
		response, err := keeper.QueryParams(sdk.WrapSDKContext(ctx), nil)
		require.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
	})
}

func TestQuerySubscriberGenesis(t *testing.T) {
	keeper, ctx, _ := testutil.NewCoordinatorKeeper(t)

	t.Run("Existing subscriber genesis", func(t *testing.T) {
		expectedGenesis := commontypes.SubscriberGenesisState{
			Params:      commontypes.SubscriberParams{},
			Coordinator: commontypes.CoordinatorInfo{},
			// Add other fields as necessary
		}
		keeper.SetSubscriberGenesis(ctx, "test-chain", &expectedGenesis)

		response, err := keeper.QuerySubscriberGenesis(sdk.WrapSDKContext(ctx), &types.QuerySubscriberGenesisRequest{
			Chain: "test-chain",
		})
		require.NoError(t, err)
		assert.Equal(t, expectedGenesis, response.SubscriberGenesis)
	})

	t.Run("Non-existent subscriber genesis", func(t *testing.T) {
		response, err := keeper.QuerySubscriberGenesis(sdk.WrapSDKContext(ctx), &types.QuerySubscriberGenesisRequest{
			Chain: "non-existent-chain",
		})
		require.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, codes.NotFound, status.Code(err))
	})

	t.Run("Nil request", func(t *testing.T) {
		response, err := keeper.QuerySubscriberGenesis(sdk.WrapSDKContext(ctx), nil)
		require.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
	})
}
