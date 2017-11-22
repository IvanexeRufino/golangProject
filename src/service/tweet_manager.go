package service

import (
	"fmt"

	"github.com/golangProject/src/domain"
)

//Tweet es un tweet
var tweet *domain.Tweet

//PublishTweet qe hace nada
func PublishTweet(tweet2 *domain.Tweet) error {
	var err error
	if tweet2.User == "" {
		err = fmt.Errorf("user is required")
	} else {
		tweet = tweet2
	}
	return err
}

//GetTweet getter
func GetTweet() *domain.Tweet {
	return tweet
}

//CleanTweet limpia el texto
func CleanTweet() {
	tweet = nil
}
