package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furya-official/mage/app"
)

func init() {
	mageConfig := sdk.GetConfig()
	app.SetBech32AddressPrefixes(mageConfig)
	app.SetBip44CoinType(mageConfig)
	mageConfig.Seal()
}
