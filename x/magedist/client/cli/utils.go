package cli

import (
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/furya-official/mage/x/magedist/types"
)

// ParseCommunityPoolMultiSpendProposalJSON reads and parses a CommunityPoolMultiSpendProposalJSON from a file.
func ParseCommunityPoolMultiSpendProposalJSON(cdc codec.JSONCodec, proposalFile string) (types.CommunityPoolMultiSpendProposalJSON, error) {
	proposal := types.CommunityPoolMultiSpendProposalJSON{}
	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}
