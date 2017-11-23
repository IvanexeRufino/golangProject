package service

import (
	"github.com/golangProject/src/domain"
)

//NewTweetManager constructor
func NewTweetManager() *domain.TweetManager {
	tweetMan := domain.TweetManager{}
	tweetMan.InitializeService()

	return &tweetMan
}
