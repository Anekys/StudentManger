package admin

import (
	"StudentManger/module"
	"StudentManger/service"
	"StudentManger/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Teacher struct {
	Tid      string `form:"tid"`
	Name     string `form:"tname" msg:"请填写姓名!" binding:"required"`
	Email    string `form:"email" msg:"请填写邮箱地址!" binding:"required"`
	PassWord string `form:"password" msg:"请填写密码!" binding:"required"`
}

func EditTeacher(c *gin.Context) {
	tid := c.Query("tid")
	session := sessions.Default(c)
	value := session.Get("aid")
	uid := value.(string)
	admin := service.FindAdminByAid(uid)
	if tid == "" {
		// 新增教师
		c.HTML(200, "adminEditTeacher.html", gin.H{
			"name": admin.Name,
		})
	} else {
		// 编辑教师
		tea := service.FindTeacherByTid(tid)
		c.HTML(200, "adminEditTeacher.html", gin.H{
			"tname":    tea.Name,
			"email":    tea.Email,
			"password": tea.PassWord,
			"tid":      tea.TID,
		})
	}
}

func PushTeacher(c *gin.Context) {
	var teacher Teacher
	err := c.ShouldBind(&teacher)
	if teacher.Tid == "" {
		if err != nil {
			errStr := utils.GetVaildMsg(err, &teacher)
			c.Redirect(301, "/admin/editTeacher?msg="+errStr)
		} else {
			tid := utils.Md5Encrypt(teacher.Email)
			tea := module.Teacher{
				TID:      tid,
				Name:     teacher.Name,
				Email:    teacher.Email,
				PassWord: teacher.PassWord,
			}
			if service.AddTeacher(tea) {
				c.Redirect(301, "/admin/teacherMange")
			} else {
				c.Redirect(301, "/admin/editTeacher?msg=添加教师失败!请检查后重试!")
			}
		}
	} else {
		if err != nil {
			errStr := utils.GetVaildMsg(err, &teacher)
			c.Redirect(301, "/admin/editTeacher?tid="+teacher.Tid+"&msg="+errStr)
		} else {
			tea := module.Teacher{
				TID:      teacher.Tid,
				Name:     teacher.Name,
				Email:    teacher.Email,
				PassWord: teacher.PassWord,
			}
			if service.UpdateTeacherByTid(teacher.Tid, tea) {
				c.Redirect(301, "/admin/teacherMange")
			} else {
				c.Redirect(301, "/admin/editTeacher?tid="+teacher.Tid+"&msg=msg=添加教师失败!请检查后重试!")
			}
		}
	}
}
