package testutil

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furya-official/mage/app"

	mageminttypes "github.com/furya-official/mage/x/magemint/types"
)

// magemintGenesisBuilder is a tool for creating a mint genesis state.
// Helper methods add values onto a default genesis state.
// All methods are immutable and return updated copies of the builder.
type magemintGenesisBuilder struct {
	mageminttypes.GenesisState
}

var _ GenesisBuilder = (*magemintGenesisBuilder)(nil)

func NewmagemintGenesisBuilder() magemintGenesisBuilder {
	gen := mageminttypes.DefaultGenesisState()
	gen.Params.CommunityPoolInflation = sdk.ZeroDec()
	gen.Params.StakingRewardsApy = sdk.ZeroDec()

	return magemintGenesisBuilder{
		GenesisState: *gen,
	}
}

func (builder magemintGenesisBuilder) Build() mageminttypes.GenesisState {
	return builder.GenesisState
}

func (builder magemintGenesisBuilder) BuildMarshalled(cdc codec.JSONCodec) app.GenesisState {
	built := builder.Build()

	return app.GenesisState{
		mageminttypes.ModuleName: cdc.MustMarshalJSON(&built),
	}
}

func (builder magemintGenesisBuilder) WithPreviousBlockTime(t time.Time) magemintGenesisBuilder {
	builder.PreviousBlockTime = t
	return builder
}

func (builder magemintGenesisBuilder) WithStakingRewardsApy(apy sdk.Dec) magemintGenesisBuilder {
	builder.Params.StakingRewardsApy = apy
	return builder
}

func (builder magemintGenesisBuilder) WithCommunityPoolInflation(
	inflation sdk.Dec,
) magemintGenesisBuilder {
	builder.Params.CommunityPoolInflation = inflation
	return builder
}
