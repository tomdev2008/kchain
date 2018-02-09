package types

type Tx struct {
	ID     string        `json:"id" binding:"required"`
	Sync   string        `json:"sync" binding:"required"`
	Event  string        `json:"event" binding:"required"`
	Params interface{}   `json:"params" binding:"required"`
	isDone bool
}