package liquid_test

import (
	"testing"

	keepertest "github.com/Srikarrao1/liquidity/testutil/keeper"
	"github.com/Srikarrao1/liquidity/testutil/nullify"
	liquid "github.com/Srikarrao1/liquidity/x/liquid/module"
	"github.com/Srikarrao1/liquidity/x/liquid/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LiquidKeeper(t)
	liquid.InitGenesis(ctx, k, genesisState)
	got := liquid.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
