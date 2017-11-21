package service_test

import (
	"testing"

	"github.com/golangProject/src/domain"

	"github.com/golangProject/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	var tweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	service.PublishTweet(tweet)

	publishedTweet := service.GetTweet()

	if publishedTweet.User != user && publishedTweet.Text != text {
		t.Error("Expected tweet is", tweet)
	}

	if publishedTweet.Date == nil {
		t.Error("Expected a date")
	}

}
func TestClean(t *testing.T) {
	var tweet2 *domain.Tweet

	user2 := "grupoesfera"
	text2 := "This is my first tweet"
	tweet2 = domain.NewTweet(user2, text2)

	service.PublishTweet(tweet2)
	service.CleanTweet()

	if service.GetTweet().Text != "" {
		t.Error("Expected tweet is", tweet2)
	}
}
