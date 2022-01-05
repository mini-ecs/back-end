package router

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/mini-ecs/back-end/api/v1"
	"github.com/mini-ecs/back-end/pkg/common/error_msg"
	"github.com/mini-ecs/back-end/pkg/common/response"
	"github.com/mini-ecs/back-end/pkg/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strings"
)

var logger = log.GetGlobalLogger()

func NewRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(Cors())
	store := cookie.NewStore([]byte("secret12345"))
	r.Use(sessions.Sessions("mini-ecs", store))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	group := r.Group("api/v1")
	{
		group.GET("/welcome", Auth(), v1.Welcome)

		group.POST("/user/login", v1.Login)
		group.POST("/user/register", v1.RegisterUser)
		group.POST("/user/modify", Auth(), v1.ModifyUser)
		group.GET("/user/currentUser", Auth(), v1.CurrentUser)

		group.GET("/course", Auth(), v1.GetCourseList)
		group.GET("/course/configs", Auth(), v1.GetMachineConfig)
		group.GET("/course/:uuid", Auth(), v1.GetSpecificCourse)
		group.PUT("/course/:uuid", Auth(), v1.ModifyCourse)
		group.DELETE("/course/:uuid", Auth(), v1.DeleteCourse)
		group.POST("/course", Auth(), v1.CreateCourse)

		group.GET("/image", Auth(), v1.GetImageList)
		group.GET("/image/:uuid", Auth(), v1.GetSpecificImage)
		group.PUT("/image/:uuid", Auth(), v1.ModifyImage)
		group.DELETE("/image/:uuid", Auth(), v1.DeleteImage)
		group.POST("/image", Auth(), v1.CreateImage)

		group.GET("/vm", Auth(), v1.GetVMList)
		group.GET("/vm/:uuid", Auth(), v1.GetSpecificVM)
		group.PUT("/vm/:uuid", Auth(), v1.ModifyVM)
		group.DELETE("/vm/:uuid", Auth(), v1.DeleteVM)
		group.POST("/vm", Auth(), v1.CreateVM)
		group.POST("/vm/image", Auth(), v1.MakeImageWithVM)
		group.POST("/vm/snapshot", Auth(), v1.MakeSnapshotWithVM)
		group.PATCH("/vm/snapshot", Auth(), v1.ResetVMWithSnapshot)

	}
	return r
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}

// Auth 关于登录授权问题
// 目前的实现方式只能保证在登录时，客户端和后端都能保证用户是在线的。检验的方式为：
// 1. 登录时，前端不带cookie地向后端验证用户名和密码是否正确
// 2. 后端验证正确，则向客户端返回的消息里带上cookie，里面存放的是该用户的uuid。后端本地维护了session，以用户的uuid为key来标志用户的状态为在线
// 3. 此后，客户端每次发送请求都附带该cookie，后端使用 Auth() 函数来确保用户发来的请求是他本人。
// 4. 客户端方面则在本地维护了一个变量来标志用户是否在线。即使用一条api请求后端确认用户是否在线。
// 在正常的情况下上面的逻辑没有问题，但如果用户在登录后清除了cookie，则会导致用户后续的请求都是不带cookie的，后端则将其标记为不在线。但前端自己
// 本身维护的状态里，因为他已经在登录时验证了自己是在线的，所以它认为自己此时的请求是合法的。因此会出现问题。
//
// 一种比较合适的解决方式是在前端也使用中间件，每次收到后盾的返回消息时先检测后端返回的是否是标志自己不在线的消息，再进行后续操作。如果不这样做，
// 则客户端每个get post delete 操作都需要加上一段代码来执行上面的逻辑，非常冗余。
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cook, err := c.Cookie("uuid")
		if err != nil {
			c.JSON(http.StatusOK, response.FailCodeMsg(error_msg.ErrorUnauthorized, "You should login first"))
			//c.Redirect(http.StatusMovedPermanently, "http://localhost:8000/user/login")
			c.Abort()
		}
		session := sessions.Default(c)
		if session.Get(cook) != "online" {
			c.JSON(http.StatusOK, response.FailCodeMsg(error_msg.ErrorUnauthorized, "You should login first"))
			//c.Redirect(http.StatusMovedPermanently, "http://localhost:8000/user/login")
			c.Abort()
		}
	}
}
