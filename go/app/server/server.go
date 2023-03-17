package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/mizutanimeen/free-room/controllers"
	"github.com/mizutanimeen/free-room/models"
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

func (aServer *Server) TestGet(w http.ResponseWriter, r *http.Request) {
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

func (aServer *Server) TestPost(w http.ResponseWriter, r *http.Request) {
	tBody := make([]byte, r.ContentLength)
	r.Body.Read(tBody)

	tPing := &Ping{}
	if tError := json.Unmarshal(tBody, tPing); tError != nil {
		http.Error(w, tError.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (aServer *Server) SignUp(aResWriter http.ResponseWriter, aRequests *http.Request) {
	tSignUpUser, tStatus, tError := controllers.GetSignUpUser(aRequests)
	if tError != nil {
		http.Error(aResWriter, tError.Error(), tStatus)
	}

	if tStatus, tError := models.SignUpConditionCheck(tSignUpUser); tError != nil {
		http.Error(aResWriter, tError.Error(), tStatus)
	}

	//パスワードのハッシュ化

	//URLを作成、URLを登録するメールに送信、GETを検知、時間が３０分以内か確認、DBに保存
	//URLは /signup/email?xxx=yyyyyyyyyyみたいにするのかな？
	//https://qiita.com/isotai/items/f810493dd192e0597f3a
	//https://colo-ri.jp/develop/2012/07/web-how-to-mail-activation.html
	//https://www.google.com/search?q=golang+%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC%E7%99%BB%E9%8C%B2&oq=&aqs=chrome.0.35i39i362l8.277420j0j7&sourceid=chrome&ie=UTF-8
	// メール認証は一旦後回し

	//認証コードもあり

	//セッションベース OAuth OpenID
	log.Println(tSignUpUser)
	return
}
