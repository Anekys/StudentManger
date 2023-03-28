package admin

import (
	"StudentManger/module"
	"StudentManger/service"
	"StudentManger/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CourseInfoForm struct {
	KID       string `form:"kid"`      // 课程ID
	CName     string `form:"name"`     // 课程名
	CAbstract string `form:"abstract"` //课程简介
	Email     string `form:"email"`    //教师名
}

func courseMange(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("aid")
	uid := value.(string)
	admin := service.FindAdminByAid(uid)
	courses := service.FindAllCourse(1)
	c.HTML(200, "adminCourseMange.html", gin.H{
		"name":       admin.Name,
		"courseList": courses,
	})
}

func EditCourse(c *gin.Context) {
	kid := c.Query("kid")
	if kid == "" {
		// 新增课程
		c.HTML(200, "adminEditCourse.html", gin.H{})
	} else {
		// 编辑教师
		cour := service.FindCourseById(kid)
		c.HTML(200, "adminEditCourse.html", gin.H{
			"name":     cour.CName,
			"abstract": cour.CAbstract,
			"teacher":  cour.TName,
			"kid":      cour.KID,
		})
	}
}

func DeleteCourse(c *gin.Context) {
	kid := c.Query("kid")
	fmt.Println("删除课程:", kid)
	if kid == "" {
		// 新增课程
		c.Redirect(301, "/admin/courseMange?msg=缺少关键参数!")
	} else {
		// 编辑教师
		if service.DeleteCourseById(kid) {
			c.Redirect(301, "/admin/courseMange?msg=删除课程成功!")
		} else {
			c.Redirect(301, "/admin/courseMange?msg=删除失败!请重试!")
		}
	}
}

func PushCourse(c *gin.Context) {
	var courseForm CourseInfoForm
	err := c.ShouldBind(&courseForm)
	if courseForm.KID == "" {
		if err != nil {
			errStr := utils.GetVaildMsg(err, &courseForm)
			c.Redirect(301, "/admin/editCourse?msg="+errStr)
		} else {
			tid := utils.Md5Encrypt(courseForm.Email)
			teacher := service.FindTeacherByTid(tid)
			course := module.CourseInfo{
				CName:     courseForm.CName,
				CAbstract: courseForm.CAbstract,
				TID:       tid,
				TName:     teacher.Name,
			}
			if service.AddCourse(course) {
				c.Redirect(301, "/admin/courseMange")
			} else {
				c.Redirect(301, "/admin/courseTeacher?msg=添加课程失败!请检查后重试!")
			}
		}
	} else {
		if err != nil {
			errStr := utils.GetVaildMsg(err, &courseForm)
			c.Redirect(301, "/admin/editTeacher?tid="+courseForm.KID+"&msg="+errStr)
		} else {
			tid := utils.Md5Encrypt(courseForm.Email)
			teacher := service.FindTeacherByTid(tid)
			course := module.CourseInfo{
				CName:     courseForm.CName,
				CAbstract: courseForm.CAbstract,
				KID:       courseForm.KID,
				TID:       tid,
				TName:     teacher.Name,
			}
			if service.UpdateCourseById(course.KID, course) {
				c.Redirect(301, "/admin/courseMange")
			} else {
				c.Redirect(301, "/admin/editCourse?kid="+course.KID+"&msg=添加教师失败!请检查后重试!")
			}
		}
	}

}
