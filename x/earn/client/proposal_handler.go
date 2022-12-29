package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/furya-official/mage/x/earn/client/cli"
	"github.com/furya-official/mage/x/earn/client/rest"
)

// community-pool deposit/withdraw proposal handlers
var (
	DepositProposalHandler  = govclient.NewProposalHandler(cli.GetCmdSubmitCommunityPoolDepositProposal, rest.DepositProposalRESTHandler)
	WithdrawProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitCommunityPoolWithdrawProposal, rest.WithdrawProposalRESTHandler)
)
