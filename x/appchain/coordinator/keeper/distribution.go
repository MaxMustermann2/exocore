package keeper

import (
	"github.com/ExocoreNetwork/exocore/x/appchain/coordinator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetSubscriberRewardsPoolAddressStr gets the subscriber rewards pool address string.
// It is the bech32 string corresponding to a hardcoded module account.
func (k Keeper) GetSubscriberRewardsPoolAddressStr(ctx sdk.Context) string {
	return k.accountKeeper.GetModuleAccount(
		ctx, types.SubscriberRewardsPool,
	).GetAddress().String()
}

// TODO: distribution implementation
// (1) validate that the subscriber has sent rewards to the pool and if not penalize them
// (2) forward rewards to the fee collector account
