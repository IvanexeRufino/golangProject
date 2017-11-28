package service_test

import (
	"testing"

	"github.com/golangProject/src/domain"

	"github.com/golangProject/src/service"
)

func BenchmarkPublishTweet(b *testing.B) {
	fileTweetWriter := service.NewMemoryTweetWriter()
	tweetWriter := service.NewChannelTweetWriter(fileTweetWriter)
	tweetManager := service.NewTweetManager()
	tweetManager.Writer = tweetWriter

	tweet := domain.NewTextTweet("hola", "hola")

	for n := 0; n < b.N; n++ {
		tweetManager.PublishTweet(tweet)
	}

}
