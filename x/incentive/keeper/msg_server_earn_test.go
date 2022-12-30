package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/x/distribution"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	earntypes "github.com/furya-official/mage/x/earn/types"
	"github.com/furya-official/mage/x/incentive"
	"github.com/furya-official/mage/x/incentive/testutil"
	"github.com/furya-official/mage/x/incentive/types"
	"github.com/furya-official/mage/x/Magemint"
	liquidtypes "github.com/furya-official/mage/x/liquid/types"
)

func (suite *HandlerTestSuite) TestEarnLiquidClaim() {
	userAddr1, userAddr2, validatorAddr1, validatorAddr2 := suite.addrs[0], suite.addrs[1], suite.addrs[2], suite.addrs[3]

	valAddr1 := sdk.ValAddress(validatorAddr1)
	valAddr2 := sdk.ValAddress(validatorAddr2)

	authBuilder := suite.authBuilder().
		WithSimpleAccount(userAddr1, cs(c("umage", 1e12))).
		WithSimpleAccount(userAddr2, cs(c("umage", 1e12))).
		WithSimpleAccount(validatorAddr1, cs(c("umage", 1e12))).
		WithSimpleAccount(validatorAddr2, cs(c("umage", 1e12)))

	incentBuilder := suite.incentiveBuilder().
		WithSimpleEarnRewardPeriod("bmage", cs())

	savingsBuilder := testutil.NewSavingsGenesisBuilder().
		WithSupportedDenoms("bmage")

	earnBuilder := testutil.NewEarnGenesisBuilder().
		WithAllowedVaults(earntypes.AllowedVault{
			Denom:             "bmage",
			Strategies:        earntypes.StrategyTypes{earntypes.STRATEGY_TYPE_SAVINGS},
			IsPrivateVault:    false,
			AllowedDepositors: nil,
		})

	suite.SetupWithGenState(
		authBuilder,
		incentBuilder,
		earnBuilder,
		savingsBuilder,
	)

	// ak := suite.App.GetAccountKeeper()
	// bk := suite.App.GetBankKeeper()
	sk := suite.App.GetStakingKeeper()
	lq := suite.App.GetLiquidKeeper()
	mk := suite.App.GetMagemintKeeper()
	dk := suite.App.GetDistrKeeper()
	ik := suite.App.GetIncentiveKeeper()

	iParams := ik.GetParams(suite.Ctx)
	period, found := iParams.EarnRewardPeriods.GetMultiRewardPeriod("bmage")
	suite.Require().True(found)
	suite.Require().Equal("bmage", period.CollateralType)

	// Ensure nonzero staking APY
	mParams := mk.GetParams(suite.Ctx)
	mParams.StakingRewardsApy = sdk.NewDecWithPrec(20, 2)
	mk.SetParams(suite.Ctx, mParams)

	bmageDenom1 := lq.GetLiquidStakingTokenDenom(valAddr1)
	bmageDenom2 := lq.GetLiquidStakingTokenDenom(valAddr2)

	err := suite.App.FundModuleAccount(suite.Ctx, distrtypes.ModuleName, cs(c("umage", 1e12)))
	suite.NoError(err)

	// Create validators
	err = suite.DeliverMsgCreateValidator(valAddr1, c("umage", 1e9))
	suite.Require().NoError(err)

	err = suite.DeliverMsgCreateValidator(valAddr2, c("umage", 1e9))
	suite.Require().NoError(err)

	// new block required to bond validator
	suite.NextBlockAfter(7 * time.Second)
	// Now the delegation is bonded, accumulate some delegator rewards
	suite.NextBlockAfter(7 * time.Second)

	// Create delegations from users
	// User 1: 1e9 umage to validator 1
	// User 2: 99e9 umage to validator 1 AND 2
	err = suite.DeliverMsgDelegate(userAddr1, valAddr1, c("umage", 1e9))
	suite.Require().NoError(err)

	err = suite.DeliverMsgDelegate(userAddr2, valAddr1, c("umage", 99e9))
	suite.Require().NoError(err)

	err = suite.DeliverMsgDelegate(userAddr2, valAddr2, c("umage", 99e9))
	suite.Require().NoError(err)

	// Mint liquid tokens
	_, err = suite.DeliverMsgMintDerivative(userAddr1, valAddr1, c("umage", 1e9))
	suite.Require().NoError(err)

	_, err = suite.DeliverMsgMintDerivative(userAddr2, valAddr1, c("umage", 99e9))
	suite.Require().NoError(err)

	_, err = suite.DeliverMsgMintDerivative(userAddr2, valAddr2, c("umage", 99e9))
	suite.Require().NoError(err)

	// Deposit liquid tokens to earn
	err = suite.DeliverEarnMsgDeposit(userAddr1, c(bmageDenom1, 1e9), earntypes.STRATEGY_TYPE_SAVINGS)
	suite.Require().NoError(err)

	err = suite.DeliverEarnMsgDeposit(userAddr2, c(bmageDenom1, 99e9), earntypes.STRATEGY_TYPE_SAVINGS)
	suite.Require().NoError(err)
	err = suite.DeliverEarnMsgDeposit(userAddr2, c(bmageDenom2, 99e9), earntypes.STRATEGY_TYPE_SAVINGS)
	suite.Require().NoError(err)

	// BeginBlocker to update minter annual provisions as it starts at 0 which results in no minted coins
	_ = suite.App.BeginBlocker(suite.Ctx, abci.RequestBeginBlock{})

	// DeliverMsgCreateValidator uses a generated pubkey, so we need to fetch
	// the validator to get the correct pubkey
	validator1, found := sk.GetValidator(suite.Ctx, valAddr1)
	suite.Require().True(found)

	pk, err := validator1.ConsPubKey()
	suite.Require().NoError(err)

	val := abci.Validator{
		Address: pk.Address(),
		Power:   100,
	}

	// Query for next block to get staking rewards
	suite.Ctx = suite.Ctx.
		WithBlockHeight(suite.Ctx.BlockHeight() + 1).
		WithBlockTime(suite.Ctx.BlockTime().Add(7 * time.Second))

	// Mint tokens
	Magemint.BeginBlocker(suite.Ctx, suite.App.GetMagemintKeeper())
	// Distribute to validators, block needs votes
	distribution.BeginBlocker(
		suite.Ctx,
		abci.RequestBeginBlock{
			LastCommitInfo: abci.LastCommitInfo{
				Votes: []abci.VoteInfo{{
					Validator:       val,
					SignedLastBlock: true,
				}},
			},
		},
		dk,
	)

	liquidMacc := suite.App.GetAccountKeeper().GetModuleAccount(suite.Ctx, liquidtypes.ModuleAccountName)
	delegation, found := sk.GetDelegation(suite.Ctx, liquidMacc.GetAddress(), valAddr1)
	suite.Require().True(found)

	// Get amount of rewards
	endingPeriod := dk.IncrementValidatorPeriod(suite.Ctx, validator1)

	// Zero rewards since this block is the same as the block it was last claimed

	// This needs to run **after** staking rewards are minted/distributed in
	// x/Magemint + x/distribution but **before** the x/incentive BeginBlocker.

	// Order of operations:
	// 1. x/Magemint + x/distribution BeginBlocker
	// 2. CalculateDelegationRewards
	// 3. x/incentive BeginBlocker to claim staking rewards
	delegationRewards := dk.CalculateDelegationRewards(suite.Ctx, validator1, delegation, endingPeriod)
	suite.Require().False(delegationRewards.IsZero(), "expected non-zero delegation rewards")

	// Claim staking rewards via incentive.
	// Block height was updated earlier.
	incentive.BeginBlocker(
		suite.Ctx,
		ik,
	)

	preClaimBal1 := suite.GetBalance(userAddr1)
	preClaimBal2 := suite.GetBalance(userAddr2)

	// Claim umage staking rewards
	denomsToClaim := map[string]string{"umage": "large"}
	selections := types.NewSelectionsFromMap(denomsToClaim)

	msg1 := types.NewMsgClaimEarnReward(userAddr1.String(), selections)
	msg2 := types.NewMsgClaimEarnReward(userAddr2.String(), selections)

	err = suite.DeliverIncentiveMsg(&msg1)
	suite.Require().NoError(err)

	err = suite.DeliverIncentiveMsg(&msg2)
	suite.Require().NoError(err)

	// Check rewards were paid out
	// User 1 gets 1% of rewards
	// User 2 gets 99% of rewards
	stakingRewards1 := delegationRewards.
		AmountOf("umage").
		Quo(sdk.NewDec(100)).
		RoundInt()
	suite.BalanceEquals(userAddr1, preClaimBal1.Add(sdk.NewCoin("umage", stakingRewards1)))

	// Total * 99 / 100
	stakingRewards2 := delegationRewards.
		AmountOf("umage").
		Mul(sdk.NewDec(99)).
		Quo(sdk.NewDec(100)).
		TruncateInt()
	suite.BalanceEquals(userAddr2, preClaimBal2.Add(sdk.NewCoin("umage", stakingRewards2)))

	suite.Equal(delegationRewards.AmountOf("umage").TruncateInt(), stakingRewards1.Add(stakingRewards2))

	// Check that claimed coins have been removed from a claim's reward
	suite.EarnRewardEquals(userAddr1, cs())
	suite.EarnRewardEquals(userAddr2, cs())
}
