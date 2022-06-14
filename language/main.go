package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/faithcomesbyhearing/language-api/language/controllers"
	"github.com/faithcomesbyhearing/language-api/language/models"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	r := gin.Default()
	models.ConnectDatabase()
	//r.Use(controllers.Cors())

	r.GET("/", controllers.FindLanguages)
	r.GET("/language", controllers.FindLanguages)
	r.GET("/language/:id", controllers.GetLanguage)
	// r.POST("/", controllers.PostLanguage)
	// r.PUT("/:id", controllers.UpdateLanguage)

	// r.NoRoute(func(c *gin.Context) {
	// 	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	// })

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
	//return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
