package keeper

import (
	"fmt"

	assetstype "github.com/ExocoreNetwork/exocore/x/assets/types"
	delegationtype "github.com/ExocoreNetwork/exocore/x/delegation/types"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
// Since this action typically occurs on chain starts, this function is allowed to panic.
func (k Keeper) InitGenesis(
	ctx sdk.Context,
	genState delegationtype.GenesisState,
) []abci.ValidatorUpdate {
	for _, a := range genState.DelegationsByStakerAssetOperator {
		stakerID := a.StakerID
		for _, b := range a.DelegationsByAssetOperator {
			assetID := b.AssetID
			for _, c := range b.DelegationsByOperator {
				operatorAddress := c.OperatorAddress
				// #nosec G703 // already validated
				accAddress, _ := sdk.AccAddressFromBech32(operatorAddress)
				amount := c.Amount // delegation amount
				if !k.operatorKeeper.IsOperator(ctx, accAddress) {
					// the operator must be registered first, so the
					// genesis of the operator module comes before this module
					panic(
						fmt.Sprintf(
							"%s: %s", delegationtype.ErrOperatorNotExist,
							fmt.Sprintf("input operatorAddr is:%s", operatorAddress),
						),
					)
				}
				// at genesis, the operator cannot be frozen so skip that check.
				// validate that enough deposits exist before delegation.
				info, err := k.assetsKeeper.GetStakerSpecifiedAssetInfo(ctx, stakerID, assetID)
				if err != nil {
					panic(err)
				}
				if amount.GT(info.WithdrawableAmount) {
					panic(
						fmt.Sprintf(
							"delegated amount %s is greater than the staker's asset amount %s",
							amount.String(), info.WithdrawableAmount.String(),
						),
					)
				}
				delegatorAndAmount := make(map[string]*delegationtype.DelegationAmounts, 1)
				delegatorAndAmount[operatorAddress] = &delegationtype.DelegationAmounts{
					UndelegatableAmount: amount,
				}
				if err := k.UpdateDelegationState(
					ctx, stakerID, assetID, delegatorAndAmount,
				); err != nil {
					panic(err)
				}
				if err := k.UpdateStakerDelegationTotalAmount(
					ctx, stakerID, assetID, amount,
				); err != nil {
					panic(err)
				}
				// also tell the assetsKeeper to mark this as a delegation.
				if err := k.assetsKeeper.UpdateStakerAssetState(
					ctx, stakerID, assetID, assetstype.StakerSingleAssetChangeInfo{
						WithdrawableAmount: amount.Neg(),
					}); err != nil {
					panic(err)
				}
				// we have checked that delegation amount > deposit amount for each asset.
				// we don't need to check for the total amount, since this genesis only handles
				// delegation amounts (others are implicitly zero).
			}
		}
	}
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the module's exported genesis
func (Keeper) ExportGenesis(sdk.Context) *delegationtype.GenesisState {
	genesis := delegationtype.DefaultGenesis()
	// TODO
	return genesis
}
