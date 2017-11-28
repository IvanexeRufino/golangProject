package domain_test

import (
	"testing"

	"github.com/golangProject/src/domain"
	"github.com/golangProject/src/service"
)

func TestAHundredPercentCoverage(t *testing.T) {

	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	tweet2 := domain.NewTextTweet("grupoesfera2", "This is my tweet")
	tweet3 := domain.NewImageTweet("algo", "otro", "nuevo")
	tweet4 := domain.NewQuoteTweet("algo", "otro", tweet3)

	us := domain.NewUser("hola")

	tm := service.NewTweetManager()

	dm := domain.NewDirectMessage(us.Name, "otro")

	tm.PublishTweet(tweet)
	tm.PublishTweet(tweet2)
	tm.PublishTweet(tweet3)
	tm.PublishTweet(tweet4)
	string1 := tweet3.String()
	string2 := tweet4.String()

	tm.SendMessage("grupoesfera", "grupoesfera2", "otro")

	if string1 != tweet3.PrintableTweet() {
		t.Errorf("Expected to be the same")
	}

	if string2 != tweet4.PrintableTweet() {
		t.Errorf("Expected to be the same")
	}

	if tweet3.GetDate() == nil {
		t.Error("Expected a date")
	}

	if tweet.GetDate() == nil {
		t.Error("Expected a date")
	}

	if tweet4.GetDate() == nil {
		t.Error("Expected a date")
	}

	if dm.From == "otro" {
		t.Errorf("goodbye")
	}

}

func TestTextTweetPrintsUserAndText(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	tweet := domain.NewImageTweet("grupoesfera", "This is my image", "http://www.grupoesfera.com.ar/common/img/grupoesfera.png")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my image http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {

	// Initialization
	quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	tweet := domain.NewQuoteTweet("nick", "Awesome", quotedTweet)

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}

}
