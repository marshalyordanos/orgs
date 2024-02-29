package routers

import (
	"auth/config"

	"github.com/labstack/echo/v4"
)

func SetupOrgsRoutes(group *echo.Group, cfg *config.Config) {
	// db := database.NewPostgresDatabase(cfg)
	// logger := log.New(os.Stdout, "example ", log.LstdFlags)
	// repo, err := psql.New(logger, db.GetDb())
	// if err != nil {
	// 	fmt.Errorf("Eror %v", err.Error())
	// }
	// tin := etrade.New(logger)
	// interactor := usecase.New(logger, repo, tin)
	// serveMux := http.NewServeMux()

	// controller := rest.New(logger, interactor, serveMux)

	// group.POST("/", todoHandler.TodoHandlerFunc)
	// group.GET("/check-tin", todoHandler.GetTodoHandlerFunc)

}
