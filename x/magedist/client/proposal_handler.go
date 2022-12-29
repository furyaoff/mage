package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/furya-official/mage/x/magedist/client/cli"
	"github.com/furya-official/mage/x/magedist/client/rest"
)

// community-pool multi-spend proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
