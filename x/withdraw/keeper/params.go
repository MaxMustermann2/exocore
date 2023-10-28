package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	types "github.com/exocore/x/withdraw/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (*types.Params, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ParamsKey)
	ifExist := store.Has(types.ParamsKey)
	if !ifExist {
		return nil, types.ErrNoParamsKey
	}

	value := store.Get(types.ParamsKey)

	ret := &types.Params{}
	k.cdc.MustUnmarshal(value, ret)
	return ret, nil
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params *types.Params) error {
	//check if addr is evm address
	if !common.IsHexAddress(params.ExoCoreLzAppAddress) {
		return types.ErrInvalidEvmAddressFormat
	}
	if len(common.FromHex(params.ExoCoreLzAppEventTopic)) != common.HashLength {
		return types.ErrInvalidLzUaTopicIdLength
	}
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(params)
	store.Set(types.ParamsKey, bz)
	return nil
}
