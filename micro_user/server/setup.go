package server

import (
	"micro_user/db"
	"micro_user/modules/controllers"
	"micro_user/modules/repository"
	"micro_user/modules/services"

	"github.com/go-playground/validator/v10"
)

type Controllers struct {
	UserCtrl controllers.UserCtrl
}

func NewSetup() Controllers {
	db := db.NewDBPostges()
	validate := validator.New()
	userRepo := repository.NewUserRepoImpl(db)
	userServ := services.NewUserServImpl(userRepo)
	userCtrl := controllers.NewUserCtrlImpl(userServ, *validate)

	return Controllers{
		UserCtrl: userCtrl,
	}

}
