package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Ping struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func tServer(w http.ResponseWriter, r *http.Request) {
	ping := &Ping{http.StatusOK, "ok"}

	res, err := json.Marshal(ping)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
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
	tRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{os.Getenv("REACT_URL")}, // Use this to allow specific origin hosts
		// AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	tRouter.Use(middleware.Logger)
	tRouter.Route("/free-room/api/v1", func(tRouter chi.Router) {
		tRouter.Get("/", tServer)
	})
	http.ListenAndServe(":8080", tRouter)
}
