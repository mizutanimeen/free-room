package models

import (
	"fmt"
	"net/http"
	"reflect"
	"time"
)

// 要件チェックで利用中。type変える時はここも変える
const (
	EMAIL    = "Email"
	NAME     = "Name"
	PASSWORD = "Password"
	SETTING  = "Setting"
)

type User struct {
	UserID   int       `json:"userid"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Setting  int       `json:"setting"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type SignUpUser struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Setting  int    `json:"setting"`
}

// 値が要件を満たしているかチェック
func SignUpConditionCheck(aSignUpUser *SignUpUser) (int, error) {
	tValues := reflect.ValueOf(*aSignUpUser)
	tTypes := tValues.Type()
	for i := 0; i < tTypes.NumField(); i++ {
		tTypeName := tTypes.Field(i).Name
		tVal := tValues.FieldByName(tTypeName)

		switch tTypeName {
		case EMAIL:
			if tStatus, tError := emailCondition(tVal); tError != nil {
				return tStatus, tError
			}
		case NAME:
			if tStatus, tError := nameCondition(tVal); tError != nil {
				return tStatus, tError
			}
		case PASSWORD:
			if tStatus, tError := passwordCondition(tVal); tError != nil {
				return tStatus, tError
			}
		case SETTING:
			if tStatus, tError := settingCondition(tVal); tError != nil {
				return tStatus, tError
			}
		default:
			return http.StatusInternalServerError, fmt.Errorf("%s is not defined", tTypeName)
		}
	}

	return 0, nil
}

func emailCondition(aValue reflect.Value) (int, error) {
	// tValue := aValue.String()

	//@とかメアドの判定
	//どうせメール送るけど一応ね

	return 0, nil
}

func nameCondition(aValue reflect.Value) (int, error) {
	tValue := aValue.String()

	if !(1 <= len(tValue) && len(tValue) <= 12) {
		return http.StatusBadRequest, fmt.Errorf("Name is too long or too short")
	}

	return 0, nil
}

func passwordCondition(aValue reflect.Value) (int, error) {
	tValue := aValue.String()

	// 数字と英文字が両方含まれているか

	if !(6 <= len(tValue) && len(tValue) <= 12) {
		return http.StatusBadRequest, fmt.Errorf("Password is too long or too short")
	}

	return 0, nil
}
func settingCondition(aValue reflect.Value) (int, error) {
	// tValue := aValue.Int()
	return 0, nil
}
