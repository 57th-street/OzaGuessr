package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/57th-street/oza-gueser/apperrors"
	"github.com/57th-street/oza-gueser/utils"
	"github.com/dgrijalva/jwt-go"
)

type userIDKey struct{}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authorization := req.Header.Get("Authorization")

		authHeaders := strings.Split(authorization, " ")
		if len(authHeaders) != 2 {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid request header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		bearer, tokenString := authHeaders[0], authHeaders[1]
		if bearer != "Bearer" || tokenString == "" {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid request header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		token, err := utils.ValidateToken(tokenString)
		if err != nil {
			err = apperrors.Unauthorized.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// このコンテキストをセットする箇所は今後別パッケージに切り出す予定
			ctx := context.WithValue(req.Context(), userIDKey{}, claims["user_id"])
			req = req.WithContext(ctx)
			next.ServeHTTP(w, req)
		} else {
			err := apperrors.Unauthorized.Wrap(err, "invalid token")
			apperrors.ErrorHandler(w, req, err)
		}
	}
}
