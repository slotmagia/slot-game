package service

import (
	"context"

	v1 "goframe-project/api/v1"
)

type (
	IHello interface {
		GetHello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error)
	}
)

var (
	localHello IHello
)

func Hello() IHello {
	if localHello == nil {
		panic("implement not found for interface IHello, forgot register?")
	}
	return localHello
}

func RegisterHello(i IHello) {
	localHello = i
}