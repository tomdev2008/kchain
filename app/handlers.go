package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kchain/types"
	cnst "kchain/types/cnst"
)

func _tx_handler(c *gin.Context) {
	t := &types.Tx{}
	if err := c.ShouldBindJSON(t); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": "error",
			"msg":err.Error(),
		})
		return
	}

	switch t.Event {
	case cnst.DbSet:


	}

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