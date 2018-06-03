package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorkemsari/golang-rest-api/helper"
	"github.com/gorkemsari/golang-rest-api/model"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secret")

func Token(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var u model.User
	err := decoder.Decode(&u)
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = u.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString(mySigningKey)

	helper.RespondWithJSON(w, http.StatusOK, tokenString)
}
