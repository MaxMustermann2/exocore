package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ExocoreNetwork/exocore/testutil/keeper"
	commontypes "github.com/ExocoreNetwork/exocore/x/appchain/common/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
)

func TestSetGetSubscriberValidatorForChain(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)

	chainID := "test-chain"
	pubKey := ed25519.GenPrivKey().PubKey()
	consAddr := sdk.ConsAddress(pubKey.Address())
	validator, err := commontypes.NewSubscriberValidator(consAddr, 100, pubKey)
	require.NoError(t, err)

	// Test setting and getting
	keeper.SetSubscriberValidatorForChain(ctx, chainID, validator)
	gotValidator, found := keeper.GetSubscriberValidatorForChain(ctx, chainID, consAddr)

	require.True(t, found)
	require.Equal(t, validator, gotValidator)

	// Test getting non-existent validator
	_, found = keeper.GetSubscriberValidatorForChain(ctx, chainID, sdk.ConsAddress("non-existent"))
	require.False(t, found)
}

func TestGetAllSubscriberValidatorsForChain(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)

	chainID := "test-chain"
	validators := []commontypes.SubscriberValidator{}

	for i := 0; i < 3; i++ {
		pubKey := ed25519.GenPrivKey().PubKey()
		consAddr := sdk.ConsAddress(pubKey.Address())
		validator, err := commontypes.NewSubscriberValidator(consAddr, 100+int64(i), pubKey)
		require.NoError(t, err)
		validators = append(validators, validator)
		keeper.SetSubscriberValidatorForChain(ctx, chainID, validator)
	}

	gotValidators := keeper.GetAllSubscriberValidatorsForChain(ctx, chainID)
	require.Equal(t, len(validators), len(gotValidators))
	require.ElementsMatch(t, validators, gotValidators)
}

func TestDeleteSubscriberValidatorForChain(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)

	chainID := "test-chain"
	pubKey := ed25519.GenPrivKey().PubKey()
	consAddr := sdk.ConsAddress(pubKey.Address())
	validator, err := commontypes.NewSubscriberValidator(consAddr, 100, pubKey)
	require.NoError(t, err)

	keeper.SetSubscriberValidatorForChain(ctx, chainID, validator)
	keeper.DeleteSubscriberValidatorForChain(ctx, chainID, consAddr)

	_, found := keeper.GetSubscriberValidatorForChain(ctx, chainID, consAddr)
	require.False(t, found)
}

func TestSetGetMaxValidatorsForChain(t *testing.T) {
	keeper, ctx, _ := testkeeper.NewCoordinatorKeeper(t)

	chainID := "test-chain"
	maxValidators := uint32(100)

	keeper.SetMaxValidatorsForChain(ctx, chainID, maxValidators)
	gotMaxValidators := keeper.GetMaxValidatorsForChain(ctx, chainID)

	require.Equal(t, maxValidators, gotMaxValidators)
}
