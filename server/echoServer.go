package server

import (
	"auth/config"
	"auth/database"
	help_rst "auth/src/help/adapter/controller/rest"
	"auth/src/org/adapter/controller/rest"

	help_psql "auth/src/help/adapter/gateway/repo/psql"
	hel_usecase "auth/src/help/usecase"
	"auth/src/org/adapter/gateway/repo/psql"
	"auth/src/org/adapter/gateway/tin_checker/etrade"
	"auth/src/org/usecase"
	"fmt"
	"log"
	"net/http"
	"os"
	// "auth/database"
	// "auth/routers"
)

type AppServer struct {
	App *http.ServeMux
	Cfg *config.Config
}

func NewEchoServer(cfg *config.Config) Server {
	return &AppServer{
		App: http.NewServeMux(),
		Cfg: cfg,
	}
}

func (s *AppServer) Start() {
	serverUrl := fmt.Sprintf(":%d", s.Cfg.App.Port)

	db := database.NewPostgresDatabase(s.Cfg)
	logger := log.New(os.Stdout, "example ", log.LstdFlags)
	repo, err := psql.New(logger, db.GetDb())
	if err != nil {
		log.Fatal(err)
	}
	tin := etrade.New(logger)
	interactor := usecase.New(logger, repo, tin)
	// serveMux := http.NewServeMux()

	repo2, err := help_psql.New(logger, db.GetDb())

	if err != nil {
		log.Fatal(err)
	}

	help_interactor := hel_usecase.New(logger, repo2)

	rest.New(logger, interactor, s.App)

	help_rst.New(logger, help_interactor, s.App)

	corsHandler := enableCORS(s.App)

	server := &http.Server{
		Addr:    serverUrl,
		Handler: corsHandler,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if req.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, req)
	})
}

// func (s *AppServer) initializeRoutes() {
// todosGroup := s.App.Group("v1/todos")
// routers.SetupTodosRoutes(todosGroup, s.Cfg)

// userHandler := handlers.NewUserHandlerImp(db)
// s.App.POST("v1/system-key", userHandler.AddSystemKeyhandler)
// s.App.GET("v1/system-key", userHandler.GetSystemKeyHandler)

// }
