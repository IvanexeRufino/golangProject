package service_test

import (
	"testing"

	"github.com/golangProject/src/domain"
	"github.com/golangProject/src/service"
)

func TestWriter(t *testing.T) {

	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "Async tweet")
	tweet2 := domain.NewTextTweet("grupoesfera", "Async tweet2")

	memoryTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(memoryTweetWriter)

	tweetsToWrite := make(chan domain.Tweet)
	quit := make(chan bool)

	go tweetWriter.WriteTweet(tweetsToWrite, quit)

	// Operation
	tweetsToWrite <- tweet
	tweetsToWrite <- tweet2
	close(tweetsToWrite)

	<-quit

	// Validation
	if memoryTweetWriter.Tweets[0] != tweet {
		t.Errorf("A tweet in the writer was expected")
	}

	if memoryTweetWriter.Tweets[1] != tweet2 {
		t.Errorf("A tweet in the writer was expected")
	}
}
