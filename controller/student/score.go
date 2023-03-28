package student

import (
	"StudentManger/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ScoreMange(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("uid")
	uid := value.(string)
	student := service.FindStudentByUid(uid)
	scoreList := service.FindMineScore(uid)
	c.HTML(200, "studentMineScore.html", gin.H{
		"title":     "成绩管理",
		"name":      student.Name,
		"scoreList": scoreList,
	})
}
