// Code generated by hertz generator.

package hello

import (
	"context"

	hello "apigateway1/biz/model/hello"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// HelloMethod .
// @router /hello [GET]
func HelloMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hello.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(hello.HelloResp)
	resp.RespBody = "hello," + req.Name // 添加的逻辑

	c.JSON(consts.StatusOK, resp)
}

