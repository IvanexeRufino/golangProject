package service_test

import (
	"testing"

	"github.com/golangProject/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	tweet := "This is my first tweet"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}

}
func TestClean(t *testing.T) {
	tweet2 := ""
	service.PublishTweet("aloja")
	service.CleanTweet()

	if service.GetTweet() != tweet2 {
		t.Error("Expected tweet is", tweet2)
	}
}
