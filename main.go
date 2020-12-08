package main

import (
	"github.com/francisco-serrano/sample-gokit/service"
	"github.com/francisco-serrano/sample-gokit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
)

func main() {
	svc := service.NewHelloService()

	helloHandler := httptransport.NewServer(
		transport.MakeHelloEndpoint(svc),
		transport.DecodeHelloRequest,
		transport.EncodeResponse,
	)

	http.Handle("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
