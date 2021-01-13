package externals

import (
	"fmt"
	"net/http"
	"encoding/json"

	"unit-tested-controllers.com/models"
)

type ApiClient struct {
	BaseUrl string
}

type IApiClient interface {
	GetPost(id int) (*models.Post, error)
	GetPosts() (*[]models.Post, error)
}

func (a *ApiClient) GetPost(id int) (*models.Post, error) {
	url := fmt.Sprintf(a.BaseUrl+"/posts/%d", id)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var r models.Post

	err = json.NewDecoder(resp.Body).Decode(&r)

	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (a *ApiClient) GetPosts() (*[]models.Post, error) {
	url := fmt.Sprintf(a.BaseUrl+"/posts")

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var r []models.Post

	err = json.NewDecoder(resp.Body).Decode(&r)

	if err != nil {
		return nil, err
	}

	return &r, nil
}
