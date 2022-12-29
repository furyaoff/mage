package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"

	"github.com/furya-official/mage/x/evmutil/types"
)

const (
	// EvmDenom is the gas denom used by the evm
	EvmDenom = "aMage"

	// CosmosDenom is the gas denom used by the mage app
	CosmosDenom = "umage"
)

// ConversionMultiplier is the conversion multiplier between aMage and umage
var ConversionMultiplier = sdk.NewInt(1_000_000_000_000)

var _ evmtypes.BankKeeper = EvmBankKeeper{}

// EvmBankKeeper is a BankKeeper wrapper for the x/evm module to allow the use
// of the 18 decimal aMage coin on the evm.
// x/evm consumes gas and send coins by minting and burning aMage coins in its module
// account and then sending the funds to the target account.
// This keeper uses both the umage coin and a separate aMage balance to manage the
// extra percision needed by the evm.
type EvmBankKeeper struct {
	aMageKeeper Keeper
	bk          types.BankKeeper
	ak          types.AccountKeeper
}

func NewEvmBankKeeper(aMageKeeper Keeper, bk types.BankKeeper, ak types.AccountKeeper) EvmBankKeeper {
	return EvmBankKeeper{
		aMageKeeper: aMageKeeper,
		bk:          bk,
		ak:          ak,
	}
}

// GetBalance returns the total **spendable** balance of aMage for a given account by address.
func (k EvmBankKeeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	if denom != EvmDenom {
		panic(fmt.Errorf("only evm denom %s is supported by EvmBankKeeper", EvmDenom))
	}

	spendableCoins := k.bk.SpendableCoins(ctx, addr)
	umage := spendableCoins.AmountOf(CosmosDenom)
	aMage := k.aMageKeeper.GetBalance(ctx, addr)
	total := umage.Mul(ConversionMultiplier).Add(aMage)
	return sdk.NewCoin(EvmDenom, total)
}

// SendCoinsFromModuleToAccount transfers aMage coins from a ModuleAccount to an AccAddress.
// It will panic if the module account does not exist. An error is returned if the recipient
// address is black-listed or if sending the tokens fails.
func (k EvmBankKeeper) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	umage, aMage, err := SplitAMageCoins(amt)
	if err != nil {
		return err
	}

	if umage.Amount.IsPositive() {
		if err := k.bk.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, sdk.NewCoins(umage)); err != nil {
			return err
		}
	}

	senderAddr := k.GetModuleAddress(senderModule)
	if err := k.ConvertOneUMageToAMageIfNeeded(ctx, senderAddr, aMage); err != nil {
		return err
	}

	if err := k.aMageKeeper.SendBalance(ctx, senderAddr, recipientAddr, aMage); err != nil {
		return err
	}

	return k.ConvertAMageToUMage(ctx, recipientAddr)
}

// SendCoinsFromAccountToModule transfers aMage coins from an AccAddress to a ModuleAccount.
// It will panic if the module account does not exist.
func (k EvmBankKeeper) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	umage, aMageNeeded, err := SplitAMageCoins(amt)
	if err != nil {
		return err
	}

	if umage.IsPositive() {
		if err := k.bk.SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, sdk.NewCoins(umage)); err != nil {
			return err
		}
	}

	if err := k.ConvertOneUMageToAMageIfNeeded(ctx, senderAddr, aMageNeeded); err != nil {
		return err
	}

	recipientAddr := k.GetModuleAddress(recipientModule)
	if err := k.aMageKeeper.SendBalance(ctx, senderAddr, recipientAddr, aMageNeeded); err != nil {
		return err
	}

	return k.ConvertAMageToUMage(ctx, recipientAddr)
}

// MintCoins mints aMage coins by minting the equivalent umage coins and any remaining aMage coins.
// It will panic if the module account does not exist or is unauthorized.
func (k EvmBankKeeper) MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
	umage, aMage, err := SplitAMageCoins(amt)
	if err != nil {
		return err
	}

	if umage.IsPositive() {
		if err := k.bk.MintCoins(ctx, moduleName, sdk.NewCoins(umage)); err != nil {
			return err
		}
	}

	recipientAddr := k.GetModuleAddress(moduleName)
	if err := k.aMageKeeper.AddBalance(ctx, recipientAddr, aMage); err != nil {
		return err
	}

	return k.ConvertAMageToUMage(ctx, recipientAddr)
}

// BurnCoins burns aMage coins by burning the equivalent umage coins and any remaining aMage coins.
// It will panic if the module account does not exist or is unauthorized.
func (k EvmBankKeeper) BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
	umage, aMage, err := SplitAMageCoins(amt)
	if err != nil {
		return err
	}

	if umage.IsPositive() {
		if err := k.bk.BurnCoins(ctx, moduleName, sdk.NewCoins(umage)); err != nil {
			return err
		}
	}

	moduleAddr := k.GetModuleAddress(moduleName)
	if err := k.ConvertOneUMageToAMageIfNeeded(ctx, moduleAddr, aMage); err != nil {
		return err
	}

	return k.aMageKeeper.RemoveBalance(ctx, moduleAddr, aMage)
}

// ConvertOneUMageToAMageIfNeeded converts 1 umage to aMage for an address if
// its aMage balance is smaller than the aMageNeeded amount.
func (k EvmBankKeeper) ConvertOneUMageToAMageIfNeeded(ctx sdk.Context, addr sdk.AccAddress, aMageNeeded sdk.Int) error {
	aMageBal := k.aMageKeeper.GetBalance(ctx, addr)
	if aMageBal.GTE(aMageNeeded) {
		return nil
	}

	umageToStore := sdk.NewCoins(sdk.NewCoin(CosmosDenom, sdk.OneInt()))
	if err := k.bk.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, umageToStore); err != nil {
		return err
	}

	// add 1umage equivalent of aMage to addr
	aMageToReceive := ConversionMultiplier
	if err := k.aMageKeeper.AddBalance(ctx, addr, aMageToReceive); err != nil {
		return err
	}

	return nil
}

// ConvertAMageToUMage converts all available aMage to umage for a given AccAddress.
func (k EvmBankKeeper) ConvertAMageToUMage(ctx sdk.Context, addr sdk.AccAddress) error {
	totalAMage := k.aMageKeeper.GetBalance(ctx, addr)
	umage, _, err := SplitAMageCoins(sdk.NewCoins(sdk.NewCoin(EvmDenom, totalAMage)))
	if err != nil {
		return err
	}

	// do nothing if account does not have enough aMage for a single umage
	umageToReceive := umage.Amount
	if !umageToReceive.IsPositive() {
		return nil
	}

	// remove aMage used for converting to umage
	aMageToBurn := umageToReceive.Mul(ConversionMultiplier)
	finalBal := totalAMage.Sub(aMageToBurn)
	if err := k.aMageKeeper.SetBalance(ctx, addr, finalBal); err != nil {
		return err
	}

	fromAddr := k.GetModuleAddress(types.ModuleName)
	if err := k.bk.SendCoins(ctx, fromAddr, addr, sdk.NewCoins(umage)); err != nil {
		return err
	}

	return nil
}

func (k EvmBankKeeper) GetModuleAddress(moduleName string) sdk.AccAddress {
	addr := k.ak.GetModuleAddress(moduleName)
	if addr == nil {
		panic(sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", moduleName))
	}
	return addr
}

// SplitAMageCoins splits aMage coins to the equivalent umage coins and any remaining aMage balance.
// An error will be returned if the coins are not valid or if the coins are not the aMage denom.
func SplitAMageCoins(coins sdk.Coins) (sdk.Coin, sdk.Int, error) {
	aMage := sdk.ZeroInt()
	umage := sdk.NewCoin(CosmosDenom, sdk.ZeroInt())

	if len(coins) == 0 {
		return umage, aMage, nil
	}

	if err := ValidateEvmCoins(coins); err != nil {
		return umage, aMage, err
	}

	// note: we should always have len(coins) == 1 here since coins cannot have dup denoms after we validate.
	coin := coins[0]
	remainingBalance := coin.Amount.Mod(ConversionMultiplier)
	if remainingBalance.IsPositive() {
		aMage = remainingBalance
	}
	umageAmount := coin.Amount.Quo(ConversionMultiplier)
	if umageAmount.IsPositive() {
		umage = sdk.NewCoin(CosmosDenom, umageAmount)
	}

	return umage, aMage, nil
}

// ValidateEvmCoins validates the coins from evm is valid and is the EvmDenom (aMage).
func ValidateEvmCoins(coins sdk.Coins) error {
	if len(coins) == 0 {
		return nil
	}

	// validate that coins are non-negative, sorted, and no dup denoms
	if err := coins.Validate(); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, coins.String())
	}

	// validate that coin denom is aMage
	if len(coins) != 1 || coins[0].Denom != EvmDenom {
		errMsg := fmt.Sprintf("invalid evm coin denom, only %s is supported", EvmDenom)
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, errMsg)
	}

	return nil
}
