package app

import (
	"github.com/gin-gonic/gin"
)

func Run() {

	logger = cfg().GetLogWithKeyVals("module", "app")

	app := gin.Default()

	logger.Info("init urls", "init", "urls")
	InitUrls(app)
	app.Run(cfg().App.Addr)
}