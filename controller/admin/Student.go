package admin

import (
	"StudentManger/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func deleteStudent(c *gin.Context) {
	uid := c.Query("uid")
	fmt.Println("删除学生：", uid)
	if uid == "" {
		c.Redirect(301, "/admin/main?msg="+"错误的学生ID!")
	} else {
		if service.DelStudentByUID(uid) {
			c.Redirect(301, "/admin/main?msg="+"删除成功!")
		} else {
			c.Redirect(301, "/admin/main?msg="+"删除失败!请稍后重试!")
		}
	}
}
