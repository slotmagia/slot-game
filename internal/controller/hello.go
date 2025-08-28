package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "goframe-project/api/v1"
	"goframe-project/internal/service"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) GetHello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.WriteHeader(200)
	return service.Hello().GetHello(ctx, req)
}