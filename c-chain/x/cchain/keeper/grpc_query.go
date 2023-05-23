package keeper

import (
	"github.com/conare/c-chain/x/cchain/types"
)

var _ types.QueryServer = Keeper{}
