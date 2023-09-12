package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*登陆路由注册*/

type HandlerUser struct {
}

func NewHandlerUser() *HandlerUser {
	return &HandlerUser{}
}

func (hu *HandlerUser) login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
}
