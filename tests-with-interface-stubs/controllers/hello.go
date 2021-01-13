package controllers

import (
	"fmt"
	"net/http"
)

func HelloController(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
