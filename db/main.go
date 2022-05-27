package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("db handler")

	db, err := sql.Open("mysql",
		"root:password@tcp(host.docker.internal:3306)/LANGUAGE")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	var (
		id   int
		name string
	)
	rows, err := db.Query("select Id, name, code, iso3 from fcbhLanguage where Id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return events.APIGatewayProxyResponse{Body: "response from DB endpoint", StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
