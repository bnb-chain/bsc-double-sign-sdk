package msg

import (
	"github.com/binance-chain/go-sdk/types/msg"
	"github.com/binance-chain/go-sdk/types/tx"
	"github.com/tendermint/go-amino"
)

func RegisterCodec(cdc *amino.Codec) {
	cdc.RegisterConcrete(MsgBscSubmitEvidence{}, "cosmos-sdk/MsgBscSubmitEvidence", nil)
}

func init() {
	RegisterCodec(msg.MsgCdc)
	RegisterCodec(tx.Cdc)
}
