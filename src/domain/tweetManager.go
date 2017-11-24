package domain

import (
	"fmt"
)

//TweetManager struct
type TweetManager struct {
	Tweets    map[string][]*Tweet
	LastTweet *Tweet
	Followers map[string][]string
}

//PublishTweet qe hace nada
func (tm *TweetManager) PublishTweet(tweet2 *Tweet) (int, error) {
	var err error
	if tweet2.User == "" {
		err = fmt.Errorf("user is required")
	} else if tweet2.Text == "" {
		err = fmt.Errorf("text is required")
	} else if len(tweet2.Text) > 140 {
		err = fmt.Errorf("text exceeds 140 characters")
	} else {
		tm.Tweets[tweet2.User] = append(tm.Tweets[tweet2.User], tweet2)
	}

	tm.LastTweet = tweet2

	return tweet2.ID, err
}

//InitializeService aloca espacio
func (tm *TweetManager) InitializeService() {
	tm.Tweets = make(map[string][]*Tweet)
	tm.LastTweet = nil
	tm.Followers = make(map[string][]string)

}

//GetTweets getter
func (tm *TweetManager) GetTweets() []*Tweet {
	var listOfTweets []*Tweet
	for _, listTweet := range tm.Tweets {
		listOfTweets = append(listOfTweets, listTweet...)
	}
	return listOfTweets
}

//GetLastTweet return last tweet
func (tm *TweetManager) GetLastTweet() *Tweet {
	return tm.LastTweet
}

//CleanTweet limpia el texto
func (tm *TweetManager) CleanTweet() {
	tm.Tweets = nil
	tm.InitializeService()
}

//GetTweetByID recibe
func (tm *TweetManager) GetTweetByID(id int) *Tweet {
	for _, tweet := range tm.GetTweets() {
		if tweet.ID == id {
			return tweet
		}
	}
	return nil
}

//CountTweetsByUser cuenta twees por usuario
func (tm *TweetManager) CountTweetsByUser(user string) int {
	return len(tm.Tweets[user])
}

//GetTweetsByUser return tweets by user
func (tm *TweetManager) GetTweetsByUser(user string) []*Tweet {
	return tm.Tweets[user]
}

//Follow follows
func (tm *TweetManager) Follow(follower, user string) error {
	var err error
	_, ok := tm.Tweets[user]
	if ok {
		tm.Followers[follower] = append(tm.Followers[follower], user)
	} else {
		err = fmt.Errorf("user doesnt exist")
	}
	return err
}

//GetTimeline returns followers published tweets
func (tm *TweetManager) GetTimeline(user string) []*Tweet {
	followedUsers := tm.Followers[user]
	var listOfTweets []*Tweet
	for _, users := range followedUsers {
		listOfTweets = append(listOfTweets, tm.Tweets[users]...)
	}
	return listOfTweets
}
