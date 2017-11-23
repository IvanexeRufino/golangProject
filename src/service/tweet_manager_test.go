package service_test

import (
	"testing"

	"github.com/golangProject/src/domain"

	"github.com/golangProject/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	service.InitializeService()

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

	service.InitializeService()

	var tweet2 *domain.Tweet

	user2 := "grupoesfera"
	text2 := "This is my first tweet"
	tweet2 = domain.NewTweet(user2, text2)

	service.PublishTweet(tweet2)
	service.CleanTweet()

	if service.GetTweets() != nil {
		t.Error("Expected tweet is", tweet2)
	}
}

func TestGetLastTweetReturnsLastOne(t *testing.T) {

	service.InitializeService()

	var tweet, secondTweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)

	service.PublishTweet(tweet)
	secondId, _ := service.PublishTweet(secondTweet)

	lastTweet := service.GetLastTweet()

	if !isValidTweet(t, lastTweet, secondId, user, secondText) {
		return
	}

}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	var user string
	text := "this is my first tweet"
	tweet = domain.NewTweet(user, text)
	var err error
	_, err = service.PublishTweet(tweet)

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
	_, err = service.PublishTweet(tweet)

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
	_, err = service.PublishTweet(tweet)

	if err == nil {
		t.Error("Expected error")
		return
	}
	if err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text is required")
	}

}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	service.InitializeService()

	var tweet, secondTweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)

	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	publishedTweets := service.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstPublishedTweet.ID, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondPublishedTweet.ID, user, secondText) {
		return
	}

}

func TestCanRetrieveTweetById(t *testing.T) {

	service.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	id, _ = service.PublishTweet(tweet)

	publishedTweet := service.GetTweetByID(id)

	isValidTweet(t, publishedTweet, id, user, text)

	publishedTweet2 := service.GetTweetByID(50)

	if publishedTweet2 != nil {
		t.Errorf("Expected nil")
	}
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	service.InitializeService()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)

	count := service.CountTweetsByUser(user)

	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	service.InitializeService()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	firstId, _ := service.PublishTweet(tweet)
	secondId, _ := service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)

	tweets := service.GetTweetsByUser(user)

	if len(tweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(tweets))
		return
	}

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}

}
func TestFollowuser(t *testing.T) {
	service.InitializeService()

	var tweet, secondTweet *domain.Tweet

	user := "nportas"
	anotherUser := "mercadolibre"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(anotherUser, secondText)

	firstId, _ := service.PublishTweet(tweet)
	secondId, _ := service.PublishTweet(secondTweet)

	service.Follow("grupoesfera", "nportas")
	service.Follow("grupoesfera", "mercadolibre")

	timeline := service.GetTimeline("grupoesfera")

	if len(timeline) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(timeline))
		return
	}

	firstPublishedTweet := timeline[0]
	secondPublishedTweet := timeline[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}

}

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user, text string) bool {

	if tweet.User != user && tweet.Text != text && tweet.ID != id {
		t.Errorf("Expected tweet is %s: %s %d \nbut is %s: %s %d",
			user, text, id, tweet.User, tweet.Text, tweet.ID)
		return false
	}

	if tweet.Date == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}
