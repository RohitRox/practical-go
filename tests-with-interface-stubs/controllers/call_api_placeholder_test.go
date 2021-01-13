package controllers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"

	"unit-tested-controllers.com/models"
)

type StubApiClient struct {}

func (s *StubApiClient) GetPost(id int) (*models.Post, error)  {
	return &models.Post{
		UserId: 10,
		Id: 1,
		Title: "Title 1",
		Body: "Body 1",
	}, nil
}

func (s *StubApiClient) GetPosts() (*[]models.Post, error) {

	post := models.Post{
		UserId: 10,
		Id: 1,
		Title: "Title 1",
		Body: "Body 1",
	}

	posts := []models.Post{ post }

	return &posts, nil

}
func TestCallApiPlaceholder(t *testing.T) {

	t.Run("Get /call_api_placeholder", func (t *testing.T){

		stub := &StubApiClient{}

		req, _ := http.NewRequest("GET", "/call_api_placeholder", nil)
		handler := http.HandlerFunc(CallApiPlaceholder(stub))

		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		var posts []models.Post

		json.Unmarshal([]byte(rec.Body.String()), &posts)

		if !(posts[0].Id == 1 && posts[0].UserId == 10) {
			t.Errorf("Did not get expected post")
		}

	})
}
