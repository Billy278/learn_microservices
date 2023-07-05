package repository

import (
	"context"
	"micro_user/modules/model"
)

type UserRepo interface {
	ShowAllUser(ctx context.Context) (UsersRes []model.User, err error)
	CreateUser(ctx context.Context, UserIn model.User) (UserRes model.User, err error)
	UpdateUser(ctx context.Context, UserIn model.User) (UserRes model.User, err error)
	FindByIdUser(ctx context.Context, UserId int) (UserRes model.User, err error)
	DeleteUser(ctx context.Context, UserId int) (UserRes model.User, err error)
	FindByUsername(ctx context.Context, Username string) (UserRes model.User, err error)
}
