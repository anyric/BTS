package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"api/models"
)
//Login users
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	account := models.Account{}
	postBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(postBody, &account)

	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		acc := models.Login(account.Mobile, account.Password)
		response, _ := json.Marshal(acc)
		w.Write(response)
	}

}
