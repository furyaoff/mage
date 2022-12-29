package hard

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furya-official/mage/x/hard/keeper"
)

// BeginBlocker updates interest rates
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	k.ApplyInterestRateUpdates(ctx)
}
