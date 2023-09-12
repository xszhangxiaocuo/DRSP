package router

import (
	"github.com/gin-gonic/gin"
)

// 路由注册方法接口
type Router interface {
	Route(r *gin.Engine)
}

// RegisterRouter
type RegisterRouter struct {
}

var routers []Router

// 批量注册路由
func InitRouter(r *gin.Engine) {

	for _, ro := range routers {
		ro.Route(r)
	}
}

func Register(ro ...Router) {
	routers = append(routers, ro...)
}
