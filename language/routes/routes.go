package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/workspaces/language-api/language/controllers"
)

var Router *gin.Engine

func CreateRoutes() *gin.Engine {
	Router = gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	Router.Use(controllers.Cors())
	// v1 of the API
	v1 := Router.Group("")
	{
		v1.POST("/language", controllers.PostLanguage)
		v1.PUT("/language/:id", controllers.UpdateLanguage)
		v1.GET("/language/:id", controllers.GetLanguageDetail)

		v1.GET("/language/", controllers.GetLanguage)

	}
	return Router
}
