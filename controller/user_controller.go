package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/illionillion/go-sns-api-learning/models"
	"github.com/illionillion/go-sns-api-learning/usecase"
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	GetUser(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

// GetUser godoc
// @Summary      ユーザー情報取得
// @Description  ユーザーIDを指定してユーザー情報を取得する
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        userId path string true "ユーザーID"
// @Success      200 {object} models.UserResponse
// @Failure      400 {object} string
// @Failure      500 {object} string
// @Router       /users/{userId} [get]
func (uc *userController) GetUser(c echo.Context) error {
	userId := c.Param("userId")
	user, err := uc.uu.GetUser(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

// SignUp godoc
// @Summary      新規ユーザー登録
// @Description  ユーザーを新規作成して返す
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user body models.UserSignupRequest true "ユーザー情報"
// @Success      201 {object} models.UserResponse
// @Failure      400 {object} string
// @Failure      500 {object} string
// @Router       /signup [post]
func (uc *userController) SignUp(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.Signup(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

// LogIn godoc
// @Summary      ログイン
// @Description  ユーザー認証して JWT を Cookie にセットする
// @Tags         auth
// @Accept       json
// @Produce      plain
// @Param        user body models.UserLoginRequest true "ログイン情報"
// @Success      200
// @Failure      400 {object} string
// @Failure      500 {object} string
// @Router       /login [post]
func (uc *userController) LogIn(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// クッキー
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

// LogOut godoc
// @Summary      ログアウト
// @Description  Cookie を削除することでログアウト
// @Tags         auth
// @Produce      plain
// @Success      200
// @Router       /logout [post]
func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
