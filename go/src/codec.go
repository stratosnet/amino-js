package src

import (
	// "github.com/cosmos/amino-js/go/lib"
	"../lib"
 amino	"github.com/tendermint/go-amino"
)

var codec *amino.Codec

func init() {
	codec = amino.NewCodec()
	lib.RegisterCodec(codec)
	codec.Seal()
}
