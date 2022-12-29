<!--
order: 2
-->

# State

## Parameters and Genesis State

`Parameters` define the list of conversion pairs allowed to be converted between Mage ERC20 tokens and sdk.Coins.

```protobuf
// Params defines the evmutil module params
message Params {
  // enabled_conversion_pairs defines the list of conversion pairs allowed to be
  // converted between Mage ERC20 and sdk.Coin
  repeated ConversionPair enabled_conversion_pairs = 4;
}

// ConversionPair defines a Mage ERC20 address and corresponding denom that is
// allowed to be converted between ERC20 and sdk.Coin
message ConversionPair {
  // ERC20 address of the token on the Mage EVM
  bytes Mage_erc20_address = 1;
  // Denom of the corresponding sdk.Coin
  string denom = 2;
}
```

`GenesisState` defines the state that must be persisted when the blockchain stops/restarts in order for normal function of the evmutil module to resume.

```protobuf
message GenesisState {
  repeated Account accounts = 1 [(gogoproto.nullable) = false];
  Params params = 2 [(gogoproto.nullable) = false];
}
```

## Account

An `Account` is a struct representing the excess `aMage` balance of an address.

Since an address's total `aMage` balance is derived from its `umage` balance and the excess `aMage` balance stored by the `Account` struct, the `aMage` balance here should never exceed 1 `umage` (10^12 `aMage`).

```protobuf
message Account {
  bytes address = 1;
  string balance = 2;
}
```

## Store

For complete implementation details for how items are stored, see [keys.go](../types/keys.go). `x/evmutil` store state consists of accounts.
