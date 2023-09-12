package user

import (
	router "DRSP/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	log.Println("init user router")
	router.Register(&RouterUser{})
}

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	h := NewHandlerUser()
	fmt.Println("路由注册成功！")
	r.GET("/index/login", h.login)

}
