package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/Srikarrao1/liquidity/testutil/keeper"
	"github.com/Srikarrao1/liquidity/x/liquid/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.LiquidKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
