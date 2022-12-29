package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furya-official/mage/app"
)

func init() {
	MageConfig := sdk.GetConfig()
	app.SetBech32AddressPrefixes(MageConfig)
	app.SetBip44CoinType(MageConfig)
	MageConfig.Seal()
}
