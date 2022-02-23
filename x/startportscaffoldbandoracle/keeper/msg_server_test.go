package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mkXultra/startport_scaffold_band_oracle/testutil/keeper"
	"github.com/mkXultra/startport_scaffold_band_oracle/x/startportscaffoldbandoracle/keeper"
	"github.com/mkXultra/startport_scaffold_band_oracle/x/startportscaffoldbandoracle/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.StartportscaffoldbandoracleKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
