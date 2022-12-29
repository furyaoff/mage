package Magedist

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/furya-official/mage/x/Magedist/keeper"
	"github.com/furya-official/mage/x/Magedist/types"
)

// NewCommunityPoolMultiSpendProposalHandler
func NewCommunityPoolMultiSpendProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.CommunityPoolMultiSpendProposal:
			return keeper.HandleCommunityPoolMultiSpendProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized Magedist proposal content type: %T", c)
		}
	}
}
