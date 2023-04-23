package student

import (
	"StudentManger/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AskLeaveList(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("uid")
	uid := value.(string)
	student := service.FindStudentByUid(uid)
	ask_list := service.FindAskLeaveByUid(uid)
	c.HTML(200, "", gin.H{
		"name":    student.Name,
		"askList": ask_list,
	})
}
