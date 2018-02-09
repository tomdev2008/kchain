package app

import (
	"github.com/gin-gonic/gin"

	abci_client "kchain/types/abci"
	store_client "kchain/types/store"
	event_client "kchain/types/events"
	klog "kchain/utils/log"
)

func Run() {

	klog.Init()

	logger.Info("init abci client", "init", "abci_client")
	abci_client.Init(cfg().Config.RPC.ListenAddress)

	logger.Info("init store_client", "init", "store_client")
	store_client.InitStoreClient()

	logger.Info("init event", "init", "event_client")
	event_client.Init()

	app := gin.Default()

	logger.Info("init urls", "init", "urls")
	InitUrls(app)
	app.Run(cfg().App.Addr)
}