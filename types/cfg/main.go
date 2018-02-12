package cfg

import (
	cfg "github.com/tendermint/tendermint/config"
	dbm "github.com/tendermint/tmlibs/db"
	tlog "github.com/tendermint/tmlibs/log"
	tmflags "github.com/tendermint/tmlibs/cli/flags"
	ttypes "github.com/tendermint/tendermint/types"
	crypto "github.com/tendermint/go-crypto"
	c "github.com/tendermint/tendermint/rpc/client"
	"os"
)

func (ss *services) DbSet(k, v []byte) {
	ss.store.Set(k, v)
}

func (ss *services) DbGet(k []byte) []byte {
	return ss.store.Get(k)
}

func (ss *services) GetLogWithKeyVals(keyvals ...interface{}) func() tlog.Logger {
	return ss.log.With(keyvals...)
}

func (ss *services) GetPrivKey() crypto.PrivKey {
	return ss.pk.PrivKey
}

func (ss *services) GetPubKey() crypto.PubKey {
	return ss.pk.GetPubKey()
}

func (ss *services) GetAddress() []byte {
	return ss.pk.GetAddress().Bytes()
}

func (ss *services) Sign(msg []byte) ([]byte, error) {
	if _sign, err := ss.pk.Sign(msg); err != nil {
		return nil, err
	} else {
		return _sign.Bytes(), nil
	}
}

func (ss *services) Verify(msg, sign []byte) ([]byte, error) {
	if _sign, err := crypto.SignatureFromBytes(sign); err != nil {
		return nil, err
	} else {
		return ss.pk.PubKey.VerifyBytes(msg, _sign), nil
	}
}

func (ss *services) GetAbciClient() (*c.HTTP) {
	return ss.client
}

func initServices() *services {
	s := &services{
		Config:cfg.DefaultConfig(),
		App:&appConfig{
			Name:"kchain",
			Addr:":9000",
		},
		Node:nil,
	}

	// 初始化db
	if store, err := dbm.NewGoLevelDB(
		"app",
		s.Config.DBDir(),
	); err != nil {
		panic(err.Error())
	} else {
		s.store = store
	}

	// 初始化log
	if klog, err := tmflags.ParseLogLevel(
		s.Config.LogLevel,
		tlog.NewTMLogger(tlog.NewSyncWriter(os.Stderr)),
		"error",
	); err != nil {
		panic(err.Error())
	} else {
		s.log = klog
	}

	// 私钥获得
	s.pk = ttypes.LoadOrGenPrivValidatorFS(s.Config.PrivValidatorFile())

	s.client = c.NewHTTP(s.Config.ProxyApp, "/ws")

	return s
}

func GetConfig() func() *services {
	return func() *services {
		once.Do(func() {
			instance = initServices()
		})
		return instance
	}
}


