package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	logFile "github.com/mizutanimeen/free-room/logs"
	"github.com/mizutanimeen/free-room/server"
)

func main() {
	//ログ出力先設定
	tLogFile, tError := logFile.Settings("./logs/data/go.log")
	defer func(aLogFile *os.File) {
		if tError := aLogFile.Close(); tError != nil {
			log.Fatalf("%s\n", tError.Error())
		}
	}(tLogFile)
	if tError != nil {
		log.Fatalf("%s\n", tError.Error())
	}

	log.Println("--------------Log Start--------------")

	//環境変数読み込み
	if tError := godotenv.Load(".env"); tError != nil {
		log.Fatalf("%s\n", tError.Error())
	}

	//DB接続
	tDbPath := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_IP"), os.Getenv("MYSQL_DATABASE"))
	tDb, tError := sql.Open("mysql", tDbPath)
	defer tDb.Close()
	if tError != nil {
		log.Fatalf("%s\n", tError.Error())
	}
	if tError = tDb.Ping(); tError != nil {
		log.Fatalf("データベース接続失敗: %s\n", tError.Error())
	} else {
		log.Printf("データベース接続成功\n")
	}

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
