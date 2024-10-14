package keeper

import (
	"testing"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	commontypes "github.com/ExocoreNetwork/exocore/x/appchain/common/types"
	"github.com/ExocoreNetwork/exocore/x/appchain/coordinator/keeper"
	"github.com/ExocoreNetwork/exocore/x/appchain/coordinator/types"
)

// MockedKeepers contains all the mocked keepers
type MockedKeepers struct {
	AVSKeeper        *types.MockAVSKeeper
	EpochsKeeper     *types.MockEpochsKeeper
	StakingKeeper    *types.MockStakingKeeper
	OperatorKeeper   *types.MockOperatorKeeper
	DelegationKeeper *types.MockDelegationKeeper
	ClientKeeper     *commontypes.MockClientKeeper
	PortKeeper       *commontypes.MockPortKeeper
	ScopedKeeper     *commontypes.MockScopedKeeper
	ChannelKeeper    *commontypes.MockChannelKeeper
	ConnectionKeeper *commontypes.MockConnectionKeeper
	AccountKeeper    *commontypes.MockAccountKeeper
}

func NewCoordinatorKeeper(t testing.TB) (keeper.Keeper, sdk.Context, MockedKeepers) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.TestingLogger())

	// Create mock controller
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	// Create mock keepers
	mockedKeepers := MockedKeepers{
		AVSKeeper:        types.NewMockAVSKeeper(ctrl),
		EpochsKeeper:     types.NewMockEpochsKeeper(ctrl),
		StakingKeeper:    types.NewMockStakingKeeper(ctrl),
		OperatorKeeper:   types.NewMockOperatorKeeper(ctrl),
		DelegationKeeper: types.NewMockDelegationKeeper(ctrl),
		ClientKeeper:     commontypes.NewMockClientKeeper(ctrl),
		PortKeeper:       commontypes.NewMockPortKeeper(ctrl),
		ScopedKeeper:     commontypes.NewMockScopedKeeper(ctrl),
		ChannelKeeper:    commontypes.NewMockChannelKeeper(ctrl),
		ConnectionKeeper: commontypes.NewMockConnectionKeeper(ctrl),
		AccountKeeper:    commontypes.NewMockAccountKeeper(ctrl),
	}

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		mockedKeepers.AVSKeeper,
		mockedKeepers.EpochsKeeper,
		mockedKeepers.OperatorKeeper,
		mockedKeepers.StakingKeeper,
		mockedKeepers.DelegationKeeper,
		mockedKeepers.ClientKeeper,
		mockedKeepers.PortKeeper,
		mockedKeepers.ScopedKeeper,
		mockedKeepers.ChannelKeeper,
		mockedKeepers.ConnectionKeeper,
		mockedKeepers.AccountKeeper,
	)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx, mockedKeepers
}
