package middleware

import (
	"context"
	"encoding/hex"
	"fmt"
	"math"
	"net/http"

	"strings"

	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/redis/go-redis/v9"
	cons "github.com/restuwahyu13/go-clean-architecture/shared/constants"
	helper "github.com/restuwahyu13/go-clean-architecture/shared/helpers"
	opt "github.com/restuwahyu13/go-clean-architecture/shared/output"
	"github.com/restuwahyu13/go-clean-architecture/shared/pkg"
)

func Auth(expired int, con *redis.Client) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			res := opt.Response{}

			jose := pkg.NewJose(ctx)
			crypto := helper.NewCrypto()

			headers := r.Header.Get("Authorization")
			if !strings.Contains(headers, "Bearer") {
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Authorization is required"

				helper.Api(w, r, res)
				return
			}

			token := strings.Split(headers, "Bearer ")[1]

			if len(strings.Split(token, ".")) != 3 {
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid token format"

				helper.Api(w, r, res)
				return
			}

			tokenMetadata, err := jwt.ParseRequest(r, jwt.WithHeaderKey("Authorization"), jwt.WithVerify(false))
			if err != nil {
				pkg.Logrus(cons.ERROR, err)
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, r, res)
				return
			}

			aud, ok := tokenMetadata.Audience()
			if !ok {
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, r, res)
				return
			}

			iss, ok := tokenMetadata.Issuer()
			if !ok {
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, r, res)
				return
			}

			sub, ok := tokenMetadata.Subject()
			if !ok {
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, r, res)
				return
			}

			jti, ok := tokenMetadata.JwtID()
			if !ok {
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, r, res)
				return
			}

			timestamp := ""
			if err := tokenMetadata.Get("timestamp", &timestamp); err != nil {
				pkg.Logrus(cons.ERROR, err)
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, r, res)
				return
			}

			suffix := int(math.Pow(float64(expired), float64(len(aud[0])+len(iss)+len(sub))))
			secretKey := fmt.Sprintf("%s:%s:%s:%s:%d", aud[0], iss, sub, timestamp, suffix)
			secretData := hex.EncodeToString([]byte(secretKey))

			key, err := crypto.AES256Decrypt(secretData, jti)
			if err != nil {
				pkg.Logrus(cons.ERROR, err)
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, r, res)
				return
			}

			rds, err := pkg.NewRedis(ctx, con)
			if err != nil {
				pkg.Logrus(cons.ERROR, err)
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, r, res)
				return
			}

			if _, err = jose.JwtVerify(key, token, rds); err != nil {
				pkg.Logrus(cons.ERROR, err)
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, r, res)
				return
			}

			sharingCtx := context.WithValue(r.Context(), "user_id", key)
			h.ServeHTTP(w, r.WithContext(sharingCtx))

			return
		})
	}
}
