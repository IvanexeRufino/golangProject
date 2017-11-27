package service

import (
	"github.com/golangProject/src/domain"
)

//ChannelTweetWriter struct
type ChannelTweetWriter struct {
	Writer TweetWriter
}

//TweetWriter interface
type TweetWriter interface {
	WriteTweet(chan domain.Tweet, chan bool)
}

//MemoryTweetWriter struct
type MemoryTweetWriter struct {
	Tweets []domain.Tweet
}

//FileTweetWriter struct
type FileTweetWriter struct {
}

//NewMemoryTweetWriter constructor
func NewMemoryTweetWriter() *MemoryTweetWriter {
	tw := MemoryTweetWriter{
		make([]domain.Tweet, 0),
	}

	return &tw
}

//NewChannelTweetWriter constructor
func NewChannelTweetWriter(tw TweetWriter) *ChannelTweetWriter {
	ctw := ChannelTweetWriter{
		tw,
	}

	return &ctw
}

//WriteTweet paralelized
func (ctw *ChannelTweetWriter) WriteTweet(tweets chan domain.Tweet, sync chan bool) {
	ctw.Writer.WriteTweet(tweets, sync)
}

//WriteTweet writing
func (mtw *MemoryTweetWriter) WriteTweet(tweets chan domain.Tweet, sync chan bool) {
	mtw.Tweets = append(mtw.Tweets, <-tweets)
}
