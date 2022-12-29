<!--
order: 2
-->

# State

## Parameters and Genesis State

`Parameters` define the governance parameters that control inflation of MAGE.
Each parameter is an APY/inflation used in conjunction with the block time to
calculate how much MAGE to mint for what purposes. See [params](./05_params.md) for a description of each param.

```go
// Params wraps the governance parameters for the magemint module
type Params struct {
	// yearly inflation of total token supply minted to the community pool.
	CommunityPoolInflation sdk.Dec `json:"community_pool_inflation" yaml:"community_pool_inflation"`
	// yearly inflation of bonded tokens minted for staking rewards to validators.
	StakingRewardsApy sdk.Dec `json:"staking_rewards_apy" yaml:"staking_rewards_apy"`
}
```

`GenesisState` defines the state that must be persisted when the blockchain stops/restarts in order for normal function of the magemint module to resume.

```go

// GenesisState defines the magemint module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `json:"params" yaml:"params"`
}
```

## Previous Block Time

The `PreviousBlockTime` is stored in the keeper and updated each block. In the BeginBlocker, the number of seconds between the current and previous block is determined and used to calculate how much MAGE should be minted for each
