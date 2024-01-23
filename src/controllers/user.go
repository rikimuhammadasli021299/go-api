package controllers

import (
	"encoding/json"
	"fmt"
	"go-api/src/helper"
	"go-api/src/models"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		var input models.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Invalid request body")
			return
		}

		hashedPassword,_ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		password := string(hashedPassword)
		newUser := models.User{
			Email: input.Email,
			Password: password,
		}

		res := models.CreateUser(&newUser)
		var _, _ = json.Marshal(res)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Register Successfull")
	} else{
		http.Error(w, "", http.StatusBadRequest)
	}
}

func Login(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
	} else{
		var input models.User
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Invalid request body")
			return
		}

		ValidateEmail := models.FindEmail(&input)
		if len(ValidateEmail) == 0 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "Email not found")
			return
		}

		var passwordSecond string
		for _, user := range ValidateEmail {
			passwordSecond = user.Password
		}
		if err := bcrypt.CompareHashAndPassword([]byte(passwordSecond), []byte(input.Password)); err != nil{
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Incorrect password")
			return
		}
		jwtKey := os.Getenv("JWT_SECRET_KEY")
		token, err := helper.GenerateToken(jwtKey, input.Email)
		item := map[string]string{
			"Email": input.Email,
			"Token": token,
		}
		res, _ := json.Marshal(item)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		return

	}
}