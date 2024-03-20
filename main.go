package main

import (
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"os"

	"gocrudapp/model"
	"gocrudapp/repository/sq3"
	"gocrudapp/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	slog.SetDefault(logger)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("env successfully loaded")
}

func main() {
	db := sq3dbconn(os.Getenv("DB_NAME"), model.UserSchema)
	defer db.Close()

	userRepo := sq3.UserRepository{DB: db}
	userService := usecase.UserService{UserRepo: userRepo}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is running at :4444"))
		w.WriteHeader(http.StatusOK)
	})

	r.Post("/v1/user", usecase.CreateUserHandler(&userService))
	r.Get("/v1/user", usecase.FetchAllUserHandler(&userService))
	r.Get("/v1/user/{uid}", usecase.FetchUserByIDHandler(&userService))
	r.Put("/v1/user/{uid}", usecase.UpdateUserByIDHandler(&userService))
	r.Delete("/v1/user/{uid}", usecase.DeleteUserByIDHandler(&userService))
	r.Delete("/v1/user", usecase.DeleteAllUserHandler(&userService))

	slog.Info("starting server at :4444")
	http.ListenAndServe(":4444", r)
}

func sq3dbconn(dbname, schema string) *sql.DB {
	db, err := sql.Open("sqlite3", dbname+".sqlite")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("db connected")

	if _, err = db.Exec(schema); err != nil {
		log.Fatal(err)
	}

	slog.Info("table created if not exist")

	return db
}
