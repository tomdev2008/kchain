package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	store_client "kchain/types/store"
	tx_handler "kchain/types/tx"
)

func _tx_handler(c *gin.Context) {
	store := store_client.GetStoreClient()

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
		"data":store().Get(t.ID),
	})
}

func _id_handler(c *gin.Context) {
	store := store_client.GetStoreClient()
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
		"data":store().Get(id),
	})
}