package user

import (
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler)  Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginReq)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	usr := database.Find(loginReq.Email, loginReq.Password)
	if usr == nil {
		http.Error(w, "Invalid Credentials", http.StatusBadRequest)
		return
	}

	cnf := config.GetConfig()

	accessToken, err := util.CreateJwt(cnf.JWTSecretKey, util.Payload{
		Sub:      usr.ID,
		FirsName: usr.FirstName,
		LastName: usr.LastName,
		Email:    usr.Email,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	util.SendData(w, accessToken, http.StatusCreated)
}
