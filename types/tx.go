package types

import (
	"github.com/tendermint/iavl"
	ktx "kchain/types/tx"
)

type Transaction struct {
	SignPubKey string        `json:"pubkey,omitempty"`
	Signature  string        `json:"sign,omitempty"`
	Data       interface{}   `json:"data,omitempty"`
	State      *iavl.VersionedTree
	db         *ktx.Db
	account    *ktx.Account
	validator  *ktx.Validator
}
