package main

import (
	"github.com/gin-gonic/gin"
	"sysStatus/util"
)

func main() {
	engine := gin.Default()
	util.RegisterRouter(engine)
	engine.Run(":9000")
}
