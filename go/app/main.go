package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/mizutanimeen/free-room/db"
	logFile "github.com/mizutanimeen/free-room/logs"
	"github.com/mizutanimeen/free-room/server"
)

func main() {
	//ログ出力先設定
	tLogFile, tError := logFile.Settings("./logs/data/go.log")
	if tError != nil {
		log.Fatalf("%s\n", tError.Error())
	}
	defer func(aLogFile *os.File) {
		if tError := aLogFile.Close(); tError != nil {
			log.Fatalf("%s\n", tError.Error())
		}
	}(tLogFile)

	log.Println("--------------Log Start--------------")

	//環境変数読み込み
	if tError := godotenv.Load(".env"); tError != nil {
		log.Fatalf("%s\n", tError.Error())
	}

	//DB接続
	tDb, tError := db.Open()
	if tError != nil {
		log.Fatalf("%s\n", tError.Error())
	}
	defer func(aDb *sql.DB) {
		if tError := aDb.Close(); tError != nil {
			log.Fatalf("%s\n", tError.Error())
		}
	}(tDb)

	tServer := server.New(tDb)

	//ルーティング
	tRouter := chi.NewRouter()
	tRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("REACT_URL")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	tRouter.Use(middleware.Logger)
	tRouter.Route("/free-room/api/v1", func(tRouter chi.Router) {
		tRouter.Get("/", tServer.TestGet)
		tRouter.Post("/", tServer.TestPost)
		tRouter.Post("/signup", tServer.SignUp)
	})
	http.ListenAndServe(":8080", tRouter)
}
