package cchain_test

import (
	"testing"

	keepertest "github.com/conare/c-chain/testutil/keeper"
	"github.com/conare/c-chain/testutil/nullify"
	"github.com/conare/c-chain/x/cchain"
	"github.com/conare/c-chain/x/cchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CchainKeeper(t)
	cchain.InitGenesis(ctx, *k, genesisState)
	got := cchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
