package tx

import (
	"kchain/types/cnst"
)


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
