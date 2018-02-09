package abci

import (
	c "github.com/tendermint/tendermint/rpc/client"
	"sync"
)

type abciClient struct {
	Url        string
	WsEndpoint string
	client     *c.HTTP
}

var (
	once sync.Once
	instance *abciClient
)


//tcp://127.0.0.1:46658
func Init(url string) {
	once.Do(func() {
		instance = &abciClient{
			WsEndpoint:"/ws",
			Url:url,
			client:c.NewHTTP(url, "/ws"),
		}
	})
}

func GetAbciClient() *c.HTTP {
	return instance.client
}

/*
func (a *Abci) Status() (*ctypes.ResultStatus, error) {
	return a.client.Status()
}

func (a *Abci) ABCIInfo() (*ctypes.ResultABCIInfo, error) {
	return a.client.ABCIInfo()
}

func (a *Abci) ABCIQuery(path string, data d.Bytes) (*ctypes.ResultABCIQuery, error) {
	return a.client.ABCIQuery(path, data)
}

func (a *Abci) ABCIQueryWithOptions(path string, data d.Bytes, opts c.ABCIQueryOptions) (*ctypes.ResultABCIQuery, error) {
	return a.client.ABCIQueryWithOptions(path, data, opts)
}

func (a *Abci) BroadcastTxCommit(tx types.Tx) (*ctypes.ResultBroadcastTxCommit, error) {
	return a.client.BroadcastTxCommit(tx)
}

func (a *Abci) BroadcastTxAsync(tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	return a.client.BroadcastTxAsync(tx)
}

func (a *Abci) BroadcastTxSync(tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	return a.client.BroadcastTxSync(tx)
}

func (a *Abci) NetInfo() (*ctypes.ResultNetInfo, error) {
	return a.client.NetInfo()
}

func (a *Abci) DumpConsensusState() (*ctypes.ResultDumpConsensusState, error) {
	return a.client.DumpConsensusState()
}

func (a *Abci) Genesis() (*ctypes.ResultGenesis, error) {
	return a.client.Genesis()
}

func (a *Abci) Block(height *int64) (*ctypes.ResultBlock, error) {
	return a.client.Block(height)
}

func (a *Abci) BlockResults(height *int64) (*ctypes.ResultBlockResults, error) {
	return a.client.BlockResults(height)
}

func (a *Abci) Commit(height *int64) (*ctypes.ResultCommit, error) {
	return a.client.Commit(height)
}

func (a *Abci)Tx(hash []byte, prove bool) (*ctypes.ResultTx, error) {
	return a.client.Tx(hash, prove)
}

func (a *Abci) TxSearch(query string, prove bool) ([]*ctypes.ResultTx, error) {
	return a.client.TxSearch(query, prove)
}
func (a *Abci) Validators(height *int64) (*ctypes.ResultValidators, error) {
	return a.client.Validators(height)
}
*/