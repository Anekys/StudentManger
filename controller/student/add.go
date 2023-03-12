package student

import (
	"StudentManger/module"
	"StudentManger/service"
	"StudentManger/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Form struct {
	Name     string `form:"name" msg:"请输入姓名" binding:"required"`     //姓名
	Age      int    `form:"age" msg:"请输入年龄" binding:"required"`      //年龄
	Gender   string `form:"gender" msg:"请输入性别" binding:"required"`   //性别 0.女 1.男
	Class    string `form:"class" msg:"请输入班级" binding:"required"`    //班级
	Phone    string `form:"phone" msg:"请输入联系方式" binding:"required"`  //联系方式
	Email    string `form:"email" msg:"请输入电子邮箱" binding:"required"`  //用户名
	PassWord string `form:"password" msg:"请输入密码" binding:"required"` //密码
}

func AddStudent(c *gin.Context) {
	var student Form
	if err := c.ShouldBind(&student); err != nil {
		errStr := utils.GetVaildMsg(err, &student)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  errStr,
		})
		return
	}
	//gender, _ := strconv.Atoi(student.Gender)
	Student := module.Student{
		Name:     student.Name,
		Email:    student.Email,
		PassWord: student.PassWord,
		Age:      student.Age,
		Gender:   student.Gender,
		Class:    student.Class,
		Phone:    student.Phone,
	}
	if service.AddStudent(Student) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "添加学生成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "添加学生失败",
		})
	}

}

func DelStudent(c *gin.Context) {
	var student Form
	if err := c.ShouldBind(&student); err != nil {
		errStr := utils.GetVaildMsg(err, &student)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  errStr,
		})
		return
	}
	//gender, _ := strconv.Atoi(student.Gender)
	Student := module.Student{
		Name:     student.Name,
		Email:    student.Email,
		PassWord: student.PassWord,
		Age:      student.Age,
		Gender:   student.Gender,
		Class:    student.Class,
		Phone:    student.Phone,
	}
	fmt.Println(Student)
}
