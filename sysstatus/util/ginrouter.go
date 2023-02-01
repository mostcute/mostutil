package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mostcute/mostutil/sysstatus/templates"
	"html/template"
	"log"
	"net/http"
)

func SysStatusHandle(ctx *gin.Context) {
	var sysInfo SysInfo
	UpdateSystemStatus(&sysInfo.SysStatus)
	ctx.HTML(http.StatusOK, "user/sysstatus.html", sysInfo)
}

func RegisterRouter(engine *gin.Engine) {
	templ := template.Must(template.New("").ParseFS(templates.TemplatesEmbed, "./**/*"))
	engine.SetHTMLTemplate(templ)
	engine.GET("/sysstatus", SysStatusHandle)
	engine.Use(templates.StaticServer())
}

func GetSysStatus() (sysInfo SysInfo) {
	UpdateSystemStatus(&sysInfo.SysStatus)
	return
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
