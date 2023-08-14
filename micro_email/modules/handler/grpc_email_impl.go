package handler

import (
	"context"
	"micro_email/modules/proto"
	"micro_email/modules/services"
	"micro_email/pkg/helper"
	"net/http"
)

type HandlerEmailImpl struct {
	EmailServ services.ServEmail
	proto.UnimplementedEmailSrvServer
}

func NewHandlerEmailImpl(emailserv services.ServEmail) *HandlerEmailImpl {
	return &HandlerEmailImpl{
		EmailServ: emailserv,
	}
}

func (h *HandlerEmailImpl) SendEmail(ctx context.Context, reqIn *proto.Email) (*proto.Response, error) {
	res, err := h.EmailServ.ServCreate(ctx, helper.ProtoEmailToStruct(reqIn))
	if err != nil {
		return &proto.Response{
			Code:  http.StatusInternalServerError,
			Error: err.Error(),
		}, err
	}
	return &proto.Response{
		Code: http.StatusOK,
		Data: &proto.EmailList{
			List: []*proto.Email{
				helper.StructToprotoEmail(res),
			},
		},
	}, nil
}
