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
		group.POST("/course/modify/:uuid", Auth(), v1.ModifyCourse)
		group.POST("/course/delete/:uuid", Auth(), v1.DeleteCourse)
		group.POST("/course", Auth(), v1.CreateCourse)

		group.GET("/image", Auth(), v1.GetImageList)
		group.GET("/image/:uuid", Auth(), v1.GetSpecificImage)
		group.PUT("/image/:uuid", Auth(), v1.ModifyImage)
		group.POST("/image/delete/:uuid", Auth(), v1.DeleteImage)
		group.POST("/image", Auth(), v1.CreateImage)

		group.GET("/vm", Auth(), v1.GetVMList)
		group.GET("/vm/:uuid", Auth(), v1.GetSpecificVM)
		group.GET("/vm/vnc/:uuid", Auth(), v1.GetVNCPort)
		group.PUT("/vm/:uuid", Auth(), v1.ModifyVM)
		group.POST("/vm/delete/:uuid", Auth(), v1.DeleteVM)
		group.GET("/vm/memory/:uuid", Auth(), v1.GetMemoryUsage)
		group.POST("/vm", Auth(), v1.CreateVM)
		group.POST("/vm/shutdown/:uuid", Auth(), v1.ShutDownVM)
		group.POST("/vm/reboot/:uuid", Auth(), v1.RebootVM)
		group.POST("/vm/start/:uuid", Auth(), v1.StartVM)
		group.POST("/vm/image/:uuid", Auth(), v1.MakeImageWithVM)
		group.POST("/vm/snapshot", Auth(), v1.MakeSnapshotWithVM)
		group.PATCH("/vm/snapshot", Auth(), v1.ResetVMWithSnapshot)

	}
	return r
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //????????????
		origin := c.Request.Header.Get("Origin") //????????????
		var headerKeys []string                  // ???????????????keys
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
			c.Header("Access-Control-Allow-Origin", "*")                                       // ???????????????????????????
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //?????????????????????????????????????????????,????????????????????????????????????'??????'??????
			//  header?????????
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              ??????????????????                                                                                                      ????????????????????????
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // ?????????????????? ????????????????????????
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // ?????????????????? ????????????
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  ???????????????????????????cookie?????? ???????????????true
			c.Set("content-type", "application/json")                                                                                                                                                              // ?????????????????????json
		}
		//????????????OPTIONS??????
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// ????????????
		c.Next() //  ????????????
	}
}

// Auth ????????????????????????
// ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????
// 1. ????????????????????????cookie????????????????????????????????????????????????
// 2. ????????????????????????????????????????????????????????????cookie?????????????????????????????????uuid????????????????????????session???????????????uuid???key?????????????????????????????????
// 3. ????????????????????????????????????????????????cookie??????????????? Auth() ???????????????????????????????????????????????????
// 4. ?????????????????????????????????????????????????????????????????????????????????????????????api???????????????????????????????????????
// ???????????????????????????????????????????????????????????????????????????????????????cookie????????????????????????????????????????????????cookie?????????????????????????????????????????????????????????
// ???????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????
//
// ?????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????
// ??????????????????get post delete ???????????????????????????????????????????????????????????????????????????
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
