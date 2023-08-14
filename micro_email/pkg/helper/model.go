package helper

import (
	"micro_email/modules/model"
	"micro_email/modules/proto"
)

func ProtoEmailToStruct(emailIn *proto.Email) model.Email {
	return model.Email{
		Id:             emailIn.Id,
		Sender:         emailIn.Sender,
		Name_Receiver:  emailIn.NameReceiver,
		Name_product:   emailIn.NameProduct,
		Harga:          float64(emailIn.Harga),
		Qty:            emailIn.Qty,
		Total:          float64(emailIn.Total),
		Email_Receiver: emailIn.EmailReceiver,
	}

}

func StructToprotoEmail(emailIn model.Email) *proto.Email {
	return &proto.Email{
		Id:            emailIn.Id,
		Sender:        emailIn.Sender,
		NameReceiver:  emailIn.Name_Receiver,
		NameProduct:   emailIn.Name_product,
		Harga:         float32(emailIn.Harga),
		Qty:           emailIn.Qty,
		Total:         float32(emailIn.Total),
		EmailReceiver: emailIn.Email_Receiver,
	}

}

func StructToprotoEmails(emailIn []model.Email) *proto.EmailList {
	protoEmails := &proto.EmailList{}
	for _, v := range emailIn {
		protoEmails.List = append(protoEmails.List, StructToprotoEmail(v))
	}
	return protoEmails

}
