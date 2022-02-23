package startportscaffoldbandoracle_test

import (
	"testing"

	keepertest "github.com/mkXultra/startport_scaffold_band_oracle/testutil/keeper"
	"github.com/mkXultra/startport_scaffold_band_oracle/x/startportscaffoldbandoracle"
	"github.com/mkXultra/startport_scaffold_band_oracle/x/startportscaffoldbandoracle/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StartportscaffoldbandoracleKeeper(t)
	startportscaffoldbandoracle.InitGenesis(ctx, *k, genesisState)
	got := startportscaffoldbandoracle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
