package middleware

import (
	"fmt"
	"micro_product/modules/model"
	"micro_product/pkg/crypto"
	"micro_product/pkg/responses"
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
	Key         string       = "KEY"
	//"Authorization"
	//lama
	//"Authorization"
)

func BearerOAuthZ() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//auth header
		header := ctx.GetHeader(string(Author))
		k := ctx.GetHeader(Key)

		if k != "" {
			if k == crypto.SharedKey {
				ctx.Next()
			} else {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.FailRes{
					Code:    http.StatusUnauthorized,
					Message: "Invalid KEY",
					Error:   responses.Unauthorized,
				})
				return
			}
		} else if header != "" {
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
			err := crypto.ParseJWT(token[1], &claim)
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
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.FailRes{
				Code:    http.StatusUnauthorized,
				Message: "NOT AUTH",
				Error:   responses.Unauthorized,
			})
			return
		}
	}
}
