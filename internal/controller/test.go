package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lamber92/go-cover-example/internal/logic"
)

func TestAPI1(ctx *gin.Context) {
	res, err := logic.Test1(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 1000,
			"msg":  err.Error(),
			"data": map[string]interface{}{},
		})
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"data": map[string]interface{}{"text": res},
	})
}

func TestAPI2(ctx *gin.Context) {
	flag := ctx.DefaultQuery("flag", "")
	res, err := logic.Test2(ctx, flag)
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 1000,
			"msg":  err.Error(),
			"data": map[string]interface{}{},
		})
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"data": map[string]interface{}{"text": res},
	})
}
