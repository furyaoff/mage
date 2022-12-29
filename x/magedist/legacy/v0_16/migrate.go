package v0_16

import (
	v015Magedist "github.com/furya-official/mage/x/Magedist/legacy/v0_15"
	v016Magedist "github.com/furya-official/mage/x/Magedist/types"
)

func migrateParams(oldParams v015Magedist.Params) v016Magedist.Params {
	periods := make([]v016Magedist.Period, len(oldParams.Periods))
	for i, oldPeriod := range oldParams.Periods {
		periods[i] = v016Magedist.Period{
			Start:     oldPeriod.Start,
			End:       oldPeriod.End,
			Inflation: oldPeriod.Inflation,
		}
	}
	return v016Magedist.Params{
		Periods: periods,
		Active:  oldParams.Active,
	}
}

// Migrate converts v0.15 Magedist state and returns it in v0.16 format
func Migrate(oldState v015Magedist.GenesisState) *v016Magedist.GenesisState {
	return &v016Magedist.GenesisState{
		Params:            migrateParams(oldState.Params),
		PreviousBlockTime: oldState.PreviousBlockTime,
	}
}
