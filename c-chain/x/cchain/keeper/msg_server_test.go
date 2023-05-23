package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/conare/c-chain/testutil/keeper"
	"github.com/conare/c-chain/x/cchain/keeper"
	"github.com/conare/c-chain/x/cchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
