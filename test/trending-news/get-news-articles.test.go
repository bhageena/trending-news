package getnewsarticles

import (
	"testing"
    "../../packages/trending-news/get-news-articles/index"
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

func TestOkRequest(t *testing.T) {
    response := main()
    if response.statusCode != 200  {
       t.Errorf("Request was not successful")
    }
}

