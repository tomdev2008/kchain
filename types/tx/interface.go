package tx

type Db interface {
	Set(k, v[]byte)
}
