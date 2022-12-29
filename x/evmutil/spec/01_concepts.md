<!--
order: 1
-->

# Concepts

## EVM Gas Denom

In order to use the EVM and be compatible with existing clients, the gas denom used by the EVM must be in 18 decimals. Since `umage` has 6 decimals of precision, it cannot be used as the EVM gas denom directly.

To use the mage token on the EVM, the evmutil module provides an `EvmBankKeeper` that is responsible for the conversion of `umage` and `amage`. A user's excess `amage` balance is stored in the `x/evmutil` store, while its `umage` balance remains in the cosmos-sdk `x/bank` module.

## `EvmBankKeeper` Overview

The `EvmBankKeeper` provides access to an account's total `amage` balance and the ability to transfer, mint, and burn `amage`. If anything other than the `amage` denom is requested, the `EvmBankKeeper` will panic.

This keeper implements the `x/evm` module's `BankKeeper` interface to enable the usage of `amage` denom on the EVM.

### `x/evm` Parameter Requirements

Since the EVM denom `amage` is required to use the `EvmBankKeeper`, it is necessary to set the `EVMDenom` param of the `x/evm` module to `amage`.

### Balance Calculation of `amage`

The `amage` balance of an account is derived from an account's **spendable** `umage` balance times 10^12 (to derive its `amage` equivalent), plus the account's excess `amage` balance that can be accessed via the module `Keeper`.

### `amage` <> `umage` Conversion

When an account does not have sufficient `amage` to cover a transfer or burn, the `EvmBankKeeper` will try to swap 1 `umage` to its equivalent `amage` amount. It does this by transferring 1 `umage` from the sender to the `x/evmutil` module account, then adding the equivalent `amage` amount to the sender's balance in the module state.

In reverse, if an account has enough `amage` balance for one or more `umage`, the excess `amage` balance will be converted to `umage`. This is done by removing the excess `amage` balance in the module store, then transferring the equivalent `umage` coins from the `x/evmutil` module account to the target account.

The swap logic ensures that all `amage` is backed by the equivalent `umage` balance stored in the module account.

## ERC20 token <> sdk.Coin Conversion

`x/evmutil` enables the conversion between ERC20 tokens and sdk.Coins. This done through the use of the `MsgConvertERC20ToCoin` & `MsgConvertCoinToERC20` messages (see **[Messages](03_messages.md)**).

Only ERC20 contract address that are whitelist via the `EnabledConversionPairs` param (see **[Params](05_params.md)**) can be converted via these messages.

## Module Keeper

The module Keeper provides access to an account's excess `amage` balance and the ability to update the balance.
