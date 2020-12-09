package main

//go:generate sqlboiler --wipe mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/francisco-serrano/sample-gokit/models"
	"github.com/francisco-serrano/sample-gokit/service"
	"github.com/francisco-serrano/sample-gokit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", `root:root@tcp(localhost:3306)/samplegokit?parseTime=true`)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected")

	msg := &models.Message{
		Content: "hello world",
	}

	if err := msg.Insert(context.Background(), db, boil.Infer()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("msg id:", msg.ID)

	got, err := models.Messages().One(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("got msg id:", got.ID)

	found, err := models.FindMessage(context.Background(), db, msg.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("found msg id:", found.ID)

	found.Content = "hello world 2"
	rows, err := found.Update(context.Background(), db, boil.Whitelist(models.MessageColumns.Content))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("updated", rows, "messages")

	foundAgain, err := models.Messages(qm.Where("content = ?", found.Content)).One(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("found again:", foundAgain.ID, foundAgain.Content)

	exists, err := models.MessageExists(context.Background(), db, foundAgain.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("user", foundAgain.ID, "exists?", exists)

	count, err := models.Messages().Count(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("there are", count, "users")

	messages, err := models.Messages().All(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("got:", len(messages), "messages")

	for _, m := range messages {
		fmt.Println("   got message:", m.ID)
	}

	svc := service.NewHelloService()

	helloHandler := httptransport.NewServer(
		transport.MakeHelloEndpoint(svc),
		transport.DecodeHelloRequest,
		transport.EncodeResponse,
	)

	app := fiber.New()
	app.Get("/hello", adaptor.HTTPHandler(helloHandler))

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
