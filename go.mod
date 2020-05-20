module github.com/binance-chain/bsc-double-sign-sdk

go 1.13

require (
	github.com/binance-chain/go-sdk v1.2.2-0.20200520064040-fc067ad70353
	github.com/ethereum/go-ethereum v1.9.13
	github.com/tendermint/go-amino v0.14.1
	github.com/tendermint/tendermint v0.32.3
	golang.org/x/crypto v0.0.0-20200311171314-f7b00557c8c4
)

replace github.com/zondax/hid => github.com/binance-chain/hid v0.9.1-0.20190807012304-e1ffd6f0a3cc
