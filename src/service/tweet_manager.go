package service

import (
	"fmt"

	"github.com/golangProject/src/domain"
)

//Tweet es un tweet
var tweets []*domain.Tweet

//PublishTweet qe hace nada
func PublishTweet(tweet2 *domain.Tweet) (int, error) {
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

	return tweet2.ID, err
}

//InitializeService aloca espacio
func InitializeService() {
	tweets = make([]*domain.Tweet, 0)

}

//GetTweets getter
func GetTweets() []*domain.Tweet {
	return tweets
}

//GetLastTweet return last tweet
func GetLastTweet() *domain.Tweet {
	var lastTweet *domain.Tweet
	if len(tweets) != 0 {
		lastTweet = tweets[len(tweets)-1]
	}
	return lastTweet
}

//CleanTweet limpia el texto
func CleanTweet() {
	tweets = nil
}

//GetTweetByID recibe
func GetTweetByID(id int) *domain.Tweet {
	for _, tweet := range tweets {
		if tweet.ID == id {
			return tweet
		}
	}
	return nil
}

//CountTweetsByUser cuenta twees por usuario
func CountTweetsByUser(user string) int {
	contador := 0
	for _, tweet := range tweets {
		if tweet.User == user {
			contador++
		}
	}
	return contador
}

//GetTweetsByUser return tweets by user
func GetTweetsByUser(user string) []*domain.Tweet {
	var tweetsByUser []*domain.Tweet
	tweetsByUser = make([]*domain.Tweet, 0)
	for i, tweet := range tweets {
		if tweet.User == user {
			tweetsByUser = append(tweetsByUser, tweets[i])
		}
	}
	return tweetsByUser
}
