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
	EvmDenom = "amage"

	// CosmosDenom is the gas denom used by the mage app
	CosmosDenom = "umage"
)

// ConversionMultiplier is the conversion multiplier between amage and umage
var ConversionMultiplier = sdk.NewInt(1_000_000_000_000)

var _ evmtypes.BankKeeper = EvmBankKeeper{}

// EvmBankKeeper is a BankKeeper wrapper for the x/evm module to allow the use
// of the 18 decimal amage coin on the evm.
// x/evm consumes gas and send coins by minting and burning amage coins in its module
// account and then sending the funds to the target account.
// This keeper uses both the umage coin and a separate amage balance to manage the
// extra percision needed by the evm.
type EvmBankKeeper struct {
	amageKeeper Keeper
	bk          types.BankKeeper
	ak          types.AccountKeeper
}

func NewEvmBankKeeper(amageKeeper Keeper, bk types.BankKeeper, ak types.AccountKeeper) EvmBankKeeper {
	return EvmBankKeeper{
		amageKeeper: amageKeeper,
		bk:          bk,
		ak:          ak,
	}
}

// GetBalance returns the total **spendable** balance of amage for a given account by address.
func (k EvmBankKeeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	if denom != EvmDenom {
		panic(fmt.Errorf("only evm denom %s is supported by EvmBankKeeper", EvmDenom))
	}

	spendableCoins := k.bk.SpendableCoins(ctx, addr)
	umage := spendableCoins.AmountOf(CosmosDenom)
	amage := k.amageKeeper.GetBalance(ctx, addr)
	total := umage.Mul(ConversionMultiplier).Add(amage)
	return sdk.NewCoin(EvmDenom, total)
}

// SendCoinsFromModuleToAccount transfers amage coins from a ModuleAccount to an AccAddress.
// It will panic if the module account does not exist. An error is returned if the recipient
// address is black-listed or if sending the tokens fails.
func (k EvmBankKeeper) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	umage, amage, err := SplitAmageCoins(amt)
	if err != nil {
		return err
	}

	if umage.Amount.IsPositive() {
		if err := k.bk.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, sdk.NewCoins(umage)); err != nil {
			return err
		}
	}

	senderAddr := k.GetModuleAddress(senderModule)
	if err := k.ConvertOneUmageToAmageIfNeeded(ctx, senderAddr, amage); err != nil {
		return err
	}

	if err := k.amageKeeper.SendBalance(ctx, senderAddr, recipientAddr, amage); err != nil {
		return err
	}

	return k.ConvertAmageToUmage(ctx, recipientAddr)
}

// SendCoinsFromAccountToModule transfers amage coins from an AccAddress to a ModuleAccount.
// It will panic if the module account does not exist.
func (k EvmBankKeeper) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	umage, amageNeeded, err := SplitAmageCoins(amt)
	if err != nil {
		return err
	}

	if umage.IsPositive() {
		if err := k.bk.SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, sdk.NewCoins(umage)); err != nil {
			return err
		}
	}

	if err := k.ConvertOneUmageToAmageIfNeeded(ctx, senderAddr, amageNeeded); err != nil {
		return err
	}

	recipientAddr := k.GetModuleAddress(recipientModule)
	if err := k.amageKeeper.SendBalance(ctx, senderAddr, recipientAddr, amageNeeded); err != nil {
		return err
	}

	return k.ConvertAmageToUmage(ctx, recipientAddr)
}

// MintCoins mints amage coins by minting the equivalent umage coins and any remaining amage coins.
// It will panic if the module account does not exist or is unauthorized.
func (k EvmBankKeeper) MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
	umage, amage, err := SplitAmageCoins(amt)
	if err != nil {
		return err
	}

	if umage.IsPositive() {
		if err := k.bk.MintCoins(ctx, moduleName, sdk.NewCoins(umage)); err != nil {
			return err
		}
	}

	recipientAddr := k.GetModuleAddress(moduleName)
	if err := k.amageKeeper.AddBalance(ctx, recipientAddr, amage); err != nil {
		return err
	}

	return k.ConvertAmageToUmage(ctx, recipientAddr)
}

// BurnCoins burns amage coins by burning the equivalent umage coins and any remaining amage coins.
// It will panic if the module account does not exist or is unauthorized.
func (k EvmBankKeeper) BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
	umage, amage, err := SplitAmageCoins(amt)
	if err != nil {
		return err
	}

	if umage.IsPositive() {
		if err := k.bk.BurnCoins(ctx, moduleName, sdk.NewCoins(umage)); err != nil {
			return err
		}
	}

	moduleAddr := k.GetModuleAddress(moduleName)
	if err := k.ConvertOneUmageToAmageIfNeeded(ctx, moduleAddr, amage); err != nil {
		return err
	}

	return k.amageKeeper.RemoveBalance(ctx, moduleAddr, amage)
}

// ConvertOneUmageToAmageIfNeeded converts 1 umage to amage for an address if
// its amage balance is smaller than the amageNeeded amount.
func (k EvmBankKeeper) ConvertOneUmageToAmageIfNeeded(ctx sdk.Context, addr sdk.AccAddress, amageNeeded sdk.Int) error {
	amageBal := k.amageKeeper.GetBalance(ctx, addr)
	if amageBal.GTE(amageNeeded) {
		return nil
	}

	umageToStore := sdk.NewCoins(sdk.NewCoin(CosmosDenom, sdk.OneInt()))
	if err := k.bk.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, umageToStore); err != nil {
		return err
	}

	// add 1umage equivalent of amage to addr
	amageToReceive := ConversionMultiplier
	if err := k.amageKeeper.AddBalance(ctx, addr, amageToReceive); err != nil {
		return err
	}

	return nil
}

// ConvertAmageToUmage converts all available amage to umage for a given AccAddress.
func (k EvmBankKeeper) ConvertAmageToUmage(ctx sdk.Context, addr sdk.AccAddress) error {
	totalAmage := k.amageKeeper.GetBalance(ctx, addr)
	umage, _, err := SplitAmageCoins(sdk.NewCoins(sdk.NewCoin(EvmDenom, totalAmage)))
	if err != nil {
		return err
	}

	// do nothing if account does not have enough amage for a single umage
	umageToReceive := umage.Amount
	if !umageToReceive.IsPositive() {
		return nil
	}

	// remove amage used for converting to umage
	amageToBurn := umageToReceive.Mul(ConversionMultiplier)
	finalBal := totalAmage.Sub(amageToBurn)
	if err := k.amageKeeper.SetBalance(ctx, addr, finalBal); err != nil {
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

// SplitAmageCoins splits amage coins to the equivalent umage coins and any remaining amage balance.
// An error will be returned if the coins are not valid or if the coins are not the amage denom.
func SplitAmageCoins(coins sdk.Coins) (sdk.Coin, sdk.Int, error) {
	amage := sdk.ZeroInt()
	umage := sdk.NewCoin(CosmosDenom, sdk.ZeroInt())

	if len(coins) == 0 {
		return umage, amage, nil
	}

	if err := ValidateEvmCoins(coins); err != nil {
		return umage, amage, err
	}

	// note: we should always have len(coins) == 1 here since coins cannot have dup denoms after we validate.
	coin := coins[0]
	remainingBalance := coin.Amount.Mod(ConversionMultiplier)
	if remainingBalance.IsPositive() {
		amage = remainingBalance
	}
	umageAmount := coin.Amount.Quo(ConversionMultiplier)
	if umageAmount.IsPositive() {
		umage = sdk.NewCoin(CosmosDenom, umageAmount)
	}

	return umage, amage, nil
}

// ValidateEvmCoins validates the coins from evm is valid and is the EvmDenom (amage).
func ValidateEvmCoins(coins sdk.Coins) error {
	if len(coins) == 0 {
		return nil
	}

	// validate that coins are non-negative, sorted, and no dup denoms
	if err := coins.Validate(); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, coins.String())
	}

	// validate that coin denom is amage
	if len(coins) != 1 || coins[0].Denom != EvmDenom {
		errMsg := fmt.Sprintf("invalid evm coin denom, only %s is supported", EvmDenom)
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, errMsg)
	}

	return nil
}
