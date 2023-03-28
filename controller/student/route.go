package student

import (
	"StudentManger/service"
	"fmt"
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
	group.GET("/allCourse", FindAllCourse)
	group.GET("/mineCourse", FindChooseCourse)
	group.GET("/chooseCourse", ChooseCourse)
	group.GET("/rejectCourse", RejectCourse)
	group.GET("/confirmMange", ConfirmMange)
	group.GET("/startConfirm", StartConfirm)
	group.GET("/scoreMange", ScoreMange)
	//post请求
	group.POST("/addStudent", AddStudent)
	group.POST("/confirm", Confirm)
	group.POST("/register", RegisterStudent)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("uid")
		fmt.Println(userID, c.Request.URL)
		if userID == nil {
			c.Redirect(http.StatusFound, "/login?msg=请先登录!")
			return
		}
		uid, _ := userID.(string)
		// 通过userID查询用户信息并存储在Context中
		studentStruct := service.FindStudentByUid(uid)
		c.Set("studentStruct", studentStruct)
		c.Next()
	}
}
