package controller

import (
	"StudentManger/module"
	"StudentManger/service"
	"StudentManger/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Teacher struct {
	Name     string `form:"name" msg:"请填写姓名!" binding:"required"`
	Email    string `form:"email" msg:"请填写邮箱地址!" binding:"required"`
	PassWord string `form:"password" msg:"请填写密码!" binding:"required"`
}

func ForwardToTeacher(c *gin.Context) {
	uid := c.DefaultQuery("uid", "")
	teacher := service.FindTeacherByTid(uid)
	var empty = module.Teacher{}
	if teacher != empty {
		c.HTML(http.StatusOK, "teacher.html", gin.H{
			"name": teacher.Name,
		})
	}
}

func AddTeacher(c *gin.Context) {
	var teacher Teacher

	if err := c.ShouldBind(&teacher); err != nil {
		errStr := utils.GetVaildMsg(err, &teacher)
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  errStr,
		})
		return
	}
	Teacher := module.Teacher{
		Name:     teacher.Name,
		Email:    teacher.Email,
		PassWord: teacher.PassWord,
	}
	if service.AddTeacher(Teacher) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "老师添加成功",
		})
	} else {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "老师添加失败",
		})
	}

}
