package main

import (
	router "DRSP/api"
	_ "DRSP/api/api"
	_ "DRSP/internal/database/gorms"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run()
}
