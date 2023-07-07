package admin

import (
	"StudentManger/module"
	"StudentManger/service"
	"StudentManger/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Admin struct {
	Aid      string `form:"aid"`
	Name     string `form:"aname" msg:"请填写姓名!" binding:"required"`
	Email    string `form:"email" msg:"请填写邮箱地址!" binding:"required"`
	PassWord string `form:"password" msg:"请填写密码!" binding:"required"`
}

func adminMange(c *gin.Context) {
	session := sessions.Default(c)
	value := session.Get("aid")
	uid := value.(string)
	admin := service.FindAdminByAid(uid)
	admins := service.FindAllAdmins(1)
	c.HTML(http.StatusOK, "adminAdminMange.html", gin.H{
		"name":      admin.Name,
		"adminList": admins,
	})
}

func editAdmin(c *gin.Context) {
	aid := c.Query("aid")
	session := sessions.Default(c)
	value := session.Get("aid")
	uid := value.(string)
	admin := service.FindAdminByAid(uid)
	if aid == "" {
		// 新增教师
		c.HTML(200, "adminEditAdmin.html", gin.H{
			"name": admin.Name,
		})
	} else {
		// 编辑教师
		adm := service.FindAdminByAid(aid)
		c.HTML(200, "adminEditAdmin.html", gin.H{
			"name":     admin.Name,
			"aname":    adm.Name,
			"email":    adm.Email,
			"password": adm.PassWord,
			"aid":      adm.AID,
		})
	}
}

func PushAdmin(c *gin.Context) {
	var admin Admin
	err := c.ShouldBind(&admin)
	if admin.Aid == "" {
		if err != nil {
			errStr := utils.GetVaildMsg(err, &admin)
			c.Redirect(301, "/admin/editAdmin?msg="+errStr)
		} else {
			aid := utils.Md5Encrypt(admin.Email)
			adm := module.Admin{
				AID:      aid,
				Name:     admin.Name,
				Email:    admin.Email,
				PassWord: admin.PassWord,
			}
			if service.AddAdmin(adm) {
				c.Redirect(301, "/admin/adminMange")
			} else {
				c.Redirect(301, "/admin/editAdmin?msg=添加管理员失败!请检查后重试!")
			}
		}
	} else {
		if err != nil {
			errStr := utils.GetVaildMsg(err, &admin)
			c.Redirect(301, "/admin/editAdmin?aid="+admin.Aid+"&msg="+errStr)
		} else {
			adm := module.Admin{
				AID:      admin.Aid,
				Name:     admin.Name,
				Email:    admin.Email,
				PassWord: admin.PassWord,
			}
			if service.UpdateAdminByAid(admin.Aid, adm) {
				c.Redirect(301, "/admin/adminMange")
			} else {
				c.Redirect(301, "/admin/editAdmin?aid="+admin.Aid+"&msg=添加管理员失败!请检查后重试!")
			}
		}
	}
}

func deleteAdmin(c *gin.Context) {
	aid := c.Query("aid")
	fmt.Println("删除管理员：", aid)
	if aid == "" {
		c.Redirect(301, "/admin/adminMange?msg="+"错误的管理员ID!")
	} else {
		if service.DeleteAdminByAid(aid) {
			c.Redirect(301, "/admin/adminMange?msg="+"删除成功!")
		} else {
			c.Redirect(301, "/admin/adminMange?msg="+"删除失败!请稍后重试!")
		}
	}
}
