package service

import (
	"context"

	v1 "goframe-project/api/v1"
)

type (
	sHello struct{}
)

func init() {
	RegisterHello(hello())
}

func hello() *sHello {
	return &sHello{}
}

func (s *sHello) GetHello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	name := req.Name
	if name == "" {
		name = "World"
	}
	
	res = &v1.HelloRes{
		Message: "Hello, " + name + "!",
	}
	return
}