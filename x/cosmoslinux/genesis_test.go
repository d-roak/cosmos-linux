package cosmoslinux_test

import (
	"testing"

	keepertest "cosmos-linux/testutil/keeper"
	"cosmos-linux/testutil/nullify"
	"cosmos-linux/x/cosmoslinux"
	"cosmos-linux/x/cosmoslinux/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmoslinuxKeeper(t)
	cosmoslinux.InitGenesis(ctx, *k, genesisState)
	got := cosmoslinux.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
