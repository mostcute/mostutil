package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mostcute/mostutil/sysstatus/util"
)

func main() {
	engine := gin.Default()
	util.RegisterRouter(engine)
	engine.Run(":9000")
}
