package controller

import (
	"StudentManger/module"
	"StudentManger/service"
	"StudentManger/utils"
	"github.com/gin-gonic/gin"
)

type Course struct {
	CName     string `form:"name" msg:"请输入课程名"`      // 课程名
	CAbstract string `form:"abstract" msg:"请输入课程简介"` // 课程简介
	TID       string `form:"tid" msg:"请输入任课教师"`
}
type studentCourse struct {
	KID string `form:"kid" msg:"请指定课程ID"`
	UID string `form:"uid" msg:"请指定学生ID"`
}

func AddCourseInfo(c *gin.Context) {
	var course Course
	if err := c.ShouldBind(&course); err != nil {
		errStr := utils.GetVaildMsg(err, &course)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  errStr,
		})
		return
	}
	courseInfo := module.CourseInfo{
		CName:     course.CName,
		CAbstract: course.CAbstract,
		TID:       course.TID,
	}
	if service.AddCourse(courseInfo) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "添加课程成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "添加课程失败",
		})
	}
}

func AddCourseStudent(c *gin.Context) {
	var stu studentCourse
	if err := c.ShouldBind(&stu); err != nil {
		errStr := utils.GetVaildMsg(err, &stu)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  errStr,
		})
		return
	}
	Stu := module.CourseStaff{
		KID: stu.KID,
		UID: stu.UID,
	}
	if service.AddCourseStudent(Stu) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "添加上课人员成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "添加上课人员失败",
		})
	}

}
