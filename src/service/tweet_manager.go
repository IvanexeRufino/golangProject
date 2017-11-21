package service

import "github.com/golangProject/src/domain"

//Tweet es un tweet
var tweet *domain.Tweet

//PublishTweet qe hace nada
func PublishTweet(tweet2 *domain.Tweet) {
	tweet = tweet2
}

//GetTweet getter
func GetTweet() *domain.Tweet {
	return tweet
}

//CleanTweet limpia el texto
func CleanTweet() {
	tweet.Text = ""
}
