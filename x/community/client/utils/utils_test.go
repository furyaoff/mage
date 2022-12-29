package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/furya-official/mage/x/community/client/utils"
)

func TestParseDepositProposal(t *testing.T) {
	cdc := codec.NewAminoCodec(codec.NewLegacyAmino())
	okJSON := testutil.WriteToNewTempFile(t, `
{
  "title": "Community Pool Lend Deposit",
  "description": "Deposit some MAGE from community pool to Lend!",
  "amount": [
    {
      "denom": "uMage",
      "amount": "100000000000"
    }
  ]
}
`)
	proposal, err := utils.ParseCommunityPoolLendDepositProposal(cdc, okJSON.Name())
	require.NoError(t, err)

	expectedAmount, err := sdk.ParseCoinsNormalized("100000000000uMage")
	require.NoError(t, err)

	require.Equal(t, "Community Pool Lend Deposit", proposal.Title)
	require.Equal(t, "Deposit some MAGE from community pool to Lend!", proposal.Description)
	require.Equal(t, expectedAmount, proposal.Amount)
}

func TestParseWithdrawProposal(t *testing.T) {
	cdc := codec.NewAminoCodec(codec.NewLegacyAmino())
	okJSON := testutil.WriteToNewTempFile(t, `
{
  "title": "Community Pool Lend Withdraw",
  "description": "Withdraw some MAGE from community pool to Lend!",
  "amount": [
    {
      "denom": "uMage",
      "amount": "100000000000"
    }
  ]
}
`)
	proposal, err := utils.ParseCommunityPoolLendWithdrawProposal(cdc, okJSON.Name())
	require.NoError(t, err)

	expectedAmount, err := sdk.ParseCoinsNormalized("100000000000uMage")
	require.NoError(t, err)

	require.Equal(t, "Community Pool Lend Withdraw", proposal.Title)
	require.Equal(t, "Withdraw some MAGE from community pool to Lend!", proposal.Description)
	require.Equal(t, expectedAmount, proposal.Amount)
}
