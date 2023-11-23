package response

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Fail(c *gin.Context, err error) {
	errResp := NewError(CodeSystemError)
	if v, ok := err.(Error); ok {
		errResp = v
	} else {
		_ = errResp.WithErr(err)
	}
	if errResp.GetErr() != nil {
		zap.L().Error("server error ...", zap.Error(errResp.GetErr())) // 记录错误

	}
	c.JSON(http.StatusOK, &RetData{
		Code: errResp.GetCode(),
		Msg:  errResp.GetMsg(),
		Data: nil,
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &RetData{
		Code: CodeOk,
		Msg:  Text(CodeOk),
		Data: data,
	})
}

type RetData struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
