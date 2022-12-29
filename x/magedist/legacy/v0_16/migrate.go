package v0_16

import (
	v015magedist "github.com/furya-official/mage/x/magedist/legacy/v0_15"
	v016magedist "github.com/furya-official/mage/x/magedist/types"
)

func migrateParams(oldParams v015magedist.Params) v016magedist.Params {
	periods := make([]v016magedist.Period, len(oldParams.Periods))
	for i, oldPeriod := range oldParams.Periods {
		periods[i] = v016magedist.Period{
			Start:     oldPeriod.Start,
			End:       oldPeriod.End,
			Inflation: oldPeriod.Inflation,
		}
	}
	return v016magedist.Params{
		Periods: periods,
		Active:  oldParams.Active,
	}
}

// Migrate converts v0.15 magedist state and returns it in v0.16 format
func Migrate(oldState v015magedist.GenesisState) *v016magedist.GenesisState {
	return &v016magedist.GenesisState{
		Params:            migrateParams(oldState.Params),
		PreviousBlockTime: oldState.PreviousBlockTime,
	}
}
