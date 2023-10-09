// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)
package restaking_assets_manage

import (
	"cosmossdk.io/math"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/exocore/x/restaking_assets_manage/keeper"
	types2 "github.com/exocore/x/restaking_assets_manage/types"
)

// NewGenesisState - Create a new genesis state
func NewGenesisState(chain []*types2.ClientChainInfo, token []*types2.ClientChainTokenInfo) *types2.GenesisState {
	return &types2.GenesisState{
		DefaultSupportedClientChains:      chain,
		DefaultSupportedClientChainTokens: token,
	}
}

// DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() *types2.GenesisState {
	return NewGenesisState([]*types2.ClientChainInfo{}, []*types2.ClientChainTokenInfo{})
}

// GetGenesisStateFromAppState returns x/restaking_assets_manage GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.Codec, appState map[string]json.RawMessage) types2.GenesisState {
	var genesisState types2.GenesisState

	if appState[types2.ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[types2.ModuleName], &genesisState)
	}

	return genesisState
}

// ValidateGenesis performs basic validation of restaking_assets_manage genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data types2.GenesisState) error {
	//todo: check the validation of client chain and token info
	return nil
}

// InitGenesis import module genesis
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	data types2.GenesisState,
) {
	//todo: might need to sort the clientChains and tokens before handling.

	//save default supported client chain
	for _, chain := range data.DefaultSupportedClientChains {
		k.SetClientChainInfo(chain)
	}
	//save default supported client chain assets
	for _, asset := range data.DefaultSupportedClientChainTokens {
		k.SetStakingAssetInfo(&types2.StakingAssetInfo{
			AssetBasicInfo:     asset,
			StakingTotalAmount: math.NewUint(0),
		})
	}
}

// ExportGenesis export module status
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types2.GenesisState {
	clientChainList := make([]*types2.ClientChainInfo, 0)
	clientChainInfo, _ := k.GetAllClientChainInfo()
	for _, v := range clientChainInfo {
		clientChainList = append(clientChainList, v)
	}

	clientChainAssetsList := make([]*types2.ClientChainTokenInfo, 0)
	clientChainAssets, _ := k.GetAllStakingAssetsInfo()
	for _, v := range clientChainAssets {
		clientChainAssetsList = append(clientChainAssetsList, v.AssetBasicInfo)
	}
	return &types2.GenesisState{
		DefaultSupportedClientChains:      clientChainList,
		DefaultSupportedClientChainTokens: clientChainAssetsList,
	}
}
