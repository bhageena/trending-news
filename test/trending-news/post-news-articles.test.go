package postnewsarticles

import (
	"testing"
    "../../packages/trending-news/post-news-articles/index"
)
        
func TestMain(t *testing.T) {
    response := main()
    if response != "" {
       t.Errorf("Actual response body doesn't match expected response")
    }

    expected := ""
    if response.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            response.Body.String(), expected)
    }
}

func TestPostRequest(t *testing.T) {
    response := main()
    if response.code != 201 || response.code != 202  {
       t.Errorf("POST request was not successful")
    }
}

