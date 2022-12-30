package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/furya-official/mage/x/earn/types"
)

const (
	bmageDenom  = "bmage"
	bmagePrefix = bmageDenom + "-"
)

// GetParams returns the params from the store
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var p types.Params
	k.paramSubspace.GetParamSet(ctx, &p)

	return p
}

// SetParams sets params on the store
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSubspace.SetParamSet(ctx, &params)
}

// GetAllowedVaults returns the list of allowed vaults from the module params.
func (k Keeper) GetAllowedVaults(ctx sdk.Context) types.AllowedVaults {
	return k.GetParams(ctx).AllowedVaults
}

// getAllowedVaultRaw returns a single vault from the module params specified
// by the denom.
func (k Keeper) getAllowedVaultRaw(
	ctx sdk.Context,
	vaultDenom string,
) (types.AllowedVault, bool) {
	for _, allowedVault := range k.GetAllowedVaults(ctx) {
		if allowedVault.Denom == vaultDenom {
			return allowedVault, true
		}
	}

	return types.AllowedVault{}, false
}

// GetAllowedVault returns the AllowedVault that corresponds to the
// given denom. If the denom starts with "bmage-" where it will return the
// "bmage" AllowedVault. Otherwise, it will return the exact match for the
// corresponding AllowedVault denom.
func (k *Keeper) GetAllowedVault(
	ctx sdk.Context,
	vaultDenom string,
) (types.AllowedVault, bool) {
	if strings.HasPrefix(vaultDenom, bmagePrefix) {
		return k.getAllowedVaultRaw(ctx, bmageDenom)
	}

	return k.getAllowedVaultRaw(ctx, vaultDenom)
}
