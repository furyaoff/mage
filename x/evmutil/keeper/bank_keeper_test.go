package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	tmtime "github.com/tendermint/tendermint/types/time"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vesting "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"

	"github.com/furya-official/mage/x/evmutil/keeper"
	"github.com/furya-official/mage/x/evmutil/testutil"
	"github.com/furya-official/mage/x/evmutil/types"
)

type evmBankKeeperTestSuite struct {
	testutil.Suite
}

func (suite *evmBankKeeperTestSuite) SetupTest() {
	suite.Suite.SetupTest()
}

func (suite *evmBankKeeperTestSuite) TestGetBalance_ReturnsSpendable() {
	startingCoins := sdk.NewCoins(sdk.NewInt64Coin("umage", 10))
	startingAMage := sdk.NewInt(100)

	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)
	bacc := authtypes.NewBaseAccountWithAddress(suite.Addrs[0])
	vacc := vesting.NewContinuousVestingAccount(bacc, startingCoins, now.Unix(), endTime.Unix())
	suite.AccountKeeper.SetAccount(suite.Ctx, vacc)

	err := suite.App.FundAccount(suite.Ctx, suite.Addrs[0], startingCoins)
	suite.Require().NoError(err)
	err = suite.Keeper.SetBalance(suite.Ctx, suite.Addrs[0], startingAMage)
	suite.Require().NoError(err)

	coin := suite.EvmBankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "aMage")
	suite.Require().Equal(startingAMage, coin.Amount)

	ctx := suite.Ctx.WithBlockTime(now.Add(12 * time.Hour))
	coin = suite.EvmBankKeeper.GetBalance(ctx, suite.Addrs[0], "aMage")
	suite.Require().Equal(sdk.NewIntFromUint64(5_000_000_000_100), coin.Amount)
}

func (suite *evmBankKeeperTestSuite) TestGetBalance_NotEvmDenom() {
	suite.Require().Panics(func() {
		suite.EvmBankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "umage")
	})
	suite.Require().Panics(func() {
		suite.EvmBankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "busd")
	})
}

func (suite *evmBankKeeperTestSuite) TestGetBalance() {
	tests := []struct {
		name           string
		startingAmount sdk.Coins
		expAmount      sdk.Int
	}{
		{
			"umage with aMage",
			sdk.NewCoins(
				sdk.NewInt64Coin("aMage", 100),
				sdk.NewInt64Coin("umage", 10),
			),
			sdk.NewInt(10_000_000_000_100),
		},
		{
			"just aMage",
			sdk.NewCoins(
				sdk.NewInt64Coin("aMage", 100),
				sdk.NewInt64Coin("busd", 100),
			),
			sdk.NewInt(100),
		},
		{
			"just umage",
			sdk.NewCoins(
				sdk.NewInt64Coin("umage", 10),
				sdk.NewInt64Coin("busd", 100),
			),
			sdk.NewInt(10_000_000_000_000),
		},
		{
			"no umage or aMage",
			sdk.NewCoins(),
			sdk.ZeroInt(),
		},
		{
			"with avaka that is more than 1 umage",
			sdk.NewCoins(
				sdk.NewInt64Coin("aMage", 20_000_000_000_220),
				sdk.NewInt64Coin("umage", 11),
			),
			sdk.NewInt(31_000_000_000_220),
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()

			suite.FundAccountWithMage(suite.Addrs[0], tt.startingAmount)
			coin := suite.EvmBankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "aMage")
			suite.Require().Equal(tt.expAmount, coin.Amount)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestSendCoinsFromModuleToAccount() {
	startingModuleCoins := sdk.NewCoins(
		sdk.NewInt64Coin("aMage", 200),
		sdk.NewInt64Coin("umage", 100),
	)
	tests := []struct {
		name           string
		sendCoins      sdk.Coins
		startingAccBal sdk.Coins
		expAccBal      sdk.Coins
		hasErr         bool
	}{
		{
			"send more than 1 umage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 12_000_000_000_010)),
			sdk.Coins{},
			sdk.NewCoins(
				sdk.NewInt64Coin("aMage", 10),
				sdk.NewInt64Coin("umage", 12),
			),
			false,
		},
		{
			"send less than 1 umage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 122)),
			sdk.Coins{},
			sdk.NewCoins(
				sdk.NewInt64Coin("aMage", 122),
				sdk.NewInt64Coin("umage", 0),
			),
			false,
		},
		{
			"send an exact amount of umage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 98_000_000_000_000)),
			sdk.Coins{},
			sdk.NewCoins(
				sdk.NewInt64Coin("aMage", 0o0),
				sdk.NewInt64Coin("umage", 98),
			),
			false,
		},
		{
			"send no aMage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 0)),
			sdk.Coins{},
			sdk.NewCoins(
				sdk.NewInt64Coin("aMage", 0),
				sdk.NewInt64Coin("umage", 0),
			),
			false,
		},
		{
			"errors if sending other coins",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 500), sdk.NewInt64Coin("busd", 1000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if not enough total aMage to cover",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 100_000_000_001_000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if not enough umage to cover",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 200_000_000_000_000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"converts receiver's aMage to umage if there's enough aMage after the transfer",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 99_000_000_000_200)),
			sdk.NewCoins(
				sdk.NewInt64Coin("aMage", 999_999_999_900),
				sdk.NewInt64Coin("umage", 1),
			),
			sdk.NewCoins(
				sdk.NewInt64Coin("aMage", 100),
				sdk.NewInt64Coin("umage", 101),
			),
			false,
		},
		{
			"converts all of receiver's aMage to umage even if somehow receiver has more than 1umage of aMage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 12_000_000_000_100)),
			sdk.NewCoins(
				sdk.NewInt64Coin("aMage", 5_999_999_999_990),
				sdk.NewInt64Coin("umage", 1),
			),
			sdk.NewCoins(
				sdk.NewInt64Coin("aMage", 90),
				sdk.NewInt64Coin("umage", 19),
			),
			false,
		},
		{
			"swap 1 umage for aMage if module account doesn't have enough aMage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 99_000_000_001_000)),
			sdk.NewCoins(
				sdk.NewInt64Coin("aMage", 200),
				sdk.NewInt64Coin("umage", 1),
			),
			sdk.NewCoins(
				sdk.NewInt64Coin("aMage", 1200),
				sdk.NewInt64Coin("umage", 100),
			),
			false,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()

			suite.FundAccountWithMage(suite.Addrs[0], tt.startingAccBal)
			suite.FundModuleAccountWithMage(evmtypes.ModuleName, startingModuleCoins)

			// fund our module with some umage to account for converting extra aMage back to umage
			suite.FundModuleAccountWithMage(types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin("umage", 10)))

			err := suite.EvmBankKeeper.SendCoinsFromModuleToAccount(suite.Ctx, evmtypes.ModuleName, suite.Addrs[0], tt.sendCoins)
			if tt.hasErr {
				suite.Require().Error(err)
				return
			} else {
				suite.Require().NoError(err)
			}

			// check umage
			umageSender := suite.BankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "umage")
			suite.Require().Equal(tt.expAccBal.AmountOf("umage").Int64(), umageSender.Amount.Int64())

			// check aMage
			actualAMage := suite.Keeper.GetBalance(suite.Ctx, suite.Addrs[0])
			suite.Require().Equal(tt.expAccBal.AmountOf("aMage").Int64(), actualAMage.Int64())
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestSendCoinsFromAccountToModule() {
	startingAccCoins := sdk.NewCoins(
		sdk.NewInt64Coin("aMage", 200),
		sdk.NewInt64Coin("umage", 100),
	)
	startingModuleCoins := sdk.NewCoins(
		sdk.NewInt64Coin("aMage", 100_000_000_000),
	)
	tests := []struct {
		name           string
		sendCoins      sdk.Coins
		expSenderCoins sdk.Coins
		expModuleCoins sdk.Coins
		hasErr         bool
	}{
		{
			"send more than 1 umage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 12_000_000_000_010)),
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 190), sdk.NewInt64Coin("umage", 88)),
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 100_000_000_010), sdk.NewInt64Coin("umage", 12)),
			false,
		},
		{
			"send less than 1 umage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 122)),
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 78), sdk.NewInt64Coin("umage", 100)),
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 100_000_000_122), sdk.NewInt64Coin("umage", 0)),
			false,
		},
		{
			"send an exact amount of umage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 98_000_000_000_000)),
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 200), sdk.NewInt64Coin("umage", 2)),
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 100_000_000_000), sdk.NewInt64Coin("umage", 98)),
			false,
		},
		{
			"send no aMage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 0)),
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 200), sdk.NewInt64Coin("umage", 100)),
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 100_000_000_000), sdk.NewInt64Coin("umage", 0)),
			false,
		},
		{
			"errors if sending other coins",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 500), sdk.NewInt64Coin("busd", 1000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if have dup coins",
			sdk.Coins{
				sdk.NewInt64Coin("aMage", 12_000_000_000_000),
				sdk.NewInt64Coin("aMage", 2_000_000_000_000),
			},
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if not enough total aMage to cover",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 100_000_000_001_000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"errors if not enough umage to cover",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 200_000_000_000_000)),
			sdk.Coins{},
			sdk.Coins{},
			true,
		},
		{
			"converts 1 umage to aMage if not enough aMage to cover",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 99_001_000_000_000)),
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 999_000_000_200), sdk.NewInt64Coin("umage", 0)),
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 101_000_000_000), sdk.NewInt64Coin("umage", 99)),
			false,
		},
		{
			"converts receiver's aMage to umage if there's enough aMage after the transfer",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 5_900_000_000_200)),
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 100_000_000_000), sdk.NewInt64Coin("umage", 94)),
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 200), sdk.NewInt64Coin("umage", 6)),
			false,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()
			suite.FundAccountWithMage(suite.Addrs[0], startingAccCoins)
			suite.FundModuleAccountWithMage(evmtypes.ModuleName, startingModuleCoins)

			err := suite.EvmBankKeeper.SendCoinsFromAccountToModule(suite.Ctx, suite.Addrs[0], evmtypes.ModuleName, tt.sendCoins)
			if tt.hasErr {
				suite.Require().Error(err)
				return
			} else {
				suite.Require().NoError(err)
			}

			// check sender balance
			umageSender := suite.BankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "umage")
			suite.Require().Equal(tt.expSenderCoins.AmountOf("umage").Int64(), umageSender.Amount.Int64())
			actualAMage := suite.Keeper.GetBalance(suite.Ctx, suite.Addrs[0])
			suite.Require().Equal(tt.expSenderCoins.AmountOf("aMage").Int64(), actualAMage.Int64())

			// check module balance
			moduleAddr := suite.AccountKeeper.GetModuleAddress(evmtypes.ModuleName)
			umageSender = suite.BankKeeper.GetBalance(suite.Ctx, moduleAddr, "umage")
			suite.Require().Equal(tt.expModuleCoins.AmountOf("umage").Int64(), umageSender.Amount.Int64())
			actualAMage = suite.Keeper.GetBalance(suite.Ctx, moduleAddr)
			suite.Require().Equal(tt.expModuleCoins.AmountOf("aMage").Int64(), actualAMage.Int64())
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestBurnCoins() {
	startingUMage := sdk.NewInt(100)
	tests := []struct {
		name       string
		burnCoins  sdk.Coins
		expUMage   sdk.Int
		expAMage   sdk.Int
		hasErr     bool
		aMageStart sdk.Int
	}{
		{
			"burn more than 1 umage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 12_021_000_000_002)),
			sdk.NewInt(88),
			sdk.NewInt(100_000_000_000),
			false,
			sdk.NewInt(121_000_000_002),
		},
		{
			"burn less than 1 umage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 122)),
			sdk.NewInt(100),
			sdk.NewInt(878),
			false,
			sdk.NewInt(1000),
		},
		{
			"burn an exact amount of umage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 98_000_000_000_000)),
			sdk.NewInt(2),
			sdk.NewInt(10),
			false,
			sdk.NewInt(10),
		},
		{
			"burn no aMage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 0)),
			startingUMage,
			sdk.ZeroInt(),
			false,
			sdk.ZeroInt(),
		},
		{
			"errors if burning other coins",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 500), sdk.NewInt64Coin("busd", 1000)),
			startingUMage,
			sdk.NewInt(100),
			true,
			sdk.NewInt(100),
		},
		{
			"errors if have dup coins",
			sdk.Coins{
				sdk.NewInt64Coin("aMage", 12_000_000_000_000),
				sdk.NewInt64Coin("aMage", 2_000_000_000_000),
			},
			startingUMage,
			sdk.ZeroInt(),
			true,
			sdk.ZeroInt(),
		},
		{
			"errors if burn amount is negative",
			sdk.Coins{sdk.Coin{Denom: "aMage", Amount: sdk.NewInt(-100)}},
			startingUMage,
			sdk.NewInt(50),
			true,
			sdk.NewInt(50),
		},
		{
			"errors if not enough aMage to cover burn",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 100_999_000_000_000)),
			sdk.NewInt(0),
			sdk.NewInt(99_000_000_000),
			true,
			sdk.NewInt(99_000_000_000),
		},
		{
			"errors if not enough umage to cover burn",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 200_000_000_000_000)),
			sdk.NewInt(100),
			sdk.ZeroInt(),
			true,
			sdk.ZeroInt(),
		},
		{
			"converts 1 umage to aMage if not enough aMage to cover",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 12_021_000_000_002)),
			sdk.NewInt(87),
			sdk.NewInt(980_000_000_000),
			false,
			sdk.NewInt(1_000_000_002),
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()
			startingCoins := sdk.NewCoins(
				sdk.NewCoin("umage", startingUMage),
				sdk.NewCoin("aMage", tt.aMageStart),
			)
			suite.FundModuleAccountWithMage(evmtypes.ModuleName, startingCoins)

			err := suite.EvmBankKeeper.BurnCoins(suite.Ctx, evmtypes.ModuleName, tt.burnCoins)
			if tt.hasErr {
				suite.Require().Error(err)
				return
			} else {
				suite.Require().NoError(err)
			}

			// check umage
			umageActual := suite.BankKeeper.GetBalance(suite.Ctx, suite.EvmModuleAddr, "umage")
			suite.Require().Equal(tt.expUMage, umageActual.Amount)

			// check aMage
			aMageActual := suite.Keeper.GetBalance(suite.Ctx, suite.EvmModuleAddr)
			suite.Require().Equal(tt.expAMage, aMageActual)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestMintCoins() {
	tests := []struct {
		name       string
		mintCoins  sdk.Coins
		umage      sdk.Int
		aMage      sdk.Int
		hasErr     bool
		aMageStart sdk.Int
	}{
		{
			"mint more than 1 umage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 12_021_000_000_002)),
			sdk.NewInt(12),
			sdk.NewInt(21_000_000_002),
			false,
			sdk.ZeroInt(),
		},
		{
			"mint less than 1 umage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 901_000_000_001)),
			sdk.ZeroInt(),
			sdk.NewInt(901_000_000_001),
			false,
			sdk.ZeroInt(),
		},
		{
			"mint an exact amount of umage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 123_000_000_000_000_000)),
			sdk.NewInt(123_000),
			sdk.ZeroInt(),
			false,
			sdk.ZeroInt(),
		},
		{
			"mint no aMage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 0)),
			sdk.ZeroInt(),
			sdk.ZeroInt(),
			false,
			sdk.ZeroInt(),
		},
		{
			"errors if minting other coins",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 500), sdk.NewInt64Coin("busd", 1000)),
			sdk.ZeroInt(),
			sdk.NewInt(100),
			true,
			sdk.NewInt(100),
		},
		{
			"errors if have dup coins",
			sdk.Coins{
				sdk.NewInt64Coin("aMage", 12_000_000_000_000),
				sdk.NewInt64Coin("aMage", 2_000_000_000_000),
			},
			sdk.ZeroInt(),
			sdk.ZeroInt(),
			true,
			sdk.ZeroInt(),
		},
		{
			"errors if mint amount is negative",
			sdk.Coins{sdk.Coin{Denom: "aMage", Amount: sdk.NewInt(-100)}},
			sdk.ZeroInt(),
			sdk.NewInt(50),
			true,
			sdk.NewInt(50),
		},
		{
			"adds to existing aMage balance",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 12_021_000_000_002)),
			sdk.NewInt(12),
			sdk.NewInt(21_000_000_102),
			false,
			sdk.NewInt(100),
		},
		{
			"convert aMage balance to umage if it exceeds 1 umage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 10_999_000_000_000)),
			sdk.NewInt(12),
			sdk.NewInt(1_200_000_001),
			false,
			sdk.NewInt(1_002_200_000_001),
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()
			suite.FundModuleAccountWithMage(types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin("umage", 10)))
			suite.FundModuleAccountWithMage(evmtypes.ModuleName, sdk.NewCoins(sdk.NewCoin("aMage", tt.aMageStart)))

			err := suite.EvmBankKeeper.MintCoins(suite.Ctx, evmtypes.ModuleName, tt.mintCoins)
			if tt.hasErr {
				suite.Require().Error(err)
				return
			} else {
				suite.Require().NoError(err)
			}

			// check umage
			umageActual := suite.BankKeeper.GetBalance(suite.Ctx, suite.EvmModuleAddr, "umage")
			suite.Require().Equal(tt.umage, umageActual.Amount)

			// check aMage
			aMageActual := suite.Keeper.GetBalance(suite.Ctx, suite.EvmModuleAddr)
			suite.Require().Equal(tt.aMage, aMageActual)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestValidateEvmCoins() {
	tests := []struct {
		name      string
		coins     sdk.Coins
		shouldErr bool
	}{
		{
			"valid coins",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 500)),
			false,
		},
		{
			"dup coins",
			sdk.Coins{sdk.NewInt64Coin("aMage", 500), sdk.NewInt64Coin("aMage", 500)},
			true,
		},
		{
			"not evm coins",
			sdk.NewCoins(sdk.NewInt64Coin("umage", 500)),
			true,
		},
		{
			"negative coins",
			sdk.Coins{sdk.Coin{Denom: "aMage", Amount: sdk.NewInt(-500)}},
			true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			err := keeper.ValidateEvmCoins(tt.coins)
			if tt.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestConvertOneUMageToAMageIfNeeded() {
	aMageNeeded := sdk.NewInt(200)
	tests := []struct {
		name          string
		startingCoins sdk.Coins
		expectedCoins sdk.Coins
		success       bool
	}{
		{
			"not enough umage for conversion",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 100)),
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 100)),
			false,
		},
		{
			"converts 1 umage to aMage",
			sdk.NewCoins(sdk.NewInt64Coin("umage", 10), sdk.NewInt64Coin("aMage", 100)),
			sdk.NewCoins(sdk.NewInt64Coin("umage", 9), sdk.NewInt64Coin("aMage", 1_000_000_000_100)),
			true,
		},
		{
			"conversion not needed",
			sdk.NewCoins(sdk.NewInt64Coin("umage", 10), sdk.NewInt64Coin("aMage", 200)),
			sdk.NewCoins(sdk.NewInt64Coin("umage", 10), sdk.NewInt64Coin("aMage", 200)),
			true,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()

			suite.FundAccountWithMage(suite.Addrs[0], tt.startingCoins)
			err := suite.EvmBankKeeper.ConvertOneUMageToAMageIfNeeded(suite.Ctx, suite.Addrs[0], aMageNeeded)
			moduleMage := suite.BankKeeper.GetBalance(suite.Ctx, suite.AccountKeeper.GetModuleAddress(types.ModuleName), "umage")
			if tt.success {
				suite.Require().NoError(err)
				if tt.startingCoins.AmountOf("aMage").LT(aMageNeeded) {
					suite.Require().Equal(sdk.OneInt(), moduleMage.Amount)
				}
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(sdk.ZeroInt(), moduleMage.Amount)
			}

			aMage := suite.Keeper.GetBalance(suite.Ctx, suite.Addrs[0])
			suite.Require().Equal(tt.expectedCoins.AmountOf("aMage"), aMage)
			umage := suite.BankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "umage")
			suite.Require().Equal(tt.expectedCoins.AmountOf("umage"), umage.Amount)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestConvertAMageToUMage() {
	tests := []struct {
		name          string
		startingCoins sdk.Coins
		expectedCoins sdk.Coins
	}{
		{
			"not enough umage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 100)),
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 100), sdk.NewInt64Coin("umage", 0)),
		},
		{
			"converts aMage for 1 umage",
			sdk.NewCoins(sdk.NewInt64Coin("umage", 10), sdk.NewInt64Coin("aMage", 1_000_000_000_003)),
			sdk.NewCoins(sdk.NewInt64Coin("umage", 11), sdk.NewInt64Coin("aMage", 3)),
		},
		{
			"converts more than 1 umage of aMage",
			sdk.NewCoins(sdk.NewInt64Coin("umage", 10), sdk.NewInt64Coin("aMage", 8_000_000_000_123)),
			sdk.NewCoins(sdk.NewInt64Coin("umage", 18), sdk.NewInt64Coin("aMage", 123)),
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupTest()

			err := suite.App.FundModuleAccount(suite.Ctx, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin("umage", 10)))
			suite.Require().NoError(err)
			suite.FundAccountWithMage(suite.Addrs[0], tt.startingCoins)
			err = suite.EvmBankKeeper.ConvertAMageToUMage(suite.Ctx, suite.Addrs[0])
			suite.Require().NoError(err)
			aMage := suite.Keeper.GetBalance(suite.Ctx, suite.Addrs[0])
			suite.Require().Equal(tt.expectedCoins.AmountOf("aMage"), aMage)
			umage := suite.BankKeeper.GetBalance(suite.Ctx, suite.Addrs[0], "umage")
			suite.Require().Equal(tt.expectedCoins.AmountOf("umage"), umage.Amount)
		})
	}
}

func (suite *evmBankKeeperTestSuite) TestSplitAMageCoins() {
	tests := []struct {
		name          string
		coins         sdk.Coins
		expectedCoins sdk.Coins
		shouldErr     bool
	}{
		{
			"invalid coins",
			sdk.NewCoins(sdk.NewInt64Coin("umage", 500)),
			nil,
			true,
		},
		{
			"empty coins",
			sdk.NewCoins(),
			sdk.NewCoins(),
			false,
		},
		{
			"umage & aMage coins",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 8_000_000_000_123)),
			sdk.NewCoins(sdk.NewInt64Coin("umage", 8), sdk.NewInt64Coin("aMage", 123)),
			false,
		},
		{
			"only aMage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 10_123)),
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 10_123)),
			false,
		},
		{
			"only umage",
			sdk.NewCoins(sdk.NewInt64Coin("aMage", 5_000_000_000_000)),
			sdk.NewCoins(sdk.NewInt64Coin("umage", 5)),
			false,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			umage, aMage, err := keeper.SplitAMageCoins(tt.coins)
			if tt.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tt.expectedCoins.AmountOf("umage"), umage.Amount)
				suite.Require().Equal(tt.expectedCoins.AmountOf("aMage"), aMage)
			}
		})
	}
}

func TestEvmBankKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(evmBankKeeperTestSuite))
}
