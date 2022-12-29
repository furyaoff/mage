package Magedist_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	tmtime "github.com/tendermint/tendermint/types/time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/furya-official/mage/x/Magedist"
	testutil "github.com/furya-official/mage/x/Magedist/testutil"
	"github.com/furya-official/mage/x/Magedist/types"
)

type genesisTestSuite struct {
	testutil.Suite
}

func (suite *genesisTestSuite) TestInitGenesis_ValidationPanic() {
	invalidState := types.NewGenesisState(
		types.Params{
			Active: true,
			Periods: []types.Period{
				{
					Start:     time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
					End:       tmtime.Canonical(time.Unix(1, 0)),
					Inflation: sdk.OneDec(),
				},
			},
		},
		tmtime.Canonical(time.Unix(1, 0)),
	)

	suite.Require().Panics(func() {
		Magedist.InitGenesis(suite.Ctx, suite.Keeper, suite.AccountKeeper, invalidState)
	}, "expected init genesis to panic with invalid state")
}

func (suite *genesisTestSuite) TestInitAndExportGenesis() {
	state := types.NewGenesisState(
		types.Params{
			Active: true,
			Periods: []types.Period{
				{
					Start:     time.Date(2021, 1, 1, 1, 1, 1, 1, time.UTC),
					End:       time.Date(2021, 2, 1, 1, 1, 1, 1, time.UTC),
					Inflation: sdk.OneDec(),
				},
			},
		},
		time.Date(2020, 1, 2, 1, 1, 1, 1, time.UTC),
	)

	Magedist.InitGenesis(suite.Ctx, suite.Keeper, suite.AccountKeeper, state)
	suite.Require().Equal(state.Params, suite.Keeper.GetParams(suite.Ctx))

	exportedState := Magedist.ExportGenesis(suite.Ctx, suite.Keeper)
	suite.Require().Equal(state, exportedState)
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(genesisTestSuite))
}
