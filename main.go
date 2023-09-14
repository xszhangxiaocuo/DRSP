package main

import (
	"DRSP/api"
	_ "DRSP/api/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run()
}
