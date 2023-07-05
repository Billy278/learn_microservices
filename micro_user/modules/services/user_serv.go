package services

import (
	"context"
	"micro_user/modules/model"
)

type UserServ interface {
	ServShowAllUser(ctx context.Context) (UsersRes []model.User, err error)
	ServCreateUser(ctx context.Context, UserIn model.User) (UserRes model.User, err error)
	ServUpdateUser(ctx context.Context, UserIn model.User) (UserRes model.User, err error)
	ServFindByIdUser(ctx context.Context, UserId int) (UserRes model.User, err error)
	ServDeleteUser(ctx context.Context, UserId int) (UserRes model.User, err error)
	ServLoginUser(ctx context.Context, Username string, Password string) (tokens model.Tokens, err error)
}
