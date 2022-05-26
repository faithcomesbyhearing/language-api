package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/workspaces/language-api/language/controllers"
)

var ginLambda *ginadapter.GinLambda

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("Gin cold start")
		r := gin.Default()

		r.Use(controllers.Cors())

		r.POST("/", controllers.PostLanguage)
		r.PUT("/:id", controllers.UpdateLanguage)
		r.GET("/:id", controllers.GetLanguageDetail)

		r.GET("/", controllers.GetLanguage)

		ginLambda = ginadapter.New(r)
	}

	// return ginLambda.ProxyWithContext(ctx, req)
	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
