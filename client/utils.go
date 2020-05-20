package client

import (
	"github.com/binance-chain/bsc-double-sign-sdk/types/bsc"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func EthHeaderToBscHeader(ethHeader *ethtypes.Header) *bsc.Header {
	if ethHeader == nil {
		return nil
	}
	return &bsc.Header{
		ParentHash:  ethHeader.ParentHash,
		UncleHash:   ethHeader.UncleHash,
		Coinbase:    ethHeader.Coinbase,
		Root:        ethHeader.Root,
		TxHash:      ethHeader.TxHash,
		ReceiptHash: ethHeader.ReceiptHash,
		Bloom:       ethHeader.Bloom,
		Difficulty:  ethHeader.Difficulty.Int64(),
		Number:      ethHeader.Number.Int64(),
		GasLimit:    ethHeader.GasLimit,
		GasUsed:     ethHeader.GasUsed,
		Time:        ethHeader.Time,
		Extra:       ethHeader.Extra,
		MixDigest:   ethHeader.MixDigest,
		Nonce:       ethHeader.Nonce,
	}
}