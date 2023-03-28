package teacher

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
	group.GET("/main", FindMineCourse)
	group.GET("/allCourse", FindAllCourse)
	group.GET("/editCourse", EditCourseInfo)
	group.GET("/confirmMange", ConfirmMange)
	group.GET("/startConfirm", StartConfirm)
	group.GET("/endConfirm", EndConfirm)
	group.GET("/confirmResult", ConfirmResultMange)
	group.GET("/viewConfirm", ConfirmDetail)
	group.GET("/scoreMange", ScoreMange)
	group.GET("/viewStaff", viewStaff)
	//group.GET("/main")
	//post请求
	group.POST("/updateCourse", UpdateCourseInfo)
	group.POST("/saveScore", saveScore)
	//group.POST("/register", RegisterStudent)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("tid")
		//fmt.Println(userID, c.Request.URL)
		if userID == nil {
			c.Redirect(http.StatusFound, "/login?msg=请先登录!")
			return
		}
		uid, _ := userID.(string)
		// 通过userID查询用户信息并存储在Context中
		teacherStruct := service.FindTeacherByTid(uid)
		c.Set("teacherStruct", teacherStruct)
		c.Next()
	}
}
