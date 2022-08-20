package engine

import (
	"net/http"

	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/gin-gonic/gin"
)

func Result(h Handle) Handle {
	return func(c *Context) (result interface{}, err error) {
		if r, err := h(c); err != nil {
			JsonError(c, r, err)
		} else {
			JsonResult(c, r)
		}
		return nil, nil
	}
}

func JsonError(ctx *Context, data interface{}, err error) {
	code := 500
	if v, ok := err.(e.ErrCode); ok {
		code = int(v)
	}
	ctx.Context.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  err.Error(),
	})
	ctx.Context.Abort()
}

func JsonResult(ctx *Context, data interface{}) {
	if data == nil {
		data = "success"
	}
	ctx.Context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}