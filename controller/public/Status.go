package public

import (
	"StudentManger/module"
	"StudentManger/service"
	"StudentManger/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogInfo struct {
	Username  string `form:"email" msg:"用户名及密码不正确或不存在!" binding:"required"`
	Password  string `form:"password" msg:"用户名及密码不正确或不存在!" binding:"required"`
	Group     int    `form:"group" msg:"请选择用户组!" binding:"required"`
	KeepAlive string `form:"keep-alive"`
}

func Login(c *gin.Context) {
	var info LogInfo
	if err := c.ShouldBind(&info); err != nil {
		errStr := utils.GetVaildMsg(err, &info)
		c.HTML(http.StatusOK, "login.html", gin.H{
			"msg": errStr,
		})
		return
	}
	session := sessions.Default(c)
	TTL := 86400 // 默认session保存1天
	if info.KeepAlive != "" {
		TTL = 604800 // 勾选保持登录则session过期时间为7天
	}
	session.Options(sessions.Options{MaxAge: TTL}) // 设置登录有效期
	switch info.Group {
	case 1:
		stu := service.FindStudentByEmailPassword(info.Username, info.Password)
		var empty module.Student
		if stu == empty {
			c.Redirect(301, "/login?msg="+"用户名或密码不正确")
		} else {
			session.Set("uid", stu.UID)
			err := session.Save()
			if err != nil {
				c.Redirect(301, "/login?msg="+err.Error())
			}
			c.Redirect(301, "/student/main")
		}
	case 2:
		tea := service.FindTeacherByEmailPassword(info.Username, info.Password)
		var empty module.Teacher
		if tea == empty {
			c.Redirect(301, "/login?msg="+"用户名或密码不正确")
		} else {
			session.Set("tid", tea.TID)
			err := session.Save()
			if err != nil {
				c.Redirect(301, "/login?msg="+err.Error())
			}
			c.Redirect(301, "/teacher/main")
		}
		//c.Redirect(http.StatusMovedPermanently, "/teacher?uid="+result.UID)
	case 3:
		//c.Redirect(http.StatusMovedPermanently, "/student?uid="+result.UID)
		adm := service.FindAdminByEmailPassword(info.Username, info.Password)
		var empty module.Admin
		if adm == empty {
			c.Redirect(301, "/login?msg="+"用户名或密码不正确")
		} else {
			session.Set("aid", adm.AID)
			err := session.Save()
			if err != nil {
				c.Redirect(301, "/login?msg="+err.Error())
			}
			c.Redirect(301, "/admin/main")
		}
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("uid")
	session.Delete("tid")
	session.Delete("aid")
	err := session.Save()
	if err != nil {
		fmt.Println("logout error", err)
		return
	}
	c.Redirect(301, "/login?msg=账号注销成功")
}
func ForwardToLogin(c *gin.Context) {
	msg := c.DefaultQuery("msg", "")
	c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": msg,
	})
}

func ForwardToRegister(c *gin.Context) {
	msg := c.DefaultQuery("msg", "")
	c.HTML(http.StatusOK, "register.html", gin.H{
		"msg": msg,
	})
}
