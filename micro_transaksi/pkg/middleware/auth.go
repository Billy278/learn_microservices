package middleware

import (
	"fmt"
	"micro_transaksi/modules/model"
	"micro_transaksi/pkg/crypto"
	"micro_transaksi/pkg/responses"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type (
	Headerstring string
)

const (
	AccessClaim string       = "access_claim"
	Author      Headerstring = "Authorization"
	Bearer      string       = "Bearer "
)

func TestBearerOAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//auth header
		header := ctx.GetHeader(string(Author))
		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.FailRes{
				Code:    http.StatusUnauthorized,
				Message: "NOT AUTH",
				Error:   responses.Unauthorized,
			})
			return
		}
		//get token

		token := strings.Split(header, Bearer)
		fmt.Println(token)
		if len(token) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.FailRes{
				Code:    http.StatusUnauthorized,
				Message: "Invalid token",
				Error:   responses.Unauthorized,
			})
			return
		}
		var claim model.AccessClaim
		fmt.Println(claim)
		err := crypto.ParseJWT(token[1], &claim)
		fmt.Println(claim)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.FailRes{
				Code:    http.StatusUnauthorized,
				Message: "failed token",
				Error:   responses.Unauthorized,
			})
			return
		}
		ctx.Set(AccessClaim, claim)
		ctx.Next()

	}
}

func BearerOAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//auth header
		header := ctx.GetHeader(string(Author))
		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.FailRes{
				Code:    http.StatusUnauthorized,
				Message: "NOT AUTH",
				Error:   responses.Unauthorized,
			})
			return
		}
		//get token

		token := strings.Split(header, Bearer)
		fmt.Println(token)
		if len(token) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.FailRes{
				Code:    http.StatusUnauthorized,
				Message: "Invalid token",
				Error:   responses.Unauthorized,
			})
			return
		}
		var claim model.AccessClaim
		fmt.Println(claim)
		err := crypto.ParseJWT(token[1], &claim)
		fmt.Println(claim)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.FailRes{
				Code:    http.StatusUnauthorized,
				Message: "failed token",
				Error:   responses.Unauthorized,
			})
			return
		}
		ctx.Set(AccessClaim, claim)
		ctx.Next()
	}
}
