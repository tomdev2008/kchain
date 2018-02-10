package transaction

import "strconv"

type Db struct {
	Key   string `json:"key,omitempty" mapstructure:"key"`
	Value string `json:"value,omitempty" mapstructure:"value"`
}

func (db *Db)ToSortString() []byte {
	return []byte(db.Key + db.Value)
}

type Validator struct {
	PubKey string `json:"pubkey,omitempty" mapstructure:"pubkey"`
	Power  int64  `json:"power,omitempty" mapstructure:"power"`
}

func (db *Validator)ToSortString() []byte {
	return []byte(strconv.Itoa(db.Power) + db.PubKey)
}

type Account struct {
	PubKey string `json:"pubkey,omitempty" mapstructure:"pubkey"`
	Power  int64  `json:"power,omitempty" mapstructure:"power"`
}

func (db *Account)ToSortString() []byte {
	return []byte(strconv.Itoa(db.Power) + db.PubKey)
}

