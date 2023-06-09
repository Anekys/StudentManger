package controller

import (
	"StudentManger/controller/admin"
	"StudentManger/controller/public"
	"StudentManger/controller/student"
	"StudentManger/controller/teacher"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadRouter(router *gin.Engine) {
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("StudentManger"))
	router.Use(sessions.Sessions("SMSession", store))
	stu := router.Group("/student")
	student.LoadRouter(stu)
	tea := router.Group("/teacher")
	teacher.LoadRouter(tea)
	adm := router.Group("/admin")
	admin.LoadRouter(adm)
	//router.Use(Middleware())
	// get请求
	router.GET("/", func(c *gin.Context) {
		// 根据session中的结构,转向对应的main,全都不存在则转向登录
		c.Redirect(http.StatusMovedPermanently, "/login")
	})
	router.GET("/login", public.ForwardToLogin)
	router.GET("/logout", public.Logout)
	router.GET("/register", public.ForwardToRegister)
	// post请求
	router.POST("/login", public.Login)
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		session := sessions.Default(c)
		fmt.Println("Request_URL_Path Is", path)
		if path == "/login" {
			return
		}
		// 浏览器会默认请求favicon.ico
		if result, err := c.Cookie("login"); err == nil {
			flag, boolen := c.Get("login")
			fmt.Println("获取到的flag", flag, boolen, session)
			if result != "true" {
				c.Redirect(http.StatusMovedPermanently, "/login?msg=请先登录!")
			}
		} else {

			c.Redirect(http.StatusMovedPermanently, "/login?msg=请先登录!")
		}
		//if utils.NotNeedCookie(path) {
		//	return
		//}
		//_, err := c.Cookie("login")
		//if err != nil {
		//	c.Redirect(http.StatusMovedPermanently, "/login?msg=请先登录!")
		//}else {
		//	user := module.User{}
		//	user.Username = cookie
		//	fmt.Println(cookie)
		//	result := service.FindUserByUserNameFirst(user)
		//	empty := module.User{}
		//	if result == empty {
		//		c.Redirect(http.StatusMovedPermanently, "/login?msg=用户不存在")
		//	}
		//}
	}
}
