package routes

/**
 * 配置路由
 */
import (
	"github.com/qit-team/snow/app/http/controllers"
	"github.com/qit-team/snow/app/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/qit-team/snow-core/http/middleware"
)

//api路由配置
func RegisterRoute(router *gin.Engine) {
	//middleware: 服务错误处理 => 生成请求id => access log
	router.Use(middlewares.ServerRecovery(), middleware.GenRequestId, middleware.GenContextKit, middleware.AccessLog())

	router.NoRoute(controllers.Error404)
	router.GET("/hello", controllers.HandleHello)
	router.POST("/test", controllers.HandleTest)

	//api版本
	v1 := router.Group("/v1")
	{
		v1.GET("/banner_list", controllers.GetBannerList)
	}
}
