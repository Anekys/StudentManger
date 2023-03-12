package admin

import (
	"StudentManger/service"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func teacherMange(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("aid")
	fmt.Println(c.Get("adminStruct"))
	uid := value.(string)
	admin := service.FindAdminByAid(uid)
	teachers := service.FindAllTeachers(1)
	fmt.Println(teachers)
	c.HTML(http.StatusOK, "adminTeacherMange.html", gin.H{
		"name":        admin.Name,
		"teacherList": teachers,
	})
}
