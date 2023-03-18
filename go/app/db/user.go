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

func GetAllUsers(aDb *sql.DB) ([]*models.User, int, error) {
	tUsers, tStatus, tError := getUsersByWhere("", aDb)
	if tError != nil {
		return nil, tStatus, tError
	}
	return tUsers, 0, nil
}

func GetUserByID(aID string, aDb *sql.DB) (*models.User, int, error) {
	tWhere := fmt.Sprintf("WHERE %s = %s", os.Getenv("MYSQL_USERS_TABLES_ID"), aID)
	tUsers, tStatus, tError := getUsersByWhere(tWhere, aDb)
	if tError != nil {
		return nil, tStatus, tError
	}

	if len(tUsers) <= 0 {
		return nil, http.StatusBadRequest, fmt.Errorf("%s is not found", aID)
	} else if len(tUsers) >= 2 {
		return nil, http.StatusInternalServerError, fmt.Errorf("%s is not unique", aID)
	}

	return tUsers[0], 0, nil
}

func getUsersByWhere(aWhere string, aDb *sql.DB) ([]*models.User, int, error) {
	tQuery := fmt.Sprintf("SELECT * FROM %s %s", os.Getenv("MYSQL_USERS_TABLES"), aWhere)
	tSqlRows, tError := aDb.Query(tQuery)
	if tError != nil {
		return nil, http.StatusInternalServerError, tError
	}
	defer tSqlRows.Close()

	tUsers := []*models.User{}
	for tSqlRows.Next() {
		var tUser models.User

		if tError := tSqlRows.Scan(&tUser.UserID, &tUser.Email, &tUser.Name, &tUser.Password, &tUser.Setting, &tUser.CreateAt, &tUser.UpdateAt); tError != nil {
			return nil, http.StatusInternalServerError, tError
		}

		tUsers = append(tUsers, &tUser)
	}
	return tUsers, 0, nil
}
