package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	kcfg "kchain/types/cfg"
	kts "kchain/types"
	tdata "github.com/tendermint/go-wire/data"
	"github.com/tendermint/tendermint/types"
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

	if d, err := json.MarshalToString(&kts.Transaction{
		SignPubKey:t.SignPubKey,
		Signature:t.Signature,
		Data:t.Data,
	}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "error",
			"msg":err.Error(),
		})
		return
	} else {
		if res, err := kcfg.Abci().BroadcastTxCommit(types.Tx(d)); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": "error",
				"msg":err.Error(),
			})
			return

		} else {
			d, _ := json.MarshalToString(res)
			c.JSON(http.StatusOK, gin.H{
				"code": "ok",
				"data":d,
			})
			return
		}
	}
}

func _q_tx_handler(c *gin.Context) {

	t := &Tx{}
	if err := c.ShouldBindJSON(t); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": "error",
			"msg":err.Error(),
		})
		return
	}

	if d, err := json.MarshalToString(&kts.Transaction{
		SignPubKey:t.SignPubKey,
		Signature:t.Signature,
		Data:t.Data,
	}); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "error",
			"msg":err.Error(),
		})
		return

	} else {
		if res, err := kcfg.Abci().ABCIQuery("", tdata.Bytes(d)); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": "error",
				"msg":err.Error(),
			})
			return
		} else {

			d, _ := json.MarshalToString(res.Response)
			c.JSON(http.StatusOK, gin.H{
				"code": "error",
				"data":d,
			})
		}
	}

}

func _id_handler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": "ok",
			"data":"id is null",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "ok",
		"data":kcfg.DbGet(id),
	})
}
