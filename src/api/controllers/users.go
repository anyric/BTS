package controllers

import (
	"net/http"
)

//GetUsers List all users
func GetUsers(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("List users"))
}

//CreateUser add new user
func CreateUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Create user"))
}

//GetUser returns a single user
func GetUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("A user"))
}

//UpdateUser edit a user
func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Update user"))
}

//DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Delete user"))
}