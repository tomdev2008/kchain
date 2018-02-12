package tx

import (
	"kchain/types"

	"kchain/types/events"
	"strings"
)

type Tx types.Tx

func (tx *Tx) FromBytes(bs []byte) error {
	return json.Unmarshal(bs, tx)
}

func (tx *Tx) ToBytes() []byte {
	d, _ := json.Marshal(tx)
	return d
}

func (tx *Tx) Do() {
	// 把tx缓存到事件库

	// 获得event

	// 执行event handle

	logger.Debug(tx.Event)
	if e := events.GetEvent(tx.Event); e != nil {
		if strings.Compare(tx.Sync, "true") == 0 {
			e.Handler(tx.ToBytes())
		} else {
			go e.Handler(tx.ToBytes())
		}

	} else {
		logger.Error("事件不存在")
	}
}
