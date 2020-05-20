package client

import (
	"github.com/binance-chain/bsc-double-sign-sdk/types/bsc"
	"github.com/binance-chain/bsc-double-sign-sdk/types/msg"
	"github.com/binance-chain/go-sdk/client/rpc"
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/types/tx"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

func BSCSubmitEvidence(c *rpc.HTTP, submitter types.AccAddress, headers []*bsc.Header,
	syncType rpc.SyncType, options ...tx.Option) (*coretypes.ResultBroadcastTx, error) {
	m := msg.NewMsgBscSubmitEvidence(submitter, headers)
	return c.Broadcast(m, syncType, options...)
}
