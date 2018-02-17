package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin/binding"
	//"kchain/utils/validation"

)

func InitUrls(router *gin.Engine) {

	//binding.Validator.RegisterValidation("bookabledate", validation.BookableDate)]

	service := router.Group("/service")
	service.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")

		return
	})

	v1 := router.Group("/v1")
	//v1.Use(AuthRequired())
	v1.POST("/tx", _tx_handler)
	v1.GET("/tx", _tx_handler)
	v1.GET("/id/:id", _id_handler)

}
