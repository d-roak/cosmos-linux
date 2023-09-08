package keeper_test

import (
	"testing"

	testkeeper "cosmos-linux/testutil/keeper"
	"cosmos-linux/x/cosmoslinux/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CosmoslinuxKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
