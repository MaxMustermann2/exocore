package keeper

import (
	"context"
	"strings"

	delegationtype "github.com/ExocoreNetwork/exocore/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ delegationtype.QueryServer = &Keeper{}

func (k *Keeper) QuerySingleDelegationInfo(ctx context.Context, req *delegationtype.SingleDelegationInfoReq) (*delegationtype.SingleDelegationInfoResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)
	delegationAmounts, err := k.GetSingleDelegationInfo(c, strings.ToLower(req.StakerID), strings.ToLower(req.AssetID), req.OperatorAddr)
	if err != nil {
		return nil, err
	}
	// calculate the maximum undelegatable amount
	singleAmount, err := k.UndelegatableAmount(c, strings.ToLower(req.AssetID), req.OperatorAddr, delegationAmounts)
	if err != nil {
		return nil, err
	}
	return &delegationtype.SingleDelegationInfoResponse{
		DelegationAmounts:      delegationAmounts,
		MaxUndelegatableAmount: singleAmount,
	}, nil
}

func (k *Keeper) QueryDelegationInfo(ctx context.Context, info *delegationtype.DelegationInfoReq) (*delegationtype.QueryDelegationInfoResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)
	return k.GetDelegationInfo(c, strings.ToLower(info.StakerID), strings.ToLower(info.AssetID))
}

func (k *Keeper) QueryUndelegations(ctx context.Context, req *delegationtype.UndelegationsReq) (*delegationtype.UndelegationRecordList, error) {
	c := sdk.UnwrapSDKContext(ctx)
	undelegations, err := k.GetStakerUndelegationRecords(c, strings.ToLower(req.StakerID), strings.ToLower(req.AssetID))
	if err != nil {
		return nil, err
	}
	return &delegationtype.UndelegationRecordList{
		Undelegations: undelegations,
	}, nil
}

func (k *Keeper) QueryUndelegationsByHeight(ctx context.Context, req *delegationtype.UndelegationsByHeightReq) (*delegationtype.UndelegationRecordList, error) {
	c := sdk.UnwrapSDKContext(ctx)
	undelegations, err := k.GetPendingUndelegationRecords(c, req.BlockHeight)
	if err != nil {
		return nil, err
	}
	return &delegationtype.UndelegationRecordList{
		Undelegations: undelegations,
	}, nil
}

func (k Keeper) QueryUndelegationHoldCount(ctx context.Context, req *delegationtype.UndelegationHoldCountReq) (*delegationtype.UndelegationHoldCountResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)
	res := k.GetUndelegationHoldCount(c, []byte(req.RecordKey))
	return &delegationtype.UndelegationHoldCountResponse{HoldCount: res}, nil
}

func (k Keeper) QueryAssociatedOperatorByStaker(ctx context.Context, req *delegationtype.QueryAssociatedOperatorByStakerReq) (*delegationtype.QueryAssociatedOperatorByStakerResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)
	operator, err := k.GetAssociatedOperator(c, strings.ToLower(req.StakerID))
	if err != nil {
		return nil, err
	}
	return &delegationtype.QueryAssociatedOperatorByStakerResponse{
		Operator: operator,
	}, nil
}

func (k Keeper) QueryAssociatedStakersByOperator(ctx context.Context, req *delegationtype.QueryAssociatedStakersByOperatorReq) (*delegationtype.QueryAssociatedStakersByOperatorResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)
	stakers, err := k.GetAssociatedStakers(c, req.Operator)
	if err != nil {
		return nil, err
	}
	return &delegationtype.QueryAssociatedStakersByOperatorResponse{
		Stakers: stakers,
	}, nil
}
