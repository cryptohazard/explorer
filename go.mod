module github.com/cryptohazard/explorer

go 1.16

replace github.com/ethereum/go-ethereum/ethclient/1922 => github.com/ethereum/go-ethereum/ethclient v1.9.22
 
replace github.com/ethereum/go-ethereum/ethclient/11011 => github.com/ethereum/go-ethereum/ethclient v1.10.11

require (
	github.com/btcsuite/btcd v0.22.0-beta
	github.com/eoscanada/eos-go v0.10.0
	github.com/ethereum/go-ethereum v1.10.11
	github.com/mitchellh/mapstructure v1.4.2
)
