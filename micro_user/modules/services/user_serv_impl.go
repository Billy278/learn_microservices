package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"micro_user/modules/model"
	"micro_user/modules/repository"
	"micro_user/pkg/crypto"
	"strconv"
	"sync"
	"time"
)

type UserServImpl struct {
	UserRepo repository.UserRepo
}

func NewUserServImpl(userrepo repository.UserRepo) UserServ {
	return &UserServImpl{
		UserRepo: userrepo,
	}
}
func (user_ser *UserServImpl) ServShowAllUser(ctx context.Context) (UsersRes []model.User, err error) {
	logCtx := fmt.Sprintf("%T, Services  ShowAllUser", user_ser)
	log.Println(logCtx)
	UsersRes, err = user_ser.UserRepo.ShowAllUser(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	return

}
func (user_ser *UserServImpl) ServCreateUser(ctx context.Context, UserIn model.User) (UserRes model.User, err error) {
	logCtx := fmt.Sprintf("%T, Services  ServCreateUser", user_ser)
	log.Println(logCtx)
	hashpassword, err := crypto.GenerateHash(UserIn.Password)
	if err != nil {
		return
	}
	UserIn.Password = hashpassword
	UserRes, err = user_ser.UserRepo.CreateUser(ctx, UserIn)
	if err != nil {
		log.Println(err)
		return
	}
	return

}
func (user_ser *UserServImpl) ServUpdateUser(ctx context.Context, UserIn model.User) (UserRes model.User, err error) {
	logCtx := fmt.Sprintf("%T, Services  ServUpdateUser", user_ser)
	log.Println(logCtx)
	_, err = user_ser.UserRepo.FindByIdUser(ctx, UserIn.Id)
	if err != nil {
		return
	}
	UserRes, err = user_ser.UserRepo.UpdateUser(ctx, UserIn)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
func (user_ser *UserServImpl) ServFindByIdUser(ctx context.Context, UserId int) (UserRes model.User, err error) {
	logCtx := fmt.Sprintf("%T, Services  ServFindByIdUser", user_ser)
	log.Println(logCtx)
	UserRes, err = user_ser.UserRepo.FindByIdUser(ctx, UserId)
	if err != nil {
		return
	}
	return
}
func (user_ser *UserServImpl) ServDeleteUser(ctx context.Context, UserId int) (UserRes model.User, err error) {
	logCtx := fmt.Sprintf("%T, Services  ServDeleteUser", user_ser)
	log.Println(logCtx)
	_, err = user_ser.UserRepo.FindByIdUser(ctx, UserId)
	if err != nil {
		return
	}
	UserRes, err = user_ser.UserRepo.DeleteUser(ctx, UserId)
	return
}

func (user_ser *UserServImpl) ServLoginUser(ctx context.Context, Username string, Password string) (tokens model.Tokens, err error) {
	logCtx := fmt.Sprintf("%T, Services  ServLoginUser", user_ser)
	log.Println(logCtx)
	userRes, err := user_ser.UserRepo.FindByUsername(ctx, Username)
	if err != nil {
		return
	}
	err = crypto.ComparateHash(userRes.Password, Password)
	if err != nil {
		err = errors.New("PASSWORD SALAH")
		return
	}
	//create token
	jti := "jti"
	usr_id := strconv.Itoa(userRes.Id)
	idtoken, accesstoken, refreshtoken, err := user_ser.generateAllTokenConcurency(ctx, usr_id, Username, jti, userRes.Email)
	if err != nil {
		return
	}

	return model.Tokens{
		IDToken:      idtoken,
		AccessToken:  accesstoken,
		RefreshToken: refreshtoken,
	}, err
}

func (user_ser *UserServImpl) generateAllTokenConcurency(ctx context.Context, userId string, username, jti, email string) (idtoken, accesstoken, refreshtoken string, err error) {
	logCtx := fmt.Sprintf("%T, generateAllTokenConcurency", user_ser)
	log.Println(logCtx)
	timeNow := time.Now()
	defaultClaim := model.DefaultClaim{
		Expired:   int(timeNow.Add(24 * time.Hour).Unix()),
		NotBefore: int(timeNow.Unix()),
		IssuedAt:  int(timeNow.Unix()),
		Issuer:    userId,
		Audience:  "micro_user",
		JTI:       jti,
		Type:      model.ID_TOKEN,
	}
	var wg sync.WaitGroup
	wg.Add(3)
	go func(defaultclaim_ model.DefaultClaim) {
		defer wg.Done()
		//generate id token
		idTokenClaim := struct {
			model.DefaultClaim
			model.IDclaim
		}{
			defaultclaim_,
			model.IDclaim{
				IdUser:   userId,
				Username: username,
				Email:    email,
			},
		}
		idtoken, err = crypto.SignJWT(idTokenClaim)
		if err != nil {
			log.Println("Error create Id token")
			return
		}

	}(defaultClaim)
	go func(defaultClaim_ model.DefaultClaim) {
		defer wg.Done()
		//generate access token
		defaultClaim_.Expired = int(timeNow.Add(2 * time.Hour).Unix())
		defaultClaim_.Type = model.ACCESS_TOKEN
		accesstokenClaim := struct {
			model.DefaultClaim
			model.AccessClaim
		}{
			defaultClaim_,
			model.AccessClaim{
				UserId:   userId,
				Username: username,
				Email:    email,
			},
		}
		accesstoken, err = crypto.SignJWT(accesstokenClaim)
		if err != nil {
			log.Println("Error create access token")
			return
		}
	}(defaultClaim)

	go func(defaultClaim_ model.DefaultClaim) {
		defer wg.Done()
		defaultClaim_.Expired = int(timeNow.Add(time.Hour).Unix())
		defaultClaim_.Type = model.REFRESH_TOKEN
		refreshtokenClaim := struct {
			model.DefaultClaim
		}{
			defaultClaim_,
		}
		refreshtoken, err = crypto.SignJWT(refreshtokenClaim)
		if err != nil {
			log.Printf("[ERROR] creating refresh token  :%v\n", err)
			return
		}
	}(defaultClaim)
	wg.Wait()
	return
}
