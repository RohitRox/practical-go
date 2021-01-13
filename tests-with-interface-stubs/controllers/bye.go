package controllers

import (
	"fmt"
	"net/http"
)

func ByeController(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Bye!")
}
