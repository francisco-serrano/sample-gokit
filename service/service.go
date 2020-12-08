package service

type HelloService interface {
	Hello(v string) string
}

type helloService struct {
}

func NewHelloService() HelloService {
	return helloService{}
}

func (s helloService) Hello(v string) string {
	return "hello " + v
}
