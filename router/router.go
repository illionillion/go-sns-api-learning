package router

import (
	"github.com/illionillion/go-sns-api-learning/controller"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
    _ "github.com/illionillion/go-sns-api-learning/docs" // 生成されたSwagger docsのパッケージ
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	// Swagger UIのルートを登録
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	// t := e.Group("/tasks")
	// t.Use(echojwt.WithConfig(echojwt.Config{
	// 	SigningKey: []byte(os.Getenv("SECRET")),
	// 	TokenLookup: "cookie:token",
	// }))
	// t.GET("", tc.GetAllTasks)
	// t.GET("/:taskId", tc.GetTaskById)
	// t.POST("", tc.CreateTask)
	// t.PUT("/:taskId", tc.UpdateTask)
	// t.DELETE("/:taskId", tc.DeleteTask)
	return e
}
