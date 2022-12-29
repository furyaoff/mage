package Magemint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furya-official/mage/x/Magemint/keeper"
)

// BeginBlocker mints & distributes new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.KeeperI) {
	if err := k.AccumulateAndMintInflation(ctx); err != nil {
		panic(err)
	}
}
