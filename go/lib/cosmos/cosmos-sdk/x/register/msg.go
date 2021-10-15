package register

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
	crypto "github.com/cosmos/amino-js/go/lib/tendermint/tendermint/crypto"
)

type MsgCreateResourceNode struct {
	NetworkID    string         `json:"network_id" yaml:"network_id"`
	PubKey       crypto.PubKey  `json:"pubkey" yaml:"pubkey"`
	Value        sdk.Coin       `json:"value" yaml:"value"`
	OwnerAddress sdk.AccAddress `json:"owner_address" yaml:"owner_address"`
	Description  Description    `json:"description" yaml:"description"`
	NodeType     string         `json:"node_type" yaml:"node_type"`
}

type MsgRemoveResourceNode struct {
	ResourceNodeAddress sdk.AccAddress `json:"resource_node_address" yaml:"resource_node_address"`
	OwnerAddress        sdk.AccAddress `json:"owner_address" yaml:"owner_address"`
}
type MsgCreateIndexingNode struct {
	NetworkID string        `json:"network_id" yaml:"network_id"`
	PubKey    crypto.PubKey `json:"pubkey" yaml:"pubkey"`
	//NetworkAddress sdk.AccAddress `json:"network_address" yaml:"network_address"`
	Value        sdk.Coin       `json:"value" yaml:"value"`
	OwnerAddress sdk.AccAddress `json:"owner_address" yaml:"owner_address"`
	Description  Description    `json:"description" yaml:"description"`
}

type MsgRemoveIndexingNode struct {
	IndexingNodeAddress sdk.AccAddress `json:"indexing_node_address" yaml:"indexing_node_address"`
	OwnerAddress        sdk.AccAddress `json:"owner_address" yaml:"owner_address"`
}
type MsgIndexingNodeRegistrationVote struct {
	NodeAddress  sdk.AccAddress `json:"node_address" yaml:"node_address"`   // node address of indexing node
	OwnerAddress sdk.AccAddress `json:"owner_address" yaml:"owner_address"` // owner address of indexing node
	Opinion      VoteOpinion    `json:"opinion" yaml:"opinion"`
	VoterAddress sdk.AccAddress `json:"voter_address" yaml:"voter_address"` // address of voter (other existed indexing node)
}
var (
	_ sdk.Msg = &MsgCreateResourceNode{}
	_ sdk.Msg = &MsgRemoveResourceNode{}
	_ sdk.Msg = &MsgCreateIndexingNode{}
	_ sdk.Msg = &MsgRemoveIndexingNode{}
	_ sdk.Msg = &MsgIndexingNodeRegistrationVote{}
)


type VoteOpinion bool

type Description struct {
	Moniker         string `json:"moniker" yaml:"moniker"`                   // name
	Identity        string `json:"identity" yaml:"identity"`                 // optional identity signature (ex. UPort or Keybase)
	Website         string `json:"website" yaml:"website"`                   // optional website link
	SecurityContact string `json:"security_contact" yaml:"security_contact"` // optional security contact info
	Details         string `json:"details" yaml:"details"`                   // optional details
}
