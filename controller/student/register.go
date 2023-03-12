package student

import (
	"StudentManger/module"
	"StudentManger/service"
	"StudentManger/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var empty module.Student

func RegisterStudent(c *gin.Context) {
	var student Form
	if err := c.ShouldBind(&student); err != nil {
		errStr := utils.GetVaildMsg(err, &student)
		c.Redirect(http.StatusMovedPermanently, "/student/register?msg="+errStr)
		return
	}
	//gender, _ := strconv.Atoi(student.Gender)
	uid := utils.Md5Encrypt(student.Email)
	if service.FindStudentByUid(uid) != empty {
		c.Redirect(http.StatusMovedPermanently, "/student/register?msg="+"该邮箱已注册!")
		return
	}

	Student := module.Student{
		UID:      uid,
		Name:     student.Name,
		Email:    student.Email,
		PassWord: student.PassWord,
		Age:      student.Age,
		Gender:   student.Gender,
		Class:    student.Class,
		Phone:    student.Phone,
	}

	if service.AddStudent(Student) {
		c.Redirect(http.StatusMovedPermanently, "/login?msg=注册成功!")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/student/register?msg=注册失败,请重试")
	}

}
