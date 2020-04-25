package jokes

import (
	"fmt"
	"net/http"

	_ "gopkg.in/square/go-jose.v2/json"

	"workshop/internal/api"
)

const getJokePath = "/api?format=json"

// JokeClient is a joke API client
type JokeClient struct {
	url string
}

// NewJokeClient creates a new joke client
func NewJokeClient(baseUrl string) *JokeClient {
	return &JokeClient{
		url: baseUrl,
	}
}

func (jc *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url + getJokePath

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request status: %s", http.StatusText(resp.StatusCode))
	}
}
