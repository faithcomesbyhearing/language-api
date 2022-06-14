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

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("Gin cold start")
		r := gin.Default()
		models.ConnectDatabase()
		//r.Use(controllers.Cors())

		r.GET("/language", controllers.FindLanguages)
		r.GET("/language/:id", controllers.GetLanguage)
		// r.POST("/", controllers.PostLanguage)
		// r.PUT("/:id", controllers.UpdateLanguage)

		ginLambda = ginadapter.New(r)
	}

	return ginLambda.ProxyWithContext(ctx, req)
	//return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
