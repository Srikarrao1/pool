package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/Srikarrao1/liquidity/testutil/keeper"
	"github.com/Srikarrao1/liquidity/x/liquidity/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.LiquidityKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
