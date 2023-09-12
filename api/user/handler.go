package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//用户路由处理函数

type HandlerUser struct {
}

func NewHandlerUser() *HandlerUser {
	return &HandlerUser{}
}

func (hu *HandlerUser) login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
}
