package service

import (
	"fmt"

	"github.com/golangProject/src/domain"
)

//Tweet es un tweet
var tweets []*domain.Tweet

//PublishTweet qe hace nada
func PublishTweet(tweet2 *domain.Tweet) error {
	var err error
	if tweet2.User == "" {
		err = fmt.Errorf("user is required")
	} else if tweet2.Text == "" {
		err = fmt.Errorf("text is required")
	} else if len(tweet2.Text) > 140 {
		err = fmt.Errorf("text exceeds 140 characters")
	} else {
		tweets = append(tweets, tweet2)
	}

	return err
}

//InitializeService aloca espacio
func InitializeService() {
	tweets = make([]*domain.Tweet, 0)

}

//GetTweet getter
func GetTweets() []*domain.Tweet {
	return tweets
}

//CleanTweet limpia el texto
func CleanTweet() {
	tweets = nil
}
