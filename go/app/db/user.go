package db

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mizutanimeen/free-room/models"
)

// コココードひどいなーどうにかならんかな
func SaveSignUpUser(aSignUpUser *models.SignUpUser, aDb *sql.DB) (int, error) {
	tColumns := fmt.Sprintf("(%s,%s,%s,%s)",
		os.Getenv("MYSQL_USERS_TABLES_EMAIL"), os.Getenv("MYSQL_USERS_TABLES_NAME"),
		os.Getenv("MYSQL_USERS_TABLES_PASSWORD"), os.Getenv("MYSQL_USERS_TABLES_SETTING"))
	tQueryValues := getQueryValues(aSignUpUser.Len()) //カラムの数とmodelの要素の数が揃っていることを保証する
	tQuery := fmt.Sprintf("INSERT INTO %s %s VALUES %s",
		os.Getenv("MYSQL_USERS_TABLES"), tColumns, tQueryValues)

	tStmt, tError := aDb.Prepare(tQuery)
	if tError != nil {
		return http.StatusInternalServerError, tError
	}
	defer tStmt.Close()

	tValues := aSignUpUser.GetAll()

	if _, tError := tStmt.Exec(tValues[0], tValues[1], tValues[2], tValues[3]); tError != nil {
		return http.StatusInternalServerError, tError
	}

	return 0, nil
}

func getQueryValues(aLen int) string {
	tValues := "("
	for i := 0; i < aLen-1; i++ {
		tValues += "?,"
	}
	tValues += "?)"
	return tValues
}
