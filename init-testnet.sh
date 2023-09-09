#!/bin/sh

# set variables for the chain
VALIDATOR_NAME=validator1
CHAIN_ID=linux
KEY_NAME=linux-key
KEY_2_NAME=linux-key-2
CHAINFLAG="--chain-id ${CHAIN_ID}"
TOKEN_AMOUNT="10000000000000000000000000stake"
STAKING_AMOUNT="1000000000stake"

# create a random Namespace ID for your rollup to post blocks to
NAMESPACE_ID=$(openssl rand -hex 10)
echo $NAMESPACE_ID

# query the DA Layer start height, in this case we are querying
# an RPC endpoint provided by Celestia Labs. The RPC endpoint is
# to allow users to interact with Celestia's core network by querying
# the node's state and broadcasting transactions on the Celestia
# network. This is for Arabica, if using another network, change the RPC.
DA_BLOCK_HEIGHT=$(curl https://rpc-arabica-9.consensus.celestia-arabica.com/block |jq -r '.result.block.header.height')
echo $DA_BLOCK_HEIGHT

# build the gm chain with Rollkit
ignite chain build

cosmos-linuxd tendermint unsafe-reset-all
cosmos-linuxd init $VALIDATOR_NAME --chain-id $CHAIN_ID
cosmos-linuxd keys add $KEY_NAME --keyring-backend test
cosmos-linuxd keys add $KEY_2_NAME --keyring-backend test
cosmos-linuxd add-genesis-account $KEY_NAME $TOKEN_AMOUNT --keyring-backend test
cosmos-linuxd add-genesis-account $KEY_2_NAME $TOKEN_AMOUNT --keyring-backend test
cosmos-linuxd gentx $KEY_NAME $STAKING_AMOUNT --chain-id $CHAIN_ID --keyring-backend test
cosmos-linuxd collect-gentxs

# export the Celestia light node's auth token to allow you to submit
# PayForBlobs to Celestia's data availability network
# this is for Arabica, if using another network, change the network name
export AUTH_TOKEN=$(docker exec $(docker ps -q) celestia light auth write --p2p.network arabica)

# start the chain
cosmos-linuxd start --rollkit.aggregator true --rollkit.da_layer celestia --rollkit.da_config='{"base_url":"http://localhost:26658","timeout":60000000000,"fee":600000,"gas_limit":6000000,"auth_token":"'$AUTH_TOKEN'"}' --rollkit.namespace_id $NAMESPACE_ID --rollkit.da_start_height $DA_BLOCK_HEIGHT
