package msg

import (
	"bytes"
	"fmt"

	"github.com/binance-chain/bsc-double-sign-sdk/types/bsc"
	"github.com/binance-chain/go-sdk/common/types"
	gosdkmsg "github.com/binance-chain/go-sdk/types/msg"
)

const (
	TypeSideChainSubmitEvidence = "bsc_submit_evidence"

	SideChainSlashMsgRoute = "slashing"
)

type MsgBscSubmitEvidence struct {
	Submitter types.AccAddress `json:"submitter"`
	Headers   []*bsc.Header    `json:"headers"`
}

func NewMsgBscSubmitEvidence(submitter types.AccAddress, headers []*bsc.Header) MsgBscSubmitEvidence {

	return MsgBscSubmitEvidence{
		Submitter: submitter,
		Headers:   headers,
	}
}

func (MsgBscSubmitEvidence) Route() string {
	return SideChainSlashMsgRoute
}

func (MsgBscSubmitEvidence) Type() string {
	return TypeSideChainSubmitEvidence
}

func headerEmptyCheck(header *bsc.Header) error {
	if header.Number == 0 {
		return fmt.Errorf("header number can not be zero ")
	}
	if header.Difficulty == 0 {
		return fmt.Errorf("header difficulty can not be zero")
	}
	if header.Extra == nil {
		return fmt.Errorf("header extra can not be empty")
	}

	return nil
}

func (msg MsgBscSubmitEvidence) ValidateBasic() error {
	if len(msg.Submitter) != types.AddrLen {
		return fmt.Errorf("Expected delegator address length is %d, actual length is %d", types.AddrLen, len(msg.Submitter))
	}

	if err := headerEmptyCheck(msg.Headers[0]); err != nil {
		return err
	}
	if err := headerEmptyCheck(msg.Headers[1]); err != nil {
		return err
	}
	if msg.Headers[0].Number != msg.Headers[1].Number {
		return fmt.Errorf("The numbers of two block headers are not the same")
	}
	if !bytes.Equal(msg.Headers[0].ParentHash[:], msg.Headers[1].ParentHash[:]) {
		return fmt.Errorf("The parent hash of two block headers are not the same")
	}
	signature1, err := msg.Headers[0].GetSignature()
	if err != nil {
		return fmt.Errorf("Failed to get signature from block header, %s", err.Error())
	}
	signature2, err := msg.Headers[1].GetSignature()
	if err != nil {
		return fmt.Errorf("Failed to get signature from block header, %s", err.Error())
	}
	if bytes.Compare(signature1, signature2) == 0 {
		return fmt.Errorf("The two blocks are the same")
	}

	return nil
}

func (msg MsgBscSubmitEvidence) GetSignBytes() []byte {
	bz := gosdkmsg.MsgCdc.MustMarshalJSON(msg)
	return gosdkmsg.MustSortJSON(bz)
}

func (msg MsgBscSubmitEvidence) GetSigners() []types.AccAddress {
	return []types.AccAddress{msg.Submitter}
}

func (msg MsgBscSubmitEvidence) GetInvolvedAddresses() []types.AccAddress {
	return msg.GetSigners()
}
