package controller

import (
	"StudentManger/module"
	"StudentManger/service"
	"StudentManger/utils"
	"github.com/gin-gonic/gin"
)

type ConfirmInfo struct {
	Kid string `form:"kid"`
	Tid string `form:"tid"`
}

func AddConfirmInfo(c *gin.Context) {
	var confirm ConfirmInfo
	if err := c.ShouldBind(&confirm); err != nil {
		errStr := utils.GetVaildMsg(err, &confirm)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  errStr,
		})
		return
	}
	var key = []string{
		confirm.Kid, confirm.Tid, utils.GetNowTimeStamp(),
	}
	secret := utils.CreateSecret(key)
	confirmInfo := module.ConfirmInfo{
		KID:    confirm.Kid,
		TID:    confirm.Tid,
		Secret: secret,
	}
	if service.AddConfirmInfo(confirmInfo) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "开始考勤成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "开始考勤失败",
		})
	}
}

func FindConfirmByKid(c *gin.Context) {
	kid := c.DefaultQuery("kid", "")
	if kid == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "需要kid参数",
		})
		return
	}
	confirm := service.FindConfirmInfoByKid(kid)
	var empty module.ConfirmInfo
	if confirm == empty {
		c.JSON(200, gin.H{
			"code":   200,
			"msg":    "该课程尚未开始考勤",
			"status": 0, //未考勤
		})
		return
	}
	c.JSON(200, gin.H{
		"code":   200,
		"secret": confirm.Secret,
		"status": 1, // 正在考勤
	})
}
