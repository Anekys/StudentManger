package teacher

import (
	"StudentManger/module"
	"StudentManger/service"
	"StudentManger/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
)

func ConfirmMange(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("tid")
	tid := value.(string)
	teacher := service.FindTeacherByTid(tid)
	courses := service.FindTeacherCourse(tid, 1)
	c.HTML(200, "teacherConfirm.html", gin.H{
		"title":      "考勤管理",
		"name":       teacher.Name,
		"courseList": courses,
	})
}

func StartConfirm(c *gin.Context) {
	kid := c.Query("kid")
	if kid == "" {
		c.Redirect(302, "/teacher/confirmMange?msg=不存在的课程,或课程ID错误!")
		return
	}
	if status, msg := service.StartConfirm(kid); status {
		startConfirm := module.CourseInfo{Status: 1}
		service.UpdateCourseById(kid, startConfirm)
		c.Redirect(302, "/teacher/confirmMange?msg=开始考勤成功")
	} else {
		c.Redirect(302, "/teacher/confirmMange?msg="+msg)
	}
}

func EndConfirm(c *gin.Context) {
	kid := c.Query("kid")
	if kid == "" {
		c.Redirect(302, "/teacher/confirmMange?msg=不存在的课程,或课程ID错误!")
		return
	}
	course := service.FindCourseById(kid)
	if !utils.Exists(kid) {
		// 添加考勤结束的记录
		result := module.ConfirmResult{
			KID:   kid,
			CName: course.CName,
			Count: 0,
			Time:  time.Now().Format("2006年01月02日15时04分05秒"),
		}
		service.AddConfirmResult(result)
		// 修改考勤状态为结束考勤
		flag := service.UpdateCourseByIdWithField(kid, "status", 0)
		fmt.Println(flag)
		c.Redirect(302, "/teacher/confirmMange?msg=结束考勤成功!")
		return
	}
	flag := service.UpdateCourseByIdWithField(kid, "status", 0)
	fmt.Println(flag)
	uidList := utils.HMGet(kid)
	count := utils.HMLen(kid)
	nowTime := time.Now().Format("2006年01月02日15时04分05秒")
	var ConfirmResults []module.ConfirmResult
	for k, v := range uidList {
		result := module.ConfirmResult{
			KID:     kid,
			CName:   course.CName,
			Count:   count,
			UID:     k,
			Student: v,
			Time:    nowTime,
		}
		ConfirmResults = append(ConfirmResults, result)
	}
	if service.AddConfirmResult(ConfirmResults) {
		utils.Del(kid)
		c.Redirect(302, "/teacher/confirmMange?msg=结束考勤成功!")
	} else {
		c.Redirect(302, "/teacher/confirmMange?msg=结束考勤失败!")
	}

}

func ConfirmResultMange(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("tid")
	tid := value.(string)
	teacher := service.FindTeacherByTid(tid)
	confirmResult := service.FindTeacherConfirmResult(tid)
	c.HTML(200, "teacherConfirmResult.html", gin.H{
		"title":       "考勤记录",
		"name":        teacher.Name,
		"confirmList": confirmResult,
	})

}

func ConfirmDetail(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("tid")
	tid := value.(string)
	teacher := service.FindTeacherByTid(tid)
	kid := c.Query("kid")
	timeStr := c.Query("time")
	if kid == "" || timeStr == "" {
		c.Redirect(302, "/teacher/confirmResult?msg=提交的参数有误!")
	}
	confirmResult := service.FindConfirmResultByKidAndTime(kid, timeStr)
	c.HTML(200, "teacherConfirmDetail.html", gin.H{
		"title":       "缺勤详情",
		"name":        teacher.Name,
		"confirmList": confirmResult,
	})
}
