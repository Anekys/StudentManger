package admin

import (
	"StudentManger/module"
	"StudentManger/service"
	"StudentManger/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type StuForm struct {
	UID      string `form:"uid"`
	Name     string `form:"sname" msg:"请输入姓名" binding:"required"`    //姓名
	Age      int    `form:"age" msg:"请输入年龄" binding:"required"`      //年龄
	Gender   string `form:"gender" msg:"请输入性别" binding:"required"`   //性别 0.女 1.男
	Class    string `form:"class" msg:"请输入班级" binding:"required"`    //班级
	Phone    string `form:"phone" msg:"请输入联系方式" binding:"required"`  //联系方式
	Email    string `form:"email" msg:"请输入电子邮箱" binding:"required"`  //用户名
	PassWord string `form:"password" msg:"请输入密码" binding:"required"` //密码
}

func editStudent(c *gin.Context) {
	uid := c.Query("uid")
	session := sessions.Default(c)
	value := session.Get("aid")
	aid := value.(string)
	admin := service.FindAdminByAid(aid) // todo - 此处以及管理员和教师edit界面需要展示当前登录用户的姓名
	if uid == "" {
		// 新增学生
		c.HTML(200, "adminEditStudent.html", gin.H{
			"name": admin.Name,
		})
	} else {
		// 编辑学生
		stu := service.FindStudentByUid(uid)
		c.HTML(200, "adminEditStudent.html", gin.H{
			"name":     admin.Name,
			"sname":    stu.Name,
			"class":    stu.Class,
			"age":      stu.Age,
			"email":    stu.Email,
			"password": stu.PassWord,
			"phone":    stu.Phone,
			"gender":   stu.Gender,
			"uid":      stu.UID,
		})
	}
}

func PushStudent(c *gin.Context) {
	var student StuForm
	err := c.ShouldBind(&student)
	if student.UID == "" {
		if err != nil {
			errStr := utils.GetVaildMsg(err, &student)
			c.Redirect(301, "/admin/editStudent?msg="+errStr)
		} else {
			uid := utils.Md5Encrypt(student.Email)
			stu := module.Student{
				UID:      uid,
				Name:     student.Name,
				Age:      student.Age,
				Gender:   student.Gender,
				Class:    student.Class,
				Phone:    student.Phone,
				Email:    student.Email,
				PassWord: student.PassWord,
			}
			if service.AddStudent(stu) {
				c.Redirect(301, "/admin/main?msg=添加学生成功!")
			} else {
				c.Redirect(301, "/admin/main?msg=添加学生失败!请检查后重试!")
			}
		}
	} else {
		if err != nil {
			errStr := utils.GetVaildMsg(err, &student)
			c.Redirect(301, "/admin/editStudent?uid="+student.UID+"&msg="+errStr)
		} else {
			stu := module.Student{
				UID:      student.UID,
				Name:     student.Name,
				Age:      student.Age,
				Gender:   student.Gender,
				Class:    student.Class,
				Phone:    student.Phone,
				Email:    student.Email,
				PassWord: student.PassWord,
			}
			if service.UpdateStudentByUid(student.UID, stu) {
				c.Redirect(301, "/admin/main?msg=更新学生信息成功!")
			} else {
				c.Redirect(301, "/admin/editStudent?uid="+student.UID+"&msg=添加学生失败!请检查后重试!")
			}
		}
	}
}

func deleteStudent(c *gin.Context) {
	uid := c.Query("uid")
	if uid == "" {
		c.Redirect(301, "/admin/main?msg="+"错误的学生ID!")
	} else {
		fmt.Println("删除学生：", uid)
		if service.DelStudentByUID(uid) {
			c.Redirect(301, "/admin/main?msg="+"删除成功!")
		} else {
			c.Redirect(301, "/admin/main?msg="+"删除失败!请稍后重试??")
		}
	}
}
