package user

import (
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	usr, err := h.userRepo.Find(req.Email, req.Password)
	if err != nil {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if usr == nil {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	accessToken, err := util.CreateJwt(h.cnf.JWTSecretKey, util.Payload{
		Sub:      usr.ID,
		FirsName: usr.FirstName,
		LastName: usr.LastName,
		Email:    usr.Email,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	util.SendData(w, http.StatusCreated, accessToken)
}
