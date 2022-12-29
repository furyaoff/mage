package community_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/furya-official/mage/app"
	"github.com/furya-official/mage/x/community/types"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	tApp := app.NewTestApp()
	ctx := tApp.NewContext(true, tmproto.Header{Height: 1})
	tApp.InitializeFromGenesisStates()

	accKeeper := tApp.GetAccountKeeper()
	acc := accKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.ModuleName))
	require.NotNil(t, acc)
}
