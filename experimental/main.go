package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/faithcomesbyhearing/language-api/language/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	lambda.Start(Handler)
}
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ConnectDatabase()
	return events.APIGatewayProxyResponse{
		StatusCode:        400,
		MultiValueHeaders: nil,
		Body:              "output",
		IsBase64Encoded:   true,
	}, nil
}

// func main() {
// 	ConnectDatabase()
// }

func ConnectDatabase() {
	fmt.Println("... ConnectDatabase")

	// r := &net.Resolver{
	// 	PreferGo: true,
	// Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
	// 	d := net.Dialer{
	// 		Timeout: time.Millisecond * time.Duration(10000),
	// 	}
	// 	return d.DialContext(ctx, network, "8.8.8.8:53")
	// },
	// }
	// ip, _ := r.LookupHost(context.Background(), "www.google.com")

	// print(ip[0])

	// nameserver, _ := net.LookupNS("db")
	// for _, ns := range nameserver {
	// 	addrs, errIP := net.LookupIP(ns.Host)
	// 	if errIP == nil {
	// 		fmt.Printf("%s\n", addrs[0].String())
	// 	} else {
	// 		fmt.Printf("%s\n", errIP.Error())
	// 	}
	// }

	// ips, err := net.LookupIP("db")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
	// 	os.Exit(1)
	// }
	// for _, ip := range ips {
	// 	fmt.Printf("google.com. IN A %s\n", ip.String())
	// }

	//time.Sleep(300 * time.Second)

	var dsn = util.Getenv("MYSQL_CONNECT_STRING", "root:password@tcp(docker.internal.host:3306)/LANGUAGE?charset=utf8mb4&parseTime=True&loc=Local")
	log.Printf("MYSQL_CONNECT_STRING: %s", dsn)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic("Failed to connect to database!" + err.Error())
	}
	defer db.Close()
	log.Printf("attempting ping")
	err = db.Ping()
	if err != nil {
		panic("Failed to ping database!" + err.Error())
	}
	log.Printf("success pinging datasource from go-provided SQL driver")

	log.Printf("success opening DB from go-provided SQL driver ")
	db.Close()

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	if database.Error != nil {
		panic("Failed to connect to database!")
	}
	log.Printf("success opening DB from gorm mysql dialector")
}
