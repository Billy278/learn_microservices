package crypto

import (
	"errors"

	"github.com/kataras/jwt"
)

const SharedKey = ("sercretkey5453498396898Done")

func SignJWT(claim any) (token string, err error) {
	//sign JWT
	tkn, err := jwt.Sign(jwt.HS256, []byte(SharedKey), claim)
	if err != nil {
		err = errors.New("ERROR SIGN CLAIM")
		return
	}
	token = string(tkn)
	return
}
func ParseJWT(token string, clams any) (err error) {
	//vrifikasi token
	verifyToken, err := jwt.Verify(jwt.HS256, []byte(SharedKey), []byte(token))
	if err != nil {
		err = errors.New("ERROR PARSE JWT")
		return
	}
	err = verifyToken.Claims(&clams)
	return
}
