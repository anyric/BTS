package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"api/models"
)

//GetUsers List all users
func GetUsers(w http.ResponseWriter, r *http.Request) {

	accouunts, err := models.GetUsers()
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	if len(accouunts) > 0 {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(accouunts)
		w.Write(response)
	}
}

//CreateUser add new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var accouunt models.Account
	postBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(postBody, &accouunt)

	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		accouunt.Create()
		response, _ := json.Marshal(accouunt)
		w.Write(response)
	}
}

//GetUser returns a single user
func GetUser(w http.ResponseWriter, r *http.Request) {
	pattern, _ := regexp.Compile(`/users/(\d+)`)
	matches := pattern.FindStringSubmatch(r.URL.Path)

	if len(matches) > 0 {
		ids, _ := strconv.Atoi(matches[1])
		id := uint(ids)
		w.Header().Set("Content-Type", "application/json")
		accouunt, err := models.GetUser(id)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			response, _ := json.Marshal(accouunt)
			w.Write(response)
		}
	}
}

//UpdateUser edit a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	pattern, _ := regexp.Compile(`/users/(\d+)`)
	matches := pattern.FindStringSubmatch(r.URL.Path)

	if len(matches) > 0 {
		ids, _ := strconv.Atoi(matches[1])
		id := uint(ids)
		w.Header().Set("Content-Type", "application/json")
		account := models.Account{}
		postBody, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(postBody, &account)

		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Header().Set("Content-Type", "application/json")
			row, err := models.UpdateUser(id, account)
			
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			if row != 0 {
				acc, _ := models.GetUser(id)
				response, _ := json.Marshal(acc)
				w.Write(response)
			}
			
		}
	}
}

//DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	pattern, _ := regexp.Compile(`/users/(\d+)`)
	matches := pattern.FindStringSubmatch(r.URL.Path)

	if len(matches) > 0 {
		ids, _ := strconv.Atoi(matches[1])
		id := uint(ids)
		w.Header().Set("Content-Type", "application/json")
		accouunt, err := models.DeleteUser(id)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			response, _ := json.Marshal(accouunt)
			w.Write(response)
		}
	}
}

