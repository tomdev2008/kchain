package app

import (
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin/binding"
	//"kchain/utils/validation"
)

func InitUrls(router *gin.Engine) {

	//binding.Validator.RegisterValidation("bookabledate", validation.BookableDate)]

	router.POST("/tx", func(c *gin.Context) {
		//c.ShouldBindJSON()
	})
	router.GET("/tx", func(c *gin.Context) {})
	router.GET("/id", func(c *gin.Context) {})

}
