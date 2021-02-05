package router

import (
	// 导入session包
	"github.com/gin-contrib/sessions"
	// 导入session存储引擎
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv29/controller"
	"github.com/liuhongdi/digv29/global"
	"log"
	"net/http"
	"runtime/debug"
)

func Router() *gin.Engine {
	router := gin.Default()
	//处理异常
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	//use middleware
	router.Use(Recover)

	// 基于cookie创建session的存储引擎，传递一个参数，用来做加密时的密钥
	store := cookie.NewStore([]byte("secret1234"))
	//session中间件生效，参数mysession，是浏览器端cookie的名字
	router.Use(sessions.Sessions("mysession", store))

	//static
	router.StaticFS("/static", http.Dir("/data/liuhongdi/digv29/static"))

	// 路径映射:index
	userc:=controller.NewUserController()
	router.POST("/user/login", userc.Login);
	router.GET("/user/session", userc.Session);
	router.GET("/user/logout", userc.Logout);
	return router
}

func HandleNotFound(c *gin.Context) {
	global.NewResult(c).Error(404,"资源未找到")
	return
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			global.NewResult(c).Error(500,"服务器内部错误")
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}