package tx

import (
	"strconv"
	"kchain/types/cnst"
)


// Account
type Account struct {
	PubKey string `json:"pubkey,omitempty" mapstructure:"pubkey"`
	Power  int64  `json:"power,omitempty" mapstructure:"power"`
}

func (act *Account)ToSortBytes() []byte {
	return []byte(strconv.Itoa(act.Power) + act.PubKey)
}
func (act *Account)ToBytes() ([]byte, error) {
	return json.Marshal(act)
}
func (act *Account) FromBytes(d []byte) error {
	return json.Unmarshal(d, act)
}
func (act *Account)GetPrefix() string {
	return cnst.AccountPrefix
}
func (act *Account) Key() []byte {
	return []byte(act.GetPrefix() + act.Key)
}
func (act *Account) ToKv() ([]byte, []byte) {
	return act.Key(), act.ToBytes()
}



