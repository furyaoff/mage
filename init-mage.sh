validatorMnemonic="equip town gesture square tomorrow volume nephew minute witness beef rich gadget actress egg sing secret pole winter alarm law today check violin uncover"
faucetMnemonic="crash sort dwarf disease change advice attract clump avoid mobile clump right junior axis book fresh mask tube front require until face effort vault"
evmFaucetMnemonic="hundred flash cattle inquiry gorilla quick enact lazy galaxy apple bitter liberty print sun hurdle oak town cash because round chalk marriage response success"
userMnemonic="news tornado sponsor drastic dolphin awful plastic select true lizard width idle ability pigeon runway lift oppose isolate maple aspect safe jungle author hole"
relayerMnemonic="never reject sniff east arctic funny twin feed upper series stay shoot vivid adapt defense economy pledge fetch invite approve ceiling admit gloom exit"

DATA=~/.mage
rm -rf $DATA

BINARY=mage

chainID="magelocalnet_8888-1"
$BINARY init validator --chain-id $chainID

sed -in-place='' 's/enable = false/enable = true/g' $DATA/config/app.toml

sed -in-place='' 's/tracer = ""/tracer = "json"/g' $DATA/config/app.toml

sed -in-place='' '/iavl-cache-size/a\
trace = true' $DATA/config/app.toml

sed -in-place='' 's/chain-id = ""/chain-id = "magelocalnet_8888-1"/g' $DATA/config/client.toml

$BINARY config keyring-backend test

validatorKeyName="validator"
printf "$validatorMnemonic\n" | $BINARY keys add $validatorKeyName --recover
$BINARY add-genesis-account $validatorKeyName 2000000000umage,100000000000bnb

faucetKeyName="faucet"
printf "$faucetMnemonic\n" | $BINARY keys add $faucetKeyName --recover
$BINARY add-genesis-account $faucetKeyName 1000000000umage,100000000000bnb

evmFaucetKeyName="evm-faucet"
printf "$evmFaucetMnemonic\n" | $BINARY keys add $evmFaucetKeyName --eth --recover
$BINARY add-genesis-account $evmFaucetKeyName 1000000000umage

userKeyName="user"
printf "$userMnemonic\n" | $BINARY keys add $userKeyName --eth --recover
$BINARY add-genesis-account $userKeyName 1000000000umage,1000000000usdx

relayerKeyName="relayer"
printf "$relayerMnemonic\n" | $BINARY keys add $relayerKeyName --eth --recover
$BINARY add-genesis-account $relayerKeyName 1000000000umage

$BINARY gentx $validatorKeyName 1000000000umage --keyring-backend test --chain-id $chainID
$BINARY collect-gentxs

sed -in-place='' 's/stake/umage/g' $DATA/config/genesis.json

sed -in-place='' 's/aphoton/amage/g' $DATA/config/genesis.json

jq '.app_state.bank.supply = []' $DATA/config/genesis.json|sponge $DATA/config/genesis.json

jq '.app_state.feemarket.params.no_base_fee = true' $DATA/config/genesis.json|sponge $DATA/config/genesis.json

jq '.app_state.evm.params.chain_config.london_block = null' $DATA/config/genesis.json|sponge $DATA/config/genesis.json
jq '.app_state.evm.params.chain_config.arrow_glacier_block = null' $DATA/config/genesis.json|sponge $DATA/config/genesis.json
jq '.app_state.evm.params.chain_config.merge_fork_block = null' $DATA/config/genesis.json|sponge $DATA/config/genesis.json

jq '.app_state.earn.params.allowed_vaults =  [
    {
        denom: "usdx",
        strategies: ["STRATEGY_TYPE_HARD"],
    },
    {
        denom: "bmage",
        strategies: ["STRATEGY_TYPE_SAVINGS"],
    }]' $DATA/config/genesis.json | sponge $DATA/config/genesis.json

jq '.app_state.savings.params.supported_denoms = ["bmage-magevaloper1ffv7nhd3z6sych2qpqkk03ec6hzkmufyz4scd0"]' $DATA/config/genesis.json | sponge $DATA/config/genesis.json


$BINARY config broadcast-mode block