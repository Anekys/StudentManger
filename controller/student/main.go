package student

import (
	"StudentManger/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Main(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("uid")
	//fmt.Println(c.Get("studentStruct"))
	uid := value.(string)
	student := service.FindStudentByUid(uid)
	classmate := service.FindStudentsByClass(student.Class)
	c.HTML(http.StatusOK, "studentMain.html", gin.H{
		"name":        student.Name,
		"class":       student.Class,
		"studentList": classmate,
	})
}

func viewInfo(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("uid")
	//fmt.Println(c.Get("studentStruct"))
	uid := value.(string)
	student := service.FindStudentByUid(uid)
	c.HTML(200, "studentViewInfo.html", gin.H{
		"name":     student.Name,
		"sname":    student.Name,
		"class":    student.Class,
		"age":      student.Age,
		"email":    student.Email,
		"password": student.PassWord,
		"phone":    student.Phone,
		"gender":   student.Gender,
		"uid":      student.UID,
	})

}