package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 标准响应格式
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *Response) Json(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, r)
	ctx.Abort()
}

func (r *Response) UnAuthorized(ctx *gin.Context)  {
	r.Code = http.StatusUnauthorized
	r.Msg = "未授权！"

	r.Json(ctx)
}

func (r *Response) BadRequest(ctx *gin.Context)  {
	r.Code = http.StatusBadRequest
	r.Msg = "错误的请求！"

	r.Json(ctx)
}

func (r *Response) ServerError(ctx *gin.Context, data interface{}) {
	r.Code = http.StatusInternalServerError
	r.Msg = "服务器内部错误！"
	r.Data = data

	r.Json(ctx)
}

func (r *Response) Success(ctx *gin.Context, data interface{})  {
	r.Code = http.StatusOK
	r.Msg = "操作成功！"
	r.Data = data

	r.Json(ctx)
}

func (r *Response) Forbidden(ctx *gin.Context)  {
	r.Code = http.StatusForbidden
	r.Msg = "账户或密码错误！"

	r.Json(ctx)
}