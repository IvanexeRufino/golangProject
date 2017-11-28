package service

import (
	"os"

	"github.com/golangProject/src/domain"
)

//ChannelTweetWriter struct
type ChannelTweetWriter struct {
	Writer TweetWriter
}

//TweetWriter interface
type TweetWriter interface {
	WriteTweet(domain.Tweet)
	GetList() map[string][]domain.Tweet
}

//MemoryTweetWriter struct
type MemoryTweetWriter struct {
	Tweets map[string][]domain.Tweet
}

//FileTweetWriter struct
type FileTweetWriter struct {
	File *os.File
}

//NewMemoryTweetWriter constructor
func NewMemoryTweetWriter() *MemoryTweetWriter {
	tw := MemoryTweetWriter{
		make(map[string][]domain.Tweet),
	}

	return &tw
}

//NewFileTweetWritter constructor
func NewFileTweetWritter() *FileTweetWriter {

	file, _ := os.OpenFile(
		"tweets.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)

	writer := new(FileTweetWriter)
	writer.File = file

	return writer
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

	tweet, open := <-tweets

	for open {
		ctw.Writer.WriteTweet(tweet)

		tweet, open = <-tweets
	}
	sync <- true
}

//GetTweets getter
func (ctw *ChannelTweetWriter) GetTweets() map[string][]domain.Tweet {

	return ctw.Writer.GetList()
}

//WriteTweet writing
func (mtw *MemoryTweetWriter) WriteTweet(tweets domain.Tweet) {

	mtw.Tweets[tweets.GetUser()] = append(mtw.Tweets[tweets.GetUser()], tweets)
}

//GetList getter
func (mtw *MemoryTweetWriter) GetList() map[string][]domain.Tweet {
	return mtw.Tweets

}

//WriteTweet writting file
func (ftw *FileTweetWriter) WriteTweet(tweets domain.Tweet) {
	if ftw.File != nil {
		byteSlice := []byte(tweets.PrintableTweet() + "\n")
		ftw.File.Write(byteSlice)
	}

}
