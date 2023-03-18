package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"

	"github.com/mizutanimeen/free-room/controllers"
	"github.com/mizutanimeen/free-room/db"
)

type Server struct {
	Db *sql.DB
}

func New(aDb *sql.DB) *Server {
	return &Server{
		Db: aDb,
	}
}

type Ping struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func TestGet(w http.ResponseWriter, r *http.Request) {
	log.Println("TestGet")
	tPing := &Ping{http.StatusOK, "ok"}

	tPingByte, tError := json.Marshal(tPing)
	if tError != nil {
		http.Error(w, tError.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if _, tError := w.Write(tPingByte); tError != nil {
		http.Error(w, tError.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func TestPost(w http.ResponseWriter, r *http.Request) {
	tBody := make([]byte, r.ContentLength)
	r.Body.Read(tBody)

	tPing := &Ping{}
	if tError := json.Unmarshal(tBody, tPing); tError != nil {
		http.Error(w, tError.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (aServer *Server) GetUsers(aResponseWriter http.ResponseWriter, aRequest *http.Request) {
	tUsers, tStatus, tError := db.GetAllUsers(aServer.Db)
	if tError != nil {
		http.Error(aResponseWriter, tError.Error(), tStatus)
	}

	tUsersByte, tError := json.Marshal(tUsers)
	if tError != nil {
		http.Error(aResponseWriter, tError.Error(), http.StatusInternalServerError)
		return
	}

	aResponseWriter.Header().Set("Content-Type", "application/json")
	aResponseWriter.WriteHeader(200)
	if _, tError := aResponseWriter.Write(tUsersByte); tError != nil {
		http.Error(aResponseWriter, tError.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (aServer *Server) GetUserByID(aResponseWriter http.ResponseWriter, aRequest *http.Request) {
	tID := chi.URLParam(aRequest, "ID")

	tUser, tStatus, tError := db.GetUserByID(tID, aServer.Db)
	if tError != nil {
		http.Error(aResponseWriter, tError.Error(), tStatus)
	}

	tUserByte, tError := json.Marshal(tUser)
	if tError != nil {
		http.Error(aResponseWriter, tError.Error(), http.StatusInternalServerError)
		return
	}

	aResponseWriter.Header().Set("Content-Type", "application/json")
	aResponseWriter.WriteHeader(200)
	if _, tError := aResponseWriter.Write(tUserByte); tError != nil {
		http.Error(aResponseWriter, tError.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (aServer *Server) SignUp(aResponseWriter http.ResponseWriter, aRequest *http.Request) {
	tSignUpUser, tStatus, tError := controllers.GetSignUpUser(aRequest)
	if tError != nil {
		http.Error(aResponseWriter, tError.Error(), tStatus)
		return
	}

	if tStatus, tError := tSignUpUser.ConditionCheck(); tError != nil {
		http.Error(aResponseWriter, tError.Error(), tStatus)
		return
	}

	//パスワードのハッシュ化

	if tStatus, tError := db.SaveSignUpUser(tSignUpUser, aServer.Db); tError != nil { //データが重複したときの例外処理、react側が欲しいなら渡すように
		log.Println(tError)
		http.Error(aResponseWriter, tError.Error(), tStatus)
		return
	}

	//URLを作成、URLを登録するメールに送信、GETを検知、時間が３０分以内か確認、DBに保存
	//URLは /signup/email?xxx=yyyyyyyyyyみたいにするのかな？
	//https://qiita.com/isotai/items/f810493dd192e0597f3a
	//https://colo-ri.jp/develop/2012/07/web-how-to-mail-activation.html
	//https://www.google.com/search?q=golang+%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC%E7%99%BB%E9%8C%B2&oq=&aqs=chrome.0.35i39i362l8.277420j0j7&sourceid=chrome&ie=UTF-8
	// メール認証は一旦後回し

	//認証コードもあり

	//セッションベース OAuth OpenID
	aResponseWriter.WriteHeader(http.StatusCreated)
	return
}
