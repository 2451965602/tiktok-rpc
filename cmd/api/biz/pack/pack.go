package pack

import (
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"tiktokrpc/cmd/api/biz/model/model"
	"tiktokrpc/cmd/api/pkg/errmsg"
)

func SendResponse(c *app.RequestContext, data interface{}) {
	c.JSON(consts.StatusOK, data)
}

func SendFailResponse(c *app.RequestContext, data *model.BaseResp) {
	c.JSON(consts.StatusBadRequest, utils.H{
		"base": data,
	})
}

func BuildBaseResp(err errmsg.ErrorMessage) *model.BaseResp {
	return &model.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}

func BuildFailResponse(c *app.RequestContext, err error) {
	if err == nil {
		SendFailResponse(c, BuildBaseResp(errmsg.NoError))

		return
	}

	e := errmsg.ErrorMessage{}
	if errors.As(err, &e) {
		SendFailResponse(c, BuildBaseResp(e))

		return
	}

	e = errmsg.ServiceError.WithMessage(err.Error())
	SendFailResponse(c, BuildBaseResp(e))
}
