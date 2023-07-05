package server

import (
	"micro_email/db"
	"micro_email/modules/controllers"
	"micro_email/modules/repository"
	"micro_email/modules/services"

	"github.com/go-playground/validator/v10"
)

type Controllers struct {
	CtrlEmail controllers.CtrlEmail
}

func NewSetup() Controllers {
	db := db.NewDBPostges()
	validate := validator.New()
	repoEmail := repository.NewRepoEmailImpl(db)
	servEmail := services.NewServEmailImpl(repoEmail)
	ctrlEmail := controllers.NewCtrlEmailImpl(servEmail, validate)
	return Controllers{
		CtrlEmail: ctrlEmail,
	}
}
