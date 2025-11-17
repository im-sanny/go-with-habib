package handlers

import (
	"log"
	"net/http"
)


func Test(w http.ResponseWriter, r *http.Request){

	log.Println("I'm handler: I'll be print in the middle")
}
