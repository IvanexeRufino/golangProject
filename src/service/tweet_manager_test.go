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

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	var user string
	text := "this is my first tweet"
	tweet = domain.NewTweet(user, text)
	var err error
	err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "user is required" {
		t.Error("User is required")
	}

}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	user := "grupoesfera"
	var text string

	tweet = domain.NewTweet(user, text)
	var err error
	err = service.PublishTweet(tweet)

	if err == nil {
		t.Error("Expected error")
		return
	}
	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}

}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet

	user := "grupoesfera"
	text := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"

	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err == nil {
		t.Error("Expected error")
		return
	}
	if err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text is required")
	}

}
