package keeper

import (
	"context"
	"github.com/Srikarrao1/liquidity/x/liquidity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"errors"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	nextPoolId, found := k.GetNextPoolID(ctx)
	if !found {
		nextPoolId = 1
	}

	// Initial amounts must be greater than zero and should form a valid constant product
	if msg.InitialAmountA <= 0 || msg.InitialAmountB <= 0 {
		return nil, errors.New("initial amounts must be greater than zero")
	}

	pool := types.Pool {
		Id:      nextPoolId,
		Builder: msg.Builder,
		AssetA: msg.AssetA,
		AssetB: msg.AssetB,
		ReserveA: msg.InitialAmountA,
		ReserveB: msg.InitialAmountB,
	}

	k.SetPool(ctx, pool)
	k.SetNextPoolID(ctx, nextPoolId+1)


	return &types.MsgCreatePoolResponse{}, nil
}