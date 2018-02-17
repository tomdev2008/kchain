package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	kcfg "kchain/types/cfg"
)

func _tx_handler(c *gin.Context) {
	t := &Tx{}
	if err := c.ShouldBindJSON(t); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": "error",
			"msg":err.Error(),
		})
		return
	}



	c.JSON(http.StatusOK, gin.H{
		"code": "ok",
		"data":kcfg.DbGet([]byte(t.ID)),
	})
}

func _id_handler(c *gin.Context) {
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
		"data":kcfg.DbGet([]byte(id)),
	})
}