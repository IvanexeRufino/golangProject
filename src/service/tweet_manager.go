package service

import (
	"fmt"

	"github.com/golangProject/src/domain"
)

//Tweet es un tweet
var tweets map[string][]*domain.Tweet
var lastTweet *domain.Tweet
var followers map[string][]string

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
		tweets[tweet2.User] = append(tweets[tweet2.User], tweet2)
	}

	lastTweet = tweet2

	return tweet2.ID, err
}

//InitializeService aloca espacio
func InitializeService() {
	tweets = make(map[string][]*domain.Tweet)
	lastTweet = nil
	followers = make(map[string][]string)

}

//GetTweets getter
func GetTweets() []*domain.Tweet {
	var listOfTweets []*domain.Tweet
	for _, listTweet := range tweets {
		listOfTweets = append(listOfTweets, listTweet...)
	}
	return listOfTweets
}

//GetLastTweet return last tweet
func GetLastTweet() *domain.Tweet {
	return lastTweet
}

//CleanTweet limpia el texto
func CleanTweet() {
	tweets = nil
	InitializeService()
}

//GetTweetByID recibe
func GetTweetByID(id int) *domain.Tweet {
	for _, tweet := range GetTweets() {
		if tweet.ID == id {
			return tweet
		}
	}
	return nil
}

//CountTweetsByUser cuenta twees por usuario
func CountTweetsByUser(user string) int {
	return len(tweets[user])
}

//GetTweetsByUser return tweets by user
func GetTweetsByUser(user string) []*domain.Tweet {
	return tweets[user]
}

//Follow follows
func Follow(follower, user string) error {
	var err error
	_, ok := tweets[user]
	if ok {
		followers[follower] = append(followers[follower], user)
	} else {
		err = fmt.Errorf("user doesnt exist")
	}
	return err
}

//GetTimeline returns followers published tweets
func GetTimeline(user string) []*domain.Tweet {
	followedUsers := followers[user]
	var listOfTweets []*domain.Tweet
	for _, users := range followedUsers {
		listOfTweets = append(listOfTweets, tweets[users]...)
	}
	return listOfTweets
}
