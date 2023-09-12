package user

import (
	router "DRSP/api"
	"github.com/gin-gonic/gin"
	"log"
)

// 使用init()通过import调包注册路由
func init() {
	log.Println("init user router")
	router.Register(&RouterUser{})
}

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	h := NewHandlerUser()

	r.GET("/index/login", h.login)

}
