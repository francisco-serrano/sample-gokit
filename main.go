package main

import (
	"github.com/francisco-serrano/sample-gokit/service"
	"github.com/francisco-serrano/sample-gokit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
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
