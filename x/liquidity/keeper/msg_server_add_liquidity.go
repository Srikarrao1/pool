package keeper

import (
    "context"
    sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
    "github.com/Srikarrao1/liquidity/x/liquidity/types"


)

func (k msgServer) AddLiquidity(goCtx context.Context, msg *types.MsgAddLiquidity) (*types.MsgAddLiquidityResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)

    // Fetch the pool
    pool, found := k.GetPool(ctx, msg.PoolId)
    if !found {
        return nil,  errorsmod.Wrap(sdkerrors.ErrUnknownRequest, "pool not found")
    }

    // Update the reserves
    pool.ReserveA += msg.AmountA
    pool.ReserveB += msg.AmountB

    // Save the updated pool
    k.SetPool(ctx, pool)

    return &types.MsgAddLiquidityResponse{}, nil
}

// func (k Keeper) GetPool(ctx sdk.Context, id uint64) (types.Pool, bool) {
// 	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))

//     key := types.KeyPrefix(types.PoolKey)
//     key = append(key, sdk.Uint64ToBigEndian(id)...)

//     if !store.Has(key) {
//         return types.Pool{}, false
//     }

//     var pool types.Pool
//     b := store.Get(key)
//     k.cdc.MustUnmarshal(b, &pool)
//     return pool, true
// }

// func (k Keeper) SetPool(ctx sdk.Context, pool types.Pool) {
//     store := ctx.KVStore(k.storeKey)
//     key := types.KeyPrefix(types.PoolKey)
//     key = append(key, sdk.Uint64ToBigEndian(pool.Id)...)

//     value := k.cdc.MustMarshal(&pool)
//     store.Set(key, value)
// }
