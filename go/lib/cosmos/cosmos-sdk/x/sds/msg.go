package sds

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"

)

type MsgFileUpload struct {
	FileHash []byte         `json:"file_hash" yaml:"file_hash"` // hash of file
	Reporter sdk.AccAddress `json:"reporter" yaml:"reporter"`   // reporter of tx
}

var _ sdk.Msg = &MsgFileUpload{}


type MsgPrepay struct {
	Sender sdk.AccAddress `json:"sender" yaml:"sender"` // sender of tx
	Coins  sdk.Coins      `json:"coins" yaml:"coins"`   // coins to send
}

var _ sdk.Msg = &MsgPrepay{}