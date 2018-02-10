package transaction

type Db struct {
	Key   string `json:"key,omitempty" mapstructure:"key"`
	Value string `json:"value,omitempty" mapstructure:"value"`
}

type Validator struct {
	PubKey string `json:"pubkey,omitempty" mapstructure:"pubkey"`
	Power  int64  `json:"power,omitempty" mapstructure:"power"`
}

type Account struct {
	PubKey string `json:"pubkey,omitempty" mapstructure:"pubkey"`
	Power  int64  `json:"power,omitempty" mapstructure:"power"`
}
