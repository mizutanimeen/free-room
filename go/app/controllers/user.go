package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mizutanimeen/free-room/models"
)

func GetSignUpUser(aRequest *http.Request) (*models.SignUpUser, int, error) {
	var tSignUpUser *models.SignUpUser
	switch aRequest.Header["Content-Type"][0] {
	case "application/json":
		var tStatus int
		var tError error
		tSignUpUser, tStatus, tError = jsonToSignUpUser(aRequest)
		if tError != nil {
			return nil, tStatus, tError
		}
	default:
		return nil, http.StatusPreconditionFailed, fmt.Errorf("%s is bad content-type", aRequest.Header["Content-Type"][0])
	}

	return tSignUpUser, 0, nil
}

// bodyがjsonのrequestからstruct userを作る
func jsonToSignUpUser(aRequest *http.Request) (*models.SignUpUser, int, error) {
	tBody := make([]byte, aRequest.ContentLength)
	aRequest.Body.Read(tBody)

	tSignUpUser := &models.SignUpUser{}
	if tError := json.Unmarshal(tBody, tSignUpUser); tError != nil {
		return nil, http.StatusInternalServerError, tError
	}

	return tSignUpUser, 0, nil
}
