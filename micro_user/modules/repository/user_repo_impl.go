package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"micro_user/modules/model"
)

type UserRepoImpl struct {
	DB *sql.DB
}

func NewUserRepoImpl(db *sql.DB) UserRepo {
	return &UserRepoImpl{
		DB: db,
	}
}
func (user_repo *UserRepoImpl) ShowAllUser(ctx context.Context) (UsersRes []model.User, err error) {
	logCtx := fmt.Sprintf("%T, Repo ShowAllUser", user_repo)
	log.Println(logCtx)
	sql := "SELECT id,email,name,username,password,no_hp,alamat,jenkel FROM tb_user"
	row, err := user_repo.DB.QueryContext(ctx, sql)
	if err != nil {
		return
	}
	defer row.Close()
	user := model.User{}
	for row.Next() {
		err = row.Scan(&user.Id, &user.Email, &user.Name, &user.Username, &user.Password, &user.No_hp, &user.Alamat, user.Jenkel)
		if err != nil {
			return
		}
		UsersRes = append(UsersRes, user)
	}
	return
}
func (user_repo *UserRepoImpl) CreateUser(ctx context.Context, UserIn model.User) (UserRes model.User, err error) {
	logCtx := fmt.Sprintf("%T, Repo CreateUser", user_repo)
	log.Println(logCtx)
	fmt.Println(UserIn)
	sql := "INSERT INTO tb_user(email,name,username,password,no_hp,alamat,jenkel) VALUES ($1,$2,$3,$4,$5,$6,$7)"
	_, err = user_repo.DB.ExecContext(ctx, sql, UserIn.Email, UserIn.Name, UserIn.Username, UserIn.Password, UserIn.No_hp, UserIn.Alamat, UserIn.Jenkel)
	if err != nil {
		return
	}
	UserRes.Username = UserIn.Username
	return
}
func (user_repo *UserRepoImpl) UpdateUser(ctx context.Context, UserIn model.User) (UserRes model.User, err error) {
	logCtx := fmt.Sprintf("%T, Repo UpdateUser", user_repo)
	log.Println(logCtx)
	sql := "UPDATE tb_user SET email=$1,name=$2,username=$3,password=$4,no_hp=$5,alamat=$6,jenkel=$7 WHERE id=$8"
	_, err = user_repo.DB.ExecContext(ctx, sql, UserIn.Email, UserIn.Name, UserIn.Username, UserIn.Password, UserIn.No_hp, UserIn.Alamat, UserIn.Jenkel, UserIn.Id)
	if err != nil {
		return
	}
	return

}
func (user_repo *UserRepoImpl) FindByIdUser(ctx context.Context, UserId int) (UserRes model.User, err error) {
	logCtx := fmt.Sprintf("%T, Repo FindByIdUser", user_repo)
	log.Println(logCtx)
	sql := "SELECT id,email,name,username,password,no_hp,alamat,jenkel FROM tb_user WHERE id=$1"
	row, err := user_repo.DB.QueryContext(ctx, sql, UserId)
	if err != nil {
		return
	}
	if row.Next() {
		err = row.Scan(&UserRes.Id, &UserRes.Email, &UserRes.Name, &UserRes.Username, &UserRes.Password, &UserRes.No_hp, &UserRes.Alamat, &UserRes.Jenkel)
		if err != nil {
			return
		}
		return
	} else {
		err = errors.New("USER NOT FOUND")
		return
	}

}
func (user_repo *UserRepoImpl) DeleteUser(ctx context.Context, UserId int) (UserRes model.User, err error) {

	logCtx := fmt.Sprintf("%T, Repo DeleteUser", user_repo)
	log.Println(logCtx)
	sql := "DELETE FROM tb_user WHERE id=$1"
	_, err = user_repo.DB.ExecContext(ctx, sql, UserId)
	if err != nil {
		return
	}
	return
}

func (user_repo *UserRepoImpl) FindByUsername(ctx context.Context, Username string) (UserRes model.User, err error) {
	logCtx := fmt.Sprintf("%T, Repo FindByIdUser", user_repo)
	log.Println(logCtx)
	sql := "SELECT id,email,name,username,password,no_hp,alamat,jenkel FROM tb_user WHERE username=$1"
	row, err := user_repo.DB.QueryContext(ctx, sql, Username)
	if err != nil {
		return
	}
	if row.Next() {
		err = row.Scan(&UserRes.Id, &UserRes.Email, &UserRes.Name, &UserRes.Username, &UserRes.Password, &UserRes.No_hp, &UserRes.Alamat, &UserRes.Jenkel)
		if err != nil {
			return
		}
		return
	} else {
		err = errors.New("USER NOT FOUND")
		return
	}

}
