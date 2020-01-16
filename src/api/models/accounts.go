package models

import (
	"fmt"
	"os"

	u "github.com/anyric/bts/src/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Token JWT Claims structure
type Token struct {
	UserID uint
	jwt.StandardClaims
}

// Account detail struct
type Account struct {
	gorm.Model
	Username string
	Password string
	Token    string
}

// Validate user accounts
func (account *Account) Validate() (map[string]interface{}, bool) {
	if len(account.Username) < 6 {
		return u.Message(false, "Username is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//username and password must be unique
	temp := &Account{}

	//check for errors and duplicate username and password
	err := db.Table("accounts").Where("Username = ?", account.Username).First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if temp.Username != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirements passed"), true
}

// Create a new user account
func (account *Account) Create() map[string]interface{} {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	db.Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_SALT")))
	account.Token = tokenString

	account.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

// Login user to account
func Login(username, password string) map[string]interface{} {

	account := &Account{}
	err := db.Table("accounts").Where("username = ?", username).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "username not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_SALT")))
	account.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

// GetUser with a given ID
func GetUser(u uint) *Account {

	acc := &Account{}
	err := db.Table("accounts").Where("id = ?", u).First(acc).Error
	if err != nil {
		return nil
	}

	acc.Password = ""
	return acc
}

// GetUsers retrieve all users
func GetUsers(user uint) []*Account {

	accounts := make([]*Account, 0)
	err := db.Table("accounts").Where("user_id = ?", user).Find(&accounts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return accounts
}
