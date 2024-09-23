package keeper

import (
	"math"
	"strings"

	"github.com/ExocoreNetwork/exocore/x/delegation/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAllUndelegations function returns all the undelegation records in the module.
// It is used during `ExportGenesis` to export the undelegation records.
func (k Keeper) GetAllUndelegations(
	ctx sdk.Context,
) (undelegations []types.UndelegationRecord, err error) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixUndelegationInfo)
	defer iterator.Close()

	ret := make([]types.UndelegationRecord, 0)
	for ; iterator.Valid(); iterator.Next() {
		var undelegation types.UndelegationRecord
		k.cdc.MustUnmarshal(iterator.Value(), &undelegation)
		ret = append(ret, undelegation)
	}
	return ret, nil
}

// SetUndelegationRecords stores the provided undelegation records.
// The records are stored with 3 different keys:
// (1) recordKey == blockNumber + lzNonce + txHash + operatorAddress => record
// (2) stakerID + assetID + lzNonce => recordKey
// (3) completeBlockNumber + lzNonce => recordKey
// If a record exists with the same key, it will be overwritten; however, that is not a big
// concern since the lzNonce and txHash are unique for each record.
func (k *Keeper) SetUndelegationRecords(
	ctx sdk.Context, records []types.UndelegationRecord,
) error {
	// initialize the 3 stores
	singleRecordStore := prefix.NewStore(
		ctx.KVStore(k.storeKey), types.KeyPrefixUndelegationInfo,
	)
	stakerUndelegationStore := prefix.NewStore(
		ctx.KVStore(k.storeKey), types.KeyPrefixStakerUndelegationInfo,
	)
	pendingUndelegationStore := prefix.NewStore(
		ctx.KVStore(k.storeKey), types.KeyPrefixPendingUndelegations,
	)
	// for validation, use the block height of the current block
	currentHeight := ctx.BlockHeight()
	for i := range records {
		record := records[i] // avoid implicit memory aliasing
		if record.CompleteBlockNumber < uint64(currentHeight) {
			return types.ErrInvalidCompletedHeight.Wrapf(
				"currentHeight:%d, CompleteBlockNumber:%d",
				currentHeight, record.CompleteBlockNumber,
			)
		}
		bz := k.cdc.MustMarshal(&record)

		// store 1
		singleRecKey := types.GetUndelegationRecordKey(
			record.BlockNumber, record.LzTxNonce, record.TxHash, record.OperatorAddr,
		)
		singleRecordStore.Set(singleRecKey, bz)

		// store 2
		stakerKey := types.GetStakerUndelegationRecordKey(
			record.StakerID, record.AssetID, record.LzTxNonce,
		)
		stakerUndelegationStore.Set(stakerKey, singleRecKey)

		// store 3
		pendingUndelegationKey := types.GetPendingUndelegationRecordKey(
			record.CompleteBlockNumber, record.LzTxNonce,
		)
		pendingUndelegationStore.Set(pendingUndelegationKey, singleRecKey)
	}
	return nil
}

// DeleteUndelegationRecord deletes the undelegation record from the module.
// The deletion is performed from all the 3 stores.
func (k *Keeper) DeleteUndelegationRecord(
	ctx sdk.Context, record *types.UndelegationRecord,
) error {
	parentStore := ctx.KVStore(k.storeKey)
	// initialize the 3 stores
	singleRecordStore := prefix.NewStore(
		parentStore, types.KeyPrefixUndelegationInfo,
	)
	stakerUndelegationStore := prefix.NewStore(
		parentStore, types.KeyPrefixStakerUndelegationInfo,
	)
	pendingUndelegationStore := prefix.NewStore(
		parentStore, types.KeyPrefixPendingUndelegations,
	)
	// store 1
	singleRecKey := types.GetUndelegationRecordKey(
		record.BlockNumber, record.LzTxNonce, record.TxHash, record.OperatorAddr,
	)
	singleRecordStore.Delete(singleRecKey)
	// store 2
	stakerKey := types.GetStakerUndelegationRecordKey(
		record.StakerID, record.AssetID, record.LzTxNonce,
	)
	stakerUndelegationStore.Delete(stakerKey)
	// store 3
	pendingUndelegationKey := types.GetPendingUndelegationRecordKey(
		record.CompleteBlockNumber, record.LzTxNonce,
	)
	pendingUndelegationStore.Delete(pendingUndelegationKey)
	// never throws an error but for future proofing
	return nil
}

// GetUndelegationRecords returns the undelegation records for the provided record keys.
func (k *Keeper) GetUndelegationRecords(
	ctx sdk.Context, singleRecordKeys []string,
) (record []*types.UndelegationRecord, err error) {
	store := prefix.NewStore(
		ctx.KVStore(k.storeKey), types.KeyPrefixUndelegationInfo,
	)
	ret := make([]*types.UndelegationRecord, 0)
	for _, singleRecordKey := range singleRecordKeys {
		keyBytes := []byte(singleRecordKey)
		value := store.Get(keyBytes)
		if value == nil {
			return nil, types.ErrNoKeyInTheStore.Wrapf(
				"GetSingleDelegationRecord: key is %s", singleRecordKey,
			)
		}
		undelegationRecord := types.UndelegationRecord{}
		k.cdc.MustUnmarshal(value, &undelegationRecord)
		ret = append(ret, &undelegationRecord)
	}
	return ret, nil
}

// IterateUndelegationsByOperator iterates over the undelegation records belonging to the
// provided operator bech32 address and filter. If the filter is non-nil, it will only
// iterate over the records for which the block height is greater than or equal to the filter.
func (k *Keeper) IterateUndelegationsByOperator(
	ctx sdk.Context, operator string, heightFilter *uint64, isUpdate bool,
	opFunc func(undelegation *types.UndelegationRecord) error,
) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixUndelegationInfo)
	iterator := sdk.KVStorePrefixIterator(store, []byte(operator))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		if heightFilter != nil {
			keyFields, err := types.ParseUndelegationRecordKey(iterator.Key())
			if err != nil {
				return err
			}
			if keyFields.BlockHeight < *heightFilter {
				continue
			}
		}
		undelegation := &types.UndelegationRecord{}
		k.cdc.MustUnmarshal(iterator.Value(), undelegation)
		err := opFunc(undelegation)
		if err != nil {
			return err
		}

		if isUpdate {
			bz := k.cdc.MustMarshal(undelegation)
			store.Set(iterator.Key(), bz)
		}
	}
	return nil
}

// GetStakerUndelegationRecKeys returns the undelegation record keys corresponding to the
// provided staker and asset.
func (k *Keeper) GetStakerUndelegationRecKeys(
	ctx sdk.Context, stakerID string, assetID string,
) (recordKeyList []string, err error) {
	store := prefix.NewStore(
		ctx.KVStore(k.storeKey), types.KeyPrefixStakerUndelegationInfo,
	)
	iterator := sdk.KVStorePrefixIterator(
		store, []byte(strings.Join([]string{stakerID, assetID}, "/")),
	)
	defer iterator.Close()

	ret := make([]string, 0)
	for ; iterator.Valid(); iterator.Next() {
		ret = append(ret, string(iterator.Value()))
	}
	return ret, nil
}

// GetStakerUndelegationRecords returns the undelegation records for the provided staker and
// asset.
func (k *Keeper) GetStakerUndelegationRecords(
	ctx sdk.Context, stakerID string, assetID string,
) (records []*types.UndelegationRecord, err error) {
	recordKeys, err := k.GetStakerUndelegationRecKeys(ctx, stakerID, assetID)
	if err != nil {
		return nil, err
	}
	return k.GetUndelegationRecords(ctx, recordKeys)
}

// GetPendingUndelegationRecKeys returns the undelegation record keys scheduled to mature at or
// before the end of the block with the provided height.
func (k *Keeper) GetPendingUndelegationRecKeys(
	ctx sdk.Context, height uint64,
) (recordKeyList []string, err error) {
	iterator := sdk.KVStorePrefixIterator(
		ctx.KVStore(k.storeKey), types.KeyPrefixPendingUndelegations,
	)
	defer iterator.Close()
	ret := make([]string, 0)
	for ; iterator.Valid(); iterator.Next() {
		// key is prefix + completeBlockNumber + lzNonce
		itemHeight, _ := types.ParsePendingUndelegationRecordKey(iterator.Key()[1:])
		if itemHeight > height {
			// the records are sorted by the block height, so we can break here
			break
		}
		ret = append(ret, string(iterator.Value()))
	}
	return ret, nil
}

// GetPendingUndelegationRecords returns the undelegation records scheduled to mature at the end
// of the block with the provided height.
func (k *Keeper) GetPendingUndelegationRecords(
	ctx sdk.Context, height uint64,
) (records []*types.UndelegationRecord, err error) {
	recordKeys, err := k.GetPendingUndelegationRecKeys(ctx, height)
	if err != nil {
		return nil, err
	}
	if len(recordKeys) == 0 {
		records = make([]*types.UndelegationRecord, 0)
		return records, nil
	}
	return k.GetUndelegationRecords(ctx, recordKeys)
}

// IncrementUndelegationHoldCount increments the hold count for the undelegation record key.
func (k Keeper) IncrementUndelegationHoldCount(ctx sdk.Context, recordKey []byte) error {
	prev := k.GetUndelegationHoldCount(ctx, recordKey)
	if prev == math.MaxUint64 {
		return types.ErrCannotIncHoldCount
	}
	now := prev + 1
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetUndelegationOnHoldKey(recordKey), sdk.Uint64ToBigEndian(now))
	return nil
}

// GetUndelegationHoldCount returns the hold count for the undelegation record key.
func (k *Keeper) GetUndelegationHoldCount(ctx sdk.Context, recordKey []byte) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetUndelegationOnHoldKey(recordKey))
	return sdk.BigEndianToUint64(bz)
}

// DecrementUndelegationHoldCount decrements the hold count for the undelegation record key.
func (k Keeper) DecrementUndelegationHoldCount(ctx sdk.Context, recordKey []byte) error {
	prev := k.GetUndelegationHoldCount(ctx, recordKey)
	if prev == 0 {
		return types.ErrCannotDecHoldCount
	}
	now := prev - 1
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetUndelegationOnHoldKey(recordKey), sdk.Uint64ToBigEndian(now))
	return nil
}
