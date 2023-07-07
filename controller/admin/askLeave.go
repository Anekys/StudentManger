package admin

import (
	"StudentManger/module"
	"StudentManger/service"
	"StudentManger/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type askForm struct {
	LID    string `form:"lid" msg:"LID错误!" binding:"required"`      // 假条的ID
	Reason string `form:"reason" msg:"请输入审批意见!" binding:"required"` // 处理原因
	Status int    `form:"status" msg:"请选择审批结果!" binding:"required"` // 状态：0.待处理,1.批准2.拒绝
}

func AskLeaveMange(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("aid")
	uid := value.(string)
	admin := service.FindAdminByAid(uid)
	ask_list := service.FindAllAskLeave(1)
	c.HTML(200, "adminAskLeaveMange.html", gin.H{
		"name":    admin.Name,
		"askList": ask_list,
	})
}

func handleAskLeave(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("aid")
	uid := value.(string)
	admin := service.FindAdminByAid(uid)
	lid := c.Query("lid")
	ask := service.FindAskLeaveByLid(lid)
	c.HTML(200, "adminHandleAskLeave.html", gin.H{
		"name": admin.Name,
		"ask":  ask,
	})
}

func PushAskLeave(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("aid")
	uid := value.(string)
	admin := service.FindAdminByAid(uid)
	var form askForm
	err := c.ShouldBind(&form)
	if err != nil {
		errStr := utils.GetVaildMsg(err, &form)
		c.Redirect(301, "/admin/askLeaveMange?msg="+errStr)
		return
	}
	askLeave := module.AskLeave{
		LID:    form.LID,
		AID:    admin.AID,
		AName:  admin.Name,
		Reason: form.Reason,
		Status: form.Status,
	}
	if service.UpdateAskLeaveByLid(form.LID, askLeave) {
		c.Redirect(301, "/admin/askLeaveMange?msg=审批成功!")
	} else {
		c.Redirect(301, "/admin/askLeaveMange?msg=审批失败!")
	}
}
