package tx

import (
	"strconv"
	"kchain/types/cnst"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Db
type Db struct {
	Key   string `json:"key,omitempty" mapstructure:"key"`
	Value string `json:"value,omitempty" mapstructure:"value"`
}

func (db *Db)ToSortBytes() []byte {
	return []byte(db.Key + db.Value)
}
func (db *Db)ToBytes() ([]byte, error) {
	return json.Marshal(db)
}
func (db *Db) FromBytes(d []byte) error {
	return json.Unmarshal(d, db)
}
func (db *Db)ToKv() ([]byte, []byte) {
	return db.Key(), db.ToBytes()
}
func (db *Db)Key() []byte {
	return []byte(db.GetPrefix() + db.Key)
}

func (db *Db)GetPrefix() string {
	return cnst.DbPrefix
}

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



