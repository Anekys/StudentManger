package student

import (
	"StudentManger/service"
	"StudentManger/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ConfirmMange(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("uid")
	uid := value.(string)
	student := service.FindStudentByUid(uid)
	courses := service.FindChooseCourse(1, uid)
	c.HTML(200, "studentConfirmMange.html", gin.H{
		"name":       student.Name,
		"courseList": courses,
		"title":      "考勤管理",
	})
}

func StartConfirm(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("uid")
	uid := value.(string)
	kid := c.Query("kid")
	if !utils.Exists(kid) {
		c.Redirect(302, "/student/confirmMange?msg=该课程还未开始考勤或所有人已打卡完毕!")
		return
	}
	student := service.FindStudentByUid(uid)
	// 已经打过卡的status为1 , 还未打卡的 status为0
	status := 0
	if utils.HMIsMember(kid, uid) {
		status = 1
	}
	c.HTML(200, "studentConfirm.html", gin.H{
		"name":   student.Name,
		"kid":    kid,
		"status": status,
	})
}

func Confirm(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("uid")
	uid := value.(string)
	kid := c.PostForm("kid")
	if count := utils.HMDel(kid, uid); count > 0 {
		c.Redirect(302, "/student/confirmMange?msg=打卡成功!")
	} else {
		c.Redirect(302, "/student/startConfirm?kid="+kid+"&msg=打卡失败")
	}
}
