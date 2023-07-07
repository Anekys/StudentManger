package student

import (
	"StudentManger/module"
	"StudentManger/service"
	"StudentManger/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AddALForm struct {
	Lid  string `form:"lid"`
	Name string `form:"name" msg:"请输入请假人" binding:"required"`
	Case string `form:"case" msg:"请输入请假原因" binding:"required"`
}

func AskLeaveList(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("uid")
	uid := value.(string)
	student := service.FindStudentByUid(uid)
	ask_list := service.FindAskLeaveByUid(uid, 1)
	c.HTML(200, "studentAskLeaveMange.html", gin.H{
		"name":    student.Name,
		"askList": ask_list,
	})
}
func pushAskLeavePage(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("uid")
	uid := value.(string)
	student := service.FindStudentByUid(uid)
	askLeave := module.AskLeave{
		UID:  uid,
		Name: student.Name,
	}
	c.HTML(200, "studentAddAskLeave.html", gin.H{
		"name": student.Name,
		"ask":  askLeave,
	})
}

func pushAskLeaveRequest(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("uid")+
	uid := value.(string)
	student := service.FindStudentByUid(uid)
	var form AddALForm
	if err := c.ShouldBind(&form); err != nil {
		errStr := utils.GetVaildMsg(err, &form)
		c.Redirect(301, "/student/addAskLeave?msg="+errStr)
		return
	}
	lid := ""
	flag := false
	if form.Lid == "" {
		lid = utils.Md5Encrypt(uid + utils.GetNowTimeStamp())
		flag = false
	} else {
		lid = form.Lid
		flag = true
	}

	askLeave := module.AskLeave{
		LID:    lid,
		UID:    uid,
		Name:   student.Name,
		Cause:  form.Case,
		Status: 0,
	}
	result := true
	if !flag {
		result = service.AddAskLeave(askLeave)
	} else {
		result = service.UpdateAskLeaveByLid(lid, askLeave)
	}
	if result {
		c.Redirect(301, "/student/askLeave?msg=提交成功!")
	} else {
		c.Redirect(301, "/student/askLeave?msg=提交失败,请稍后重试!")
	}
}

func editAskLeave(c *gin.Context) {
	lid := c.Query("lid")
	ask := service.FindAskLeaveByLid(lid)
	c.HTML(200, "studentAddAskLeave.html", gin.H{
		"name": ask.Name,
		"ask":  ask,
	})
}
func viewDetail(c *gin.Context) {
	lid := c.Query("lid")
	ask := service.FindAskLeaveByLid(lid)
	c.HTML(200, "studentViewAskLeave.html", gin.H{
		"name": ask.Name,
		"ask":  ask,
	})
}
