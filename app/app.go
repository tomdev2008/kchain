package app

import (
	"github.com/tendermint/tmlibs/log"
	"github.com/gin-gonic/gin"
	kcfg "kchain/types/cfg"
	abci_client "kchain/types/abci"
	store_client "kchain/types/store"
	event_client "kchain/types/events"
	tx_handler "kchain/types/tx"
	"net/http"
	"os"
)

func Run() {
	var cfg = kcfg.GetConfig()
	var logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout)).With("module", "app")

	logger.Info("init abci client")
	abci_client.Init(cfg.Config.RPC.ListenAddress)

	logger.Info("init store_client")
	store := store_client.InitStoreClient()

	logger.Info("init event")
	event_client.Init()

	app := gin.Default()
	service := app.Group("/service")
	{
		service.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
			return
		})
	}

	v1 := app.Group("/v1")
	{
		//v1.Use(AuthRequired())
		v1.POST("/tx", func(c *gin.Context) {
			t := &tx_handler.Tx{}
			if err := c.ShouldBindJSON(t); err != nil {
				logger.Error(err.Error())
				c.JSON(http.StatusOK, gin.H{
					"code": "error",
					"msg":err.Error(),
				})
				return
			}
			t.Do()

			c.JSON(http.StatusOK, gin.H{
				"code": "ok",
				"data":store.Get(t.ID),
			})

			return
		})
		v1.GET("/id/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				c.JSON(http.StatusOK, gin.H{
					"code": "ok",
					"data":"id参数不存在",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"code": "ok",
				"data":store.Get(id),
			})
		})
	}
	logger.Info(cfg.App.Addr)
	app.Run(cfg.App.Addr)
}