package transaction

type Db struct {
	Key   string `json:"key,omitempty" mapstructure:"key"`
	Value string `json:"value,omitempty" mapstructure:"value"`
}
