package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SysStatusHandle(ctx *gin.Context) {
	var sysInfo SysInfo
	UpdateSystemStatus(&sysInfo.SysStatus)
	ctx.HTML(http.StatusOK, "user/sysstatus.html", sysInfo)
}

func RegisterRouter(engine *gin.Engine) {
	engine.LoadHTMLGlob("templates/**/*")
	engine.Static("/css", "templates/statics")
	engine.GET("/sysstatus", SysStatusHandle)
}

func GetSysStatusJson() string {
	var sysInfo SysInfo
	UpdateSystemStatus(&sysInfo.SysStatus)
	res, err := json.MarshalIndent(sysInfo, "", "\t")
	if err != nil {
		log.Println(err)
	}
	return string(res)
}
func Loginit() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
