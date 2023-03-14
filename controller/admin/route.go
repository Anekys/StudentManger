package admin

import (
	"StudentManger/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadRouter(group *gin.RouterGroup) {
	group.Use(AuthMiddleware())
	// 静态资源
	group.Static("/images", "web/images")
	group.Static("/fonts", "web/fonts")
	group.Static("/vendors", "web/vendors")
	group.Static("/css", "web/css")
	group.Static("/js", "web/js")
	//get请求
	group.GET("/main", Main)
	group.GET("/teacherMange", teacherMange)
	group.GET("/editTeacher", EditTeacher)
	group.GET("/deleteTeacher", deleteTeacher)
	group.GET("/courseMange", courseMange)
	group.GET("/editCourse", EditCourse)
	group.GET("/deleteCourse", DeleteCourse)
	//post请求
	//group.POST("/addStudent", AddStudent)
	//group.POST("/register", RegisterStudent)
	group.POST("/pushTeacher", PushTeacher)
	group.POST("/pushCourse", PushCourse)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("aid")
		if userID == nil {
			c.Redirect(http.StatusFound, "/login?msg=请先登录!")
			return
		}
		uid, _ := userID.(string)
		// 通过userID查询用户信息并存储在Context中
		adminStruct := service.FindAdminByAid(uid)
		c.Set("adminStruct", adminStruct)
		c.Next()
	}
}
