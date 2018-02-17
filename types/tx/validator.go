package tx

import (
	"strconv"
	"kchain/types/cnst"
)


// Validator
type Validator struct {
	PubKey string `json:"pubkey,omitempty" mapstructure:"pubkey"`
	Power  int64  `json:"power,omitempty" mapstructure:"power"`
}

func (val *Validator)ToSortBytes() []byte {
	return []byte(strconv.Itoa(val.Power) + val.PubKey)
}
func (val *Validator)ToBytes() ([]byte, error) {
	return json.Marshal(val)
}
func (val *Validator) FromBytes(d []byte) error {
	return json.Unmarshal(d, val)
}
func (val *Validator)GetPrefix() string {
	return cnst.ValidatorPrefix
}

func (val *Validator)Key() []byte {
	return []byte(val.GetPrefix() + val.Key)
}

func (val *Validator) ToKv() ([]byte, []byte) {
	return val.Key(), val.ToBytes()
}
