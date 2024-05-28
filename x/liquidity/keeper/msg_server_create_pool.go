package keeper

import (
	"context"
	"errors"
	"fmt"

	"github.com/Srikarrao1/liquidity/x/liquidity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	fmt.Println("cscds==========12345")
	ctx := sdk.UnwrapSDKContext(goCtx)

	nextPoolId, found := k.GetNextPoolID(ctx)
	if !found {
		nextPoolId = 1
	}

	fmt.Println("cscds==========12345", nextPoolId)

	// Initial amounts must be greater than zero and should form a valid constant product
	if msg.InitialAmountA <= 0 || msg.InitialAmountB <= 0 {
		return nil, errors.New("initial amounts must be greater than zero")
	}

	pool := types.Pool{
		Id:       nextPoolId,
		Creator:  msg.Creator,
		Builder:  msg.Builder,
		AssetA:   msg.AssetA,
		AssetB:   msg.AssetB,
		ReserveA: msg.InitialAmountA,
		ReserveB: msg.InitialAmountB,
	}

	fmt.Println("cscds==========pool", pool)

	k.SetPool(ctx, pool)
	fmt.Println("cscds==========pool")
	k.SetNextPoolID(ctx, nextPoolId+1)
	fmt.Println("cscds==========pool78")
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreatePool,
			sdk.NewAttribute(types.AttributeKeyPoolId, fmt.Sprintf("%d", pool.Id)),
			sdk.NewAttribute(types.AttributeKeyTokenA, msg.AssetA),
			sdk.NewAttribute(types.AttributeKeyTokenA, msg.AssetB),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyBuilder, msg.Builder),
		),
	)

	return &types.MsgCreatePoolResponse{}, nil
}
