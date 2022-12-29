package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/furya-official/mage/x/auction/types"
)

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSubspace.SetParamSet(ctx, &params)
}

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSubspace.GetParamSet(ctx, &params)
	return
}
