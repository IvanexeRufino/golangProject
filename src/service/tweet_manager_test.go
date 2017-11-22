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

	publishedTweets := service.GetTweets()

	if publishedTweets[0].User != user && publishedTweets[0].Text != text {
		t.Error("Expected tweet is", tweet)
	}

	if publishedTweets[0].Date == nil {
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

	if service.GetTweets()[0].Text != "" {
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

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	// Initialization
	service.InitializeService()

	var tweet, secondTweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)

	// Operation
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, user, secondText) {
		return
	}

}

func isValidTweet(t *testing.T, tweet *domain.Tweet, user, text string) bool {

	if tweet.User != user && tweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, tweet.User, tweet.Text)
		return false
	}

	if tweet.Date == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}
