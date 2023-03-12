package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func tServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	if tError := godotenv.Load(".env"); tError != nil {
		log.Fatalf("%s\n", tError.Error())
	}

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

	tRouter := chi.NewRouter()
	tRouter.Use(middleware.Logger)
	tRouter.Route("/free-room/api/v1", func(tRouter chi.Router) {
		tRouter.Get("/", tServer)
	})
	http.ListenAndServe(":8080", tRouter)
}
