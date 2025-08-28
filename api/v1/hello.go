package v1

import "github.com/gogf/gf/v2/frame/g"

type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"Hello World API"`
	Name   string `json:"name" dc:"Name parameter"`
}

type HelloRes struct {
	g.Meta  `mime:"application/json"`
	Message string `json:"message" dc:"Hello message"`
}