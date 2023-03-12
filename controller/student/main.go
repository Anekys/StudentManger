package student

import (
	"StudentManger/service"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Main(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("uid")
	fmt.Println(c.Get("studentStruct"))
	uid := value.(string)
	student := service.FindStudentByUid(uid)
	classmate := service.FindStudentsByClass(student.Class)
	c.HTML(http.StatusOK, "studentMain.html", gin.H{
		"name":        student.Name,
		"class":       student.Class,
		"studentList": classmate,
	})
}
