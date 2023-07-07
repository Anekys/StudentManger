package admin

import (
	"StudentManger/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Main(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("aid")
	//fmt.Println(c.Get("adminStruct"))
	uid := value.(string)
	admin := service.FindAdminByAid(uid)
	studentList := service.FindAllStudents(1)
	c.HTML(http.StatusOK, "adminMain.html", gin.H{
		"name":        admin.Name,
		"studentList": studentList,
	})
}
