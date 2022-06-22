package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/faithcomesbyhearing/language-api/language/util"
	_ "github.com/go-sql-driver/mysql"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("db handler")
	var dsn = util.Getenv("MYSQL_CONNECT_STRING", "root:password@tcp(host.docker.internal:3306)/LANGUAGE")
	log.Printf("MYSQL_CONNECT_STRING: %s", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Printf("attempting ping")
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("success pinging datasource")
	var (
		id       int
		name     string
		code     string
		iso3     string
		response string
	)
	rows, err := db.Query("select Id, name, code, iso3 from fcbhLanguage limit 10")
	if err != nil {
		log.Fatal(err)
		response := "fatal error (1): " + err.Error()
		return events.APIGatewayProxyResponse{Body: response, StatusCode: 500}, nil
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &code, &iso3)
		if err != nil {
			log.Fatal(err)
			response := "fatal error (2): " + err.Error()
			return events.APIGatewayProxyResponse{Body: response, StatusCode: 500}, nil
		}
		response += fmt.Sprintf("%d %s %s %s\n", id,
			name, code, iso3)
		log.Println(id, name, code, iso3)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		response := "fatal error (3): " + err.Error()
		return events.APIGatewayProxyResponse{Body: response, StatusCode: 500}, nil
	}
	return events.APIGatewayProxyResponse{Body: response, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
