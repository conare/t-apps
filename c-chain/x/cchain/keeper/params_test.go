package keeper_test

import (
	"testing"

	testkeeper "github.com/conare/c-chain/testutil/keeper"
	"github.com/conare/c-chain/x/cchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
