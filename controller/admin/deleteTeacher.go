package admin

import (
	"StudentManger/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func deleteTeacher(c *gin.Context) {
	tid := c.Query("tid")
	fmt.Println("删除教师：", tid)
	if tid == "" {
		c.Redirect(301, "/admin/teacherMange?msg="+"错误的教师ID!")
	} else {
		if service.DeleteTeacher(tid) {
			c.Redirect(301, "/admin/teacherMange?msg="+"删除失败!请稍后重试!")
		} else {
			c.Redirect(301, "/admin/teacherMange?msg="+"删除成功!")
		}
	}
}
