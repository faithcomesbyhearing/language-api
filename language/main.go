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

	r.GET("/language", controllers.FindLanguages)
	r.GET("/language/:id", controllers.GetLanguage)
	r.POST("/language", controllers.AddLanguage)
	r.PATCH("/language/:id", controllers.UpdateLanguage)

	r.GET("/language/:id/note", controllers.FindLanguageNotes)
	r.GET("/language/:id/note/:note_id", controllers.GetLanguageNote)
	r.POST("/language/:id/note", controllers.AddLanguageNote)
	//r.PATCH("/language/:id/note/:note-id", controllers.UpdateLanguageNote)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
