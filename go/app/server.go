package app

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"web/log"
	"web/util"
	"xutil/xLog"
	"xutil/xStr"
)

var httpServer *gin.Engine
var HTTP_V1 *gin.RouterGroup

//初始化服务
func init() {
	httpServer = gin.Default()

	HTTP_V1 = httpServer.Group("/v1")
	HTTP_V1.GET("/time", func(c *gin.Context) {
		c.String(200, xStr.Int2str(util.GetTime()))
		//Response(c, 200, )
	})

	//开启API文档
	url := ginSwagger.URL("./swagger/doc.json")
	httpServer.GET("/api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	//开启静态资源访问以及GZIP
	httpServer.Use(gzip.Gzip(gzip.DefaultCompression))
	httpServer.Static("/admin", "./dist")
	httpServer.Static("/static", "./dist/static")
	httpServer.StaticFile("/favicon.ico", "./dist/favicon.ico")
	//处理非定义的请求
	httpServer.NoRoute(func(c *gin.Context) {
		Response(c, 400, "请求非法")
	})
}

//启动服务
func Start(pro string) {
	err := httpServer.Run(":" + pro)
	if err != nil {
		glog.GetLog().WithFields(xLog.Fields{"ERR": err}).Fatal("HTTP服务启动失败")
	}

	glog.GetLog().INFO("HTTP服务启动成功")
}
