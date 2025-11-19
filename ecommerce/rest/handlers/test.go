package handlers

import (
	"log"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	log.Println("I'm test handler!")
}
