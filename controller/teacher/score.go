package teacher

import (
	"StudentManger/module"
	"StudentManger/service"
	"StudentManger/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// ScoreForm 老师提交学生分数
type ScoreForm struct {
	UID     string `form:"uid" msg:"学生信息错误!请检查学生uid是否正确"`    // 学生ID
	Student string `form:"student" msg:"学生信息错误!请检查学生姓名是否正确"` // 学生名称
	KID     string `form:"kid" msg:"课程信息错误!请检查课程kid是否正确"`    // 课程ID
	Course  string `form:"course" msg:"课程信息错误!请检查课程信息是否正确"`  // 课程名
	Score   string `form:"score" msg:"分数信息错误!请检查分数是否符合格式"`   // 成绩
}

func ScoreMange(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("tid")
	tid := value.(string)
	teacher := service.FindTeacherByTid(tid)
	courses := service.FindTeacherCourse(tid, 1)
	c.HTML(200, "teacherScoreMange.html", gin.H{
		"name":       teacher.Name,
		"courseList": courses,
		"title":      "成绩管理",
	})
}

func viewStaff(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("tid")
	tid := value.(string)
	kid := c.Query("kid")
	teacher := service.FindTeacherByTid(tid)
	course := service.FindCourseById(kid)
	staffList := service.FindCourseAllStaffById(kid, 1)
	c.HTML(200, "teacherScoreStaff.html", gin.H{
		"title":     "人员详情",
		"name":      teacher.Name,
		"staffList": staffList,
		"kid":       kid,
		"tid":       tid,
		"course":    course.CName,
	})
}

func saveScore(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("tid")
	tid := value.(string)
	teacher := service.FindTeacherByTid(tid)
	var scoreForm ScoreForm
	err := c.ShouldBind(&scoreForm)
	if err != nil {
		errStr := utils.GetVaildMsg(err, &scoreForm)
		c.JSON(200, gin.H{
			"status": 400,
			"msg":    errStr,
		})
		return
	}
	score := module.Score{
		UID:     scoreForm.UID,
		Student: scoreForm.Student,
		TID:     tid,
		Teacher: teacher.Name,
		KID:     scoreForm.KID,
		Course:  scoreForm.Course,
		Score:   scoreForm.Score,
	}
	flag := true
	if service.FindScoreExist(score.KID, score.UID) {
		flag = service.UpdatesScoreByKidAndUid(score)
	} else {
		flag = service.AddScore(score)
	}
	if flag {
		c.JSON(200, gin.H{
			"status": 200,
			"msg":    "成绩保存成功",
		})
	} else {
		c.JSON(200, gin.H{
			"status": 400,
			"msg":    "成绩保存失败,请稍后重试或联系管理员!",
		})
	}
}
