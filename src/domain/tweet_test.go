package domain_test

import (
	"testing"

	"github.com/golangProject/src/domain"
)

func TestCanGetAPrintableTweet(t *testing.T) {
	tweet := domain.NewTweet("grupoesfera", "This is my tweet")

	text := tweet.PrintableTweet()

	text2 := tweet.String()

	expectedText := "@grupoesfera: This is my tweet"

	if text != expectedText && text2 != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}
}
