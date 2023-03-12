package teacher

import (
	"StudentManger/service"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Main(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("tid")
	fmt.Println(c.Get("teacherStruct"))
	uid := value.(string)
	teacher := service.FindTeacherByTid(uid)
	//classmate := service.FindStudentsByClass(student.Class)
	c.HTML(http.StatusOK, "teacherMain.html", gin.H{
		"name": teacher.Name,
		//"class":       student.Class,
		//"studentList": classmate,
	})
}
