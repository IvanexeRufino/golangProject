package service_test

import (
	"testing"

	"github.com/golangProject/src/domain"

	"github.com/golangProject/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	tm := service.NewTweetManager()

	var tweet *domain.TextTweet

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	tm.PublishTweet(tweet)

	publishedTweets := tm.GetTweets()

	if publishedTweets[0].GetUser() != user && publishedTweets[0].GetText() != text {
		t.Error("Expected tweet is", tweet)
	}

	if publishedTweets[0].GetDate() == nil {
		t.Error("Expected a date")
	}

}
func TestClean(t *testing.T) {

	tm := service.NewTweetManager()

	var tweet2 domain.Tweet

	user2 := "grupoesfera"
	text2 := "This is my first tweet"
	tweet2 = domain.NewTextTweet(user2, text2)

	tm.PublishTweet(tweet2)
	tm.CleanTweet()

	if tm.GetTweets() != nil {
		t.Error("Expected tweet is", tweet2)
	}
}

func TestGetLastTweetReturnsLastOne(t *testing.T) {

	tm := service.NewTweetManager()

	var tweet, secondTweet domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)

	tm.PublishTweet(tweet)
	secondId, _ := tm.PublishTweet(secondTweet)

	lastTweet := tm.GetLastTweet()

	if !isValidTweet(t, lastTweet, secondId, user, secondText) {
		return
	}

}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	tm := service.NewTweetManager()

	var tweet domain.Tweet
	var user string
	text := "this is my first tweet"
	tweet = domain.NewTextTweet(user, text)
	var err error
	_, err = tm.PublishTweet(tweet)

	if err != nil && err.Error() != "user is required" {
		t.Error("User is required")
	}

}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	tm := service.NewTweetManager()

	var tweet domain.Tweet
	user := "grupoesfera"
	var text string

	tweet = domain.NewTextTweet(user, text)
	var err error
	_, err = tm.PublishTweet(tweet)

	if err == nil {
		t.Error("Expected error")
		return
	}
	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}

}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	tm := service.NewTweetManager()

	var tweet domain.Tweet

	user := "grupoesfera"
	text := `AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
			AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA`

	tweet = domain.NewTextTweet(user, text)

	var err error
	_, err = tm.PublishTweet(tweet)

	if err == nil {
		t.Error("Expected error")
		return
	}
	if err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text is required")
	}

}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	tm := service.NewTweetManager()

	var tweet, secondTweet domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)

	tm.PublishTweet(tweet)
	tm.PublishTweet(secondTweet)

	publishedTweets := tm.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstPublishedTweet.GetID(), user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondPublishedTweet.GetID(), user, secondText) {
		return
	}

}

func TestCanRetrieveTweetById(t *testing.T) {

	tm := service.NewTweetManager()

	var tweet domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	id, _ = tm.PublishTweet(tweet)

	publishedTweet := tm.GetTweetByID(id)

	isValidTweet(t, publishedTweet, id, user, text)

	publishedTweet2 := tm.GetTweetByID(50)

	if publishedTweet2 != nil {
		t.Errorf("Expected nil")
	}
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	tm := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	tm.PublishTweet(tweet)
	tm.PublishTweet(secondTweet)
	tm.PublishTweet(thirdTweet)

	count := tm.CountTweetsByUser(user)

	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	tm := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	firstId, _ := tm.PublishTweet(tweet)
	secondId, _ := tm.PublishTweet(secondTweet)
	tm.PublishTweet(thirdTweet)

	tweets := tm.GetTweetsByUser(user)

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

	tm := service.NewTweetManager()

	var tweet, secondTweet domain.Tweet

	user := "nportas"
	anotherUser := "mercadolibre"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(anotherUser, secondText)

	tm.AddUser(user)
	tm.AddUser(anotherUser)

	firstId, _ := tm.PublishTweet(tweet)
	tm.PublishTweet(secondTweet)

	tm.Follow("nportas", "nportas")

	timeline := tm.GetTimeline("nportas")

	if len(timeline) != 1 {
		t.Errorf("Expected size is 1 but was %d", len(timeline))
		return
	}

	firstPublishedTweet := timeline[0]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	err := tm.Follow("grupoesfera", "mas de lo mismo")

	if err == nil {
		t.Errorf("Expected an error")
	}

}

func TestSendMessageToUser(t *testing.T) {

	tm := service.NewTweetManager()

	var tweet, secondTweet domain.Tweet

	user := "nportas"
	anotherUser := "mercadolibre"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(anotherUser, secondText)

	tm.AddUser(user)
	tm.AddUser(anotherUser)

	tm.PublishTweet(tweet)
	tm.PublishTweet(secondTweet)

	tm.SendMessage(user, anotherUser, "hola wachin")

	messages := tm.GetUnreadedMessages(anotherUser)

	if len(messages) != 1 {
		t.Errorf("Expected size is 1 but was %d", len(messages))
		return
	}

	messagesReadeds := tm.GetAllDirectMessages(anotherUser)

	if len(messagesReadeds) != 1 {
		t.Errorf("Expected size is 1 but was %d", len(messagesReadeds))
		return
	}

	messagesUnreadeds := tm.GetUnreadedMessages(anotherUser)

	if len(messagesUnreadeds) != 0 {
		t.Errorf("Expected size is 0 but was %d", len(messagesUnreadeds))
		return
	}

	err := tm.SendMessage(user, "mas de lo mismo", "otra cosa")

	if err == nil {
		t.Errorf("Expected an error")
	}

}

func TestTrendingTopicsAreTheMoreTweeted(t *testing.T) {

	tm := service.NewTweetManager()

	var tweet, secondTweet domain.Tweet

	user := "nportas"
	anotherUser := "mercadolibre"
	text := "This is my first tweet"
	secondText := "This was his second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(anotherUser, secondText)

	tm.AddUser(user)
	tm.AddUser(anotherUser)

	tm.PublishTweet(tweet)
	tm.PublishTweet(secondTweet)

	tts := tm.GetTrendingTopics()

	if !(tts[0] == "This" || tts[0] == "tweet") && !(tts[1] == "tweet" || tts[1] == "This") {
		t.Errorf("Expected this and tweet but was %s and %s", tts[0], tts[1])
	}
}

func TestRetweetearAddsToTweets(t *testing.T) {
	tm := service.NewTweetManager()

	var tweet, secondTweet domain.Tweet

	user := "nportas"
	anotherUser := "mercadolibre"
	text := "This is my first tweet"
	secondText := "This was his second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(anotherUser, secondText)

	tm.AddUser(user)
	tm.AddUser(anotherUser)

	id, _ := tm.PublishTweet(tweet)
	tm.PublishTweet(secondTweet)

	tm.Retweetear(user, id)

	if len(tm.Writer.GetTweets()[user]) != 2 {
		t.Errorf("Expected a retweeted tweet")
	}
}

func TestFavouriteList(t *testing.T) {

	tm := service.NewTweetManager()

	var tweet, secondTweet domain.Tweet

	user := "nportas"
	anotherUser := "mercadolibre"
	text := "This is my first tweet"
	secondText := "This was his second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(anotherUser, secondText)

	tm.AddUser(user)
	tm.AddUser(anotherUser)

	id, _ := tm.PublishTweet(tweet)
	tm.PublishTweet(secondTweet)

	tm.Fav(user, id)

	favourites := tm.GetTweetsFav(user)

	if len(favourites) != 1 {
		t.Errorf("Expected a favourite tweet")
	}

}

func TestAddingPlugins(t *testing.T) {

	tm := service.NewTweetManager()

	var tweet, secondTweet domain.Tweet

	user := "nportas"
	anotherUser := "mercadolibre"
	text := "This is my first tweet"
	secondText := "This was his second tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(anotherUser, secondText)

	tm.AddUser(user)
	tm.AddUser(anotherUser)

	face := service.NewPlugin(1)
	google := service.NewPlugin(2)
	tweeter := service.NewPlugin(3)

	tm.AddPlugin(face)
	tm.AddPlugin(google)
	tm.AddPlugin(tweeter)

	tm.PublishTweet(tweet)
	tm.PublishTweet(secondTweet)

	plugins := tm.Plugins

	if len(plugins) != 3 {
		t.Errorf("Expected all plugins")
	}

}

func isValidTweet(t *testing.T, tweet domain.Tweet, id int, user, text string) bool {

	if tweet.GetUser() != user && tweet.GetText() != text && tweet.GetID() != id {
		t.Errorf("Expected tweet is %s: %s %d \nbut is %s: %s %d",
			user, text, id, tweet.GetUser(), tweet.GetText(), tweet.GetID())
		return false
	}

	if tweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}
