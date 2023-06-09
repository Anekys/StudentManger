package main

import (
	"StudentManger/controller"
	_ "StudentManger/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/images", "web/images")
	r.Static("/fonts", "web/fonts")
	r.Static("/vendors", "web/vendors")
	r.Static("/css", "web/css")
	r.Static("/js", "web/js")
	fmt.Println("学生管理系统启动")
	r.LoadHTMLGlob("web/pages/**/*.html")
	controller.LoadRouter(r)
	err := r.Run(":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
