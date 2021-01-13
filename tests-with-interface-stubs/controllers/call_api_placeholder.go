package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"unit-tested-controllers.com/externals"
)

func CallApiPlaceholder(apiClient externals.IApiClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		posts, err := apiClient.GetPosts()

		if err != nil {
			fmt.Fprint(w, "Oops! Something went wrong!")
		} else {
			jData, err := json.Marshal(posts)
			if err != nil {
				fmt.Fprint(w, "Oops! Something went wrong!")
			}
			w.Write(jData)
		}

	}
}
