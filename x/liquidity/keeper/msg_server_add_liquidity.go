package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/Srikarrao1/liquidity/x/liquidity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddLiquidity(goCtx context.Context, msg *types.MsgAddLiquidity) (*types.MsgAddLiquidityResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)

    // Fetch the pool
    pool, found := k.GetPool(ctx, msg.PoolId)
    if !found {
        return nil,  errorsmod.Wrap(sdkerrors.ErrUnknownRequest, "pool not found")
    }

    // if msg.TokenA != pool.AssetA || msg.TokenB != pool.AssetB {
    //     return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "tokens do not match pool's tokens")
    // }

      // Ensure the provided amounts maintain the X*Y=K curve
      if !isValidAddition(int64(pool.ReserveA), int64(pool.ReserveB), int64(msg.AmountA), int64(msg.AmountB)) {
        return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid amounts for maintaining the constant product formula")
    }

     // Calculate the number of LP tokens to mint
     mintAmount := calculateLPTokens(uint64(pool.TotalLiquidity), uint64(pool.ReserveA)+uint64(msg.AmountA), uint64(pool.ReserveB)+uint64(msg.AmountB))


    // Update the reserves
    pool.ReserveA += msg.AmountA
    pool.ReserveB += msg.AmountB
    pool.TotalLiquidity += mintAmount

    // Save the updated pool
    k.SetPool(ctx, pool)

    return &types.MsgAddLiquidityResponse{}, nil
}

func isValidAddition(reserveA, reserveB, amountA, amountB int64) bool {
    // Calculate the product of the current reserves
    k := reserveA * reserveB
    // Calculate the product of the new reserves after adding liquidity
    newReserveA := reserveA + amountA
    newReserveB := reserveB + amountB
    newK := newReserveA * newReserveB

    // Ensure the new product is not less than the current product
    return newK >= k
}

// Function to calculate the number of LP tokens to mint based on the current pool state and the amount being added
func calculateLPTokens(totalLiquidity, newReserveA, newReserveB uint64) uint64 {
    if totalLiquidity == 0 {
        return 1000000 // Initial liquidity
    }
    return (totalLiquidity * totalLiquidity) / (newReserveA * newReserveB)
}
