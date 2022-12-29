package testutil

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furya-official/mage/app"

	Mageminttypes "github.com/furya-official/mage/x/Magemint/types"
)

// MagemintGenesisBuilder is a tool for creating a mint genesis state.
// Helper methods add values onto a default genesis state.
// All methods are immutable and return updated copies of the builder.
type MagemintGenesisBuilder struct {
	Mageminttypes.GenesisState
}

var _ GenesisBuilder = (*MagemintGenesisBuilder)(nil)

func NewMagemintGenesisBuilder() MagemintGenesisBuilder {
	gen := Mageminttypes.DefaultGenesisState()
	gen.Params.CommunityPoolInflation = sdk.ZeroDec()
	gen.Params.StakingRewardsApy = sdk.ZeroDec()

	return MagemintGenesisBuilder{
		GenesisState: *gen,
	}
}

func (builder MagemintGenesisBuilder) Build() Mageminttypes.GenesisState {
	return builder.GenesisState
}

func (builder MagemintGenesisBuilder) BuildMarshalled(cdc codec.JSONCodec) app.GenesisState {
	built := builder.Build()

	return app.GenesisState{
		Mageminttypes.ModuleName: cdc.MustMarshalJSON(&built),
	}
}

func (builder MagemintGenesisBuilder) WithPreviousBlockTime(t time.Time) MagemintGenesisBuilder {
	builder.PreviousBlockTime = t
	return builder
}

func (builder MagemintGenesisBuilder) WithStakingRewardsApy(apy sdk.Dec) MagemintGenesisBuilder {
	builder.Params.StakingRewardsApy = apy
	return builder
}

func (builder MagemintGenesisBuilder) WithCommunityPoolInflation(
	inflation sdk.Dec,
) MagemintGenesisBuilder {
	builder.Params.CommunityPoolInflation = inflation
	return builder
}
