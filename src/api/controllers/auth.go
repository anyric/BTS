package controllers

import (
	"api/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Login users
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	account := models.Account{}
	postBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(postBody, &account)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		acc := models.Login(account.Mobile, account.Password)
		response, _ := json.Marshal(acc)
		w.Write(response)
	}
}
