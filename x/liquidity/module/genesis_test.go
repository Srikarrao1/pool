package liquidity_test

import (
	"testing"

	keepertest "github.com/Srikarrao1/liquidity/testutil/keeper"
	"github.com/Srikarrao1/liquidity/testutil/nullify"
	liquidity "github.com/Srikarrao1/liquidity/x/liquidity/module"
	"github.com/Srikarrao1/liquidity/x/liquidity/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LiquidityKeeper(t)
	liquidity.InitGenesis(ctx, k, genesisState)
	got := liquidity.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
