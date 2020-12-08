package service

import (
	"errors"
	"strings"
)

type HelloService interface {
	Hello(v string) (string, error)
}

type helloService struct {
}

func NewHelloService() HelloService {
	return helloService{}
}

func (s helloService) Hello(v string) (string, error) {
	if strings.TrimSpace(v) == "" {
		return "", errors.New("cannot pass empty string")
	}

	return "hello " + v, nil
}
