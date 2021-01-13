package main

import (
	"fmt"
	"net/http"

	controllers "unit-tested-controllers.com/controllers"
	externals "unit-tested-controllers.com/externals"
)

const APP_PORT = 8090

func main() {

	http.HandleFunc("/hello", controllers.HelloController)
	http.HandleFunc("/bye", controllers.ByeController)

	apiClient := externals.ApiClient{"https://jsonplaceholder.typicode.com"}

	http.HandleFunc("/call_api_placeholder", controllers.CallApiPlaceholder(&apiClient))

	fmt.Println(fmt.Sprintf("Starting server at 0.0.0.0:%d", APP_PORT))
	http.ListenAndServe(fmt.Sprintf(":%d", APP_PORT), nil)
}
