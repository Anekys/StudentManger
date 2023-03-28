package student

import (
	"StudentManger/module"
	"StudentManger/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func FindChooseCourse(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("uid")
	uid := value.(string)
	student := service.FindStudentByUid(uid)
	courses := service.FindChooseCourse(1, uid)
	c.HTML(200, "studentChooseCourse.html", gin.H{
		"name":       student.Name,
		"courseList": courses,
		"title":      "已选课程",
	})
}
func FindAllCourse(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("uid")
	uid := value.(string)
	student := service.FindStudentByUid(uid)
	courses := service.FindNotChooseCourse(1, uid)
	c.HTML(200, "studentAllCourse.html", gin.H{
		"name":       student.Name,
		"courseList": courses,
		"title":      "全部课程",
	})
}
func ChooseCourse(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("uid")
	uid := value.(string)
	student := service.FindStudentByUid(uid)
	kid := c.Query("kid")
	if kid == "" {
		c.Redirect(302, "/student/allCourse?msg=不存在的课程或错误的课程ID!")
		return
	}
	staff := module.CourseStaff{
		KID:     kid,
		UID:     uid,
		Student: student.Name,
	}
	if service.ChooseCourse(staff) {
		c.Redirect(302, "/student/mineCourse?msg=课程报名成功!")
	} else {
		c.Redirect(302, "/student/allCourse?msg=课程报名失败,请稍后重试或联系管理员确认!")
	}

}

func RejectCourse(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("uid")
	uid := value.(string)
	kid := c.Query("kid")
	if kid == "" {
		c.Redirect(302, "/student/mineCourse?msg=不存在的课程或错误的课程ID!")
		return
	}
	staff := module.CourseStaff{
		KID: kid,
		UID: uid,
	}
	if service.RejectCourse(staff) {
		c.Redirect(302, "/student/mineCourse?msg=退订课程成功!")
	} else {
		c.Redirect(302, "/student/mineCourse?msg=退订课程失败,请稍后重试或联系管理员确认!")
	}

}
