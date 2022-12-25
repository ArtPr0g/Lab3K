package app

import (
	"awesomeProject/internal/app/ds"
	"awesomeProject/internal/app/role"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

type loginReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type loginResp struct {
	ExpiresIn   time.Duration `json:"expires_in"`
	AccessToken string        `json:"access_token"`
	TokenType   string        `json:"token_type"`
}

func (a *Application) Login(gCtx *gin.Context) {
	cfg := a.config
	req := &loginReq{}

	err := json.NewDecoder(gCtx.Request.Body).Decode(req)
	if err != nil {
		_ = gCtx.AbortWithError(http.StatusBadRequest, err)

		return
	}

	user, err := a.repo.GetUserByLogin(req.Login)
	if err != nil {
		_ = gCtx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	if req.Login == user.Name && user.Pass == generateHashString(req.Password) {
		// значит проверка пройдена
		// генерируем ему jwt
		token := jwt.NewWithClaims(cfg.JWT.SigningMethod, &ds.JWTClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(cfg.JWT.ExpiresIn).Unix(),
				IssuedAt:  time.Now().Unix(),
				Issuer:    "bitop-admin",
			},
			UserUUID: user.UUID,
			Role:     user.Role,
		})

		if token == nil {
			_ = gCtx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("token is nil"))

			return
		}

		strToken, err := token.SignedString([]byte(cfg.JWT.Token))
		if err != nil {
			_ = gCtx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("cant create str token"))

			return
		}

		gCtx.SetCookie("access_token", strToken, int(cfg.JWT.ExpiresIn), "/", "localhost", false, false)

		gCtx.JSON(http.StatusOK, loginResp{
			ExpiresIn:   cfg.JWT.ExpiresIn,
			AccessToken: strToken,
			TokenType:   "Bearer",
		})
		return
	}

	gCtx.AbortWithStatus(http.StatusForbidden) // отдаем 403 ответ в знак того что доступ запрещен
}

type pingResp struct {
	Status bool `json:"status"`
}

func (a *Application) Ping(gCtx *gin.Context) {
	gCtx.JSON(http.StatusOK, &pingResp{
		Status: true,
	})
}

type registerReq struct {
	Name string `json:"name"` // лучше назвать то же самое что login
	Pass string `json:"pass"`
}

type registerResp struct {
	Ok bool `json:"ok"`
}

func (a *Application) Register(gCtx *gin.Context) {
	req := &registerReq{}

	err := json.NewDecoder(gCtx.Request.Body).Decode(req)
	if err != nil {
		_ = gCtx.AbortWithError(http.StatusBadRequest, err)

		return
	}

	if req.Pass == "" {
		_ = gCtx.AbortWithError(http.StatusBadRequest, fmt.Errorf("pass is empty"))

		return
	}

	if req.Name == "" {
		_ = gCtx.AbortWithError(http.StatusBadRequest, fmt.Errorf("name is empty"))

		return
	}

	err = a.repo.Register(&ds.User{
		Role: role.Buyer,
		Name: req.Name,
		Pass: generateHashString(req.Pass), // пароли делаем в хешированном виде и далее будем сравнивать хеши, чтобы их не угнали с базой вместе
	})

	if err != nil {
		_ = gCtx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	gCtx.JSON(http.StatusOK, &registerResp{
		Ok: true,
	})
}

func generateHashString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

type logoutResp struct {
	Status bool `json:"status"`
}

func (a *Application) Logout(gCtx *gin.Context) {
	jwtStr := gCtx.GetHeader("Authorization")
	if !strings.HasPrefix(jwtStr, jwtPrefix) { // если нет префикса то нас дурят!
		gCtx.AbortWithStatus(http.StatusBadRequest) // отдаем что нет доступа

		return // завершаем обработку
	}

	// отрезаем префикс
	jwtStr = jwtStr[len(jwtPrefix):]

	_, err := jwt.ParseWithClaims(jwtStr, &ds.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.config.JWT.Token), nil
	})
	if err != nil {
		_ = gCtx.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)

		return
	}

	// сохраняем в блеклист редиса
	err = a.redis.WriteJWTToBlacklist(gCtx.Request.Context(), jwtStr, a.config.JWT.ExpiresIn)
	if err != nil {
		_ = gCtx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	gCtx.SetCookie("access_token", "", -1, "/", "localhost", false, true)

	gCtx.JSON(http.StatusOK, &logoutResp{
		Status: true,
	})
}
