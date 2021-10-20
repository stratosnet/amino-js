package pot

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"

)

type MsgVolumeReport struct {
	NodesVolume     []SingleNodeVolume `json:"nodes_volume" yaml:"nodes_volume"`         // volume report
	Reporter        sdk.AccAddress     `json:"reporter" yaml:"reporter"`                 // volume reporter
	Epoch           sdk.Int            `json:"report_epoch" yaml:"report_epoch"`         // volume report epoch
	ReportReference string             `json:"report_reference" yaml:"report_reference"` // volume report reference
}

type MsgWithdraw struct {
	Amount       sdk.Coin       `json:"amount" yaml:"amount"`
	NodeAddress  sdk.AccAddress `json:"node_address" yaml:"node_address"`
	OwnerAddress sdk.AccAddress `json:"owner_address" yaml:"owner_address"`
}

type MsgFoundationDeposit struct {
	Amount       sdk.Coin       `json:"amount" yaml:"amount"`
	
	From         sdk.AccAddress `json:"from" yaml:"from"`
}


var (
	_ sdk.Msg = &MsgVolumeReport{}
	_ sdk.Msg = &MsgWithdraw{}
	_ sdk.Msg = &MsgFoundationDeposit{}
)


type SingleNodeVolume struct {
	NodeAddress sdk.AccAddress `json:"node_address" yaml:"node_address"`
	Volume      sdk.Int        `json:"node_volume" yaml:"node_volume"`
}
