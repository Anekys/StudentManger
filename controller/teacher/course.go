package teacher

import (
	"StudentManger/module"
	"StudentManger/service"
	"StudentManger/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Course struct {
	CName     string `form:"name" msg:"请输入课程名"`      // 课程名
	CAbstract string `form:"abstract" msg:"请输入课程简介"` // 课程简介
	TName     string `form:"teacher_name" msg:"请输入任课教师姓名"`
	Kid       string `form:"kid" msg:"不存在的课程或课程ID错误!"`
}

func FindMineCourse(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("tid")
	tid := value.(string)
	teacher := service.FindTeacherByTid(tid)
	courses := service.FindTeacherCourse(tid, 1)
	c.HTML(200, "teacherMain.html", gin.H{
		"name":       teacher.Name,
		"courseList": courses,
		"title":      "我的课程",
	})
}

func FindAllCourse(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("tid")
	tid := value.(string)
	teacher := service.FindTeacherByTid(tid)
	courses := service.FindAllCourse(1)
	c.HTML(200, "teacherAllCourse.html", gin.H{
		"name":       teacher.Name,
		"courseList": courses,
		"title":      "全部课程",
	})
}

func EditCourseInfo(c *gin.Context) {
	kid := c.Query("kid")
	if kid == "" {
		c.Redirect(302, "/teacher/editCourse?msg=不存在的课程!")
		return
	}
	course := service.FindCourseById(kid)
	c.HTML(200, "teacherEditCourse.html", gin.H{
		"course": course,
	})
}

func UpdateCourseInfo(c *gin.Context) {
	var course Course
	err := c.ShouldBind(&course)
	if err != nil || course.Kid == "" {
		errors := utils.GetVaildMsg(err, &course)
		c.Redirect(302, "/teacher/main?msg="+errors)
		return
	}
	courseInfo := module.CourseInfo{
		KID:       course.Kid,
		CName:     course.CName,
		CAbstract: course.CAbstract,
		TName:     course.TName,
	}
	if service.UpdateCourseById(course.Kid, courseInfo) {
		c.Redirect(302, "/teacher/editCourse?kid="+course.Kid+"&msg=保存成功!")
	} else {
		c.Redirect(302, "/teacher/main?msg=保存失败!请稍后重试或联系管理员!")
	}
}
