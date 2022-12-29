package ante_test

import (
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"

	"github.com/furya-official/mage/app"
	"github.com/furya-official/mage/app/ante"
)

func mustParseDecCoins(value string) sdk.DecCoins {
	coins, err := sdk.ParseDecCoins(strings.ReplaceAll(value, ";", ","))
	if err != nil {
		panic(err)
	}

	return coins
}

func TestEvmMinGasFilter(t *testing.T) {
	tApp := app.NewTestApp()
	handler := ante.NewEvmMinGasFilter(tApp.GetEvmKeeper())

	ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
	tApp.GetEvmKeeper().SetParams(ctx, evmtypes.Params{
		EvmDenom: "aMage",
	})

	testCases := []struct {
		name                 string
		minGasPrices         sdk.DecCoins
		expectedMinGasPrices sdk.DecCoins
	}{
		{
			"no min gas prices",
			mustParseDecCoins(""),
			mustParseDecCoins(""),
		},
		{
			"zero uMage gas price",
			mustParseDecCoins("0uMage"),
			mustParseDecCoins("0uMage"),
		},
		{
			"non-zero uMage gas price",
			mustParseDecCoins("0.001uMage"),
			mustParseDecCoins("0.001uMage"),
		},
		{
			"zero uMage gas price, min aMage price",
			mustParseDecCoins("0uMage;100000aMage"),
			mustParseDecCoins("0uMage"), // aMage is removed
		},
		{
			"zero uMage gas price, min aMage price, other token",
			mustParseDecCoins("0uMage;100000aMage;0.001other"),
			mustParseDecCoins("0uMage;0.001other"), // aMage is removed
		},
		{
			"non-zero uMage gas price, min aMage price",
			mustParseDecCoins("0.25uMage;100000aMage;0.001other"),
			mustParseDecCoins("0.25uMage;0.001other"), // aMage is removed
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := tApp.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})

			ctx = ctx.WithMinGasPrices(tc.minGasPrices)
			mmd := MockAnteHandler{}

			_, err := handler.AnteHandle(ctx, nil, false, mmd.AnteHandle)
			require.NoError(t, err)
			require.True(t, mmd.WasCalled)

			assert.NoError(t, mmd.CalledCtx.MinGasPrices().Validate())
			assert.Equal(t, tc.expectedMinGasPrices, mmd.CalledCtx.MinGasPrices())
		})
	}
}
