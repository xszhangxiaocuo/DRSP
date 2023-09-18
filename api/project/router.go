package project

import (
	router "DRSP/api"
	"github.com/gin-gonic/gin"
	"log"
)

// 使用init()通过import调包注册路由
func init() {
	log.Println("init project router")
	router.Register(&RouterProject{})
}

type RouterProject struct {
}

func (rp *RouterProject) Route(r *gin.Engine) {
	h := NewHandlerProject()

	r.POST("/project/upload", h.upload)
	r.POST("/project/chat", h.chat)
}
