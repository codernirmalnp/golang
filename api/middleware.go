package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/codernirmalnp/golang/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(token token.Maker) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("Authorization Header is not provided")

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		field := strings.Fields(authorizationHeader)
		if len(field) < 2 {
			err := errors.New("Invalid Authorization Header format")

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		authorizationType := strings.ToLower(field[0])
		if authorizationType != authorizationTypeBearer {
			err := errors.New("unsupported Authorization type")

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return

		}
		accessToken := field[1]
		payload, err := token.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}

}
