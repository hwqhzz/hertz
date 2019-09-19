package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hertz/pkg/setting"
	"hertz/routes"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "hertz/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {

	//设置框架启动模式
	gin.SetMode(setting.RunMode)

	//获取路由
	route := gin.Default()

	routes.CreateRoutes(route)
	if (setting.AppDebug) {
		url := ginSwagger.URL("http://localhost:9000/swagger/doc.json") // The url pointing to API definition
		route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	//配置服务
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        route,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	//服务启动
	server.ListenAndServe()
}
