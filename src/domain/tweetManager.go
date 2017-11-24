package domain

import (
	"fmt"
	"strings"
)

//TweetManager struct
type TweetManager struct {
	Tweets    map[string][]*Tweet
	LastTweet *Tweet
	Users     []*User
	TTs       map[string]int
}

//
func (tm *TweetManager) iterateTweetForTTs(tweet *Tweet) {

	listOfWords := strings.Fields(tweet.Text)
	for _, word := range listOfWords {
		tm.TTs[word] = tm.TTs[word] + 1
	}

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
		_, ok := tm.Tweets[tweet2.User]
		if !ok {
			usuarioNuevo := NewUser(tweet2.User)
			tm.Users = append(tm.Users, usuarioNuevo)
		}
		tm.Tweets[tweet2.User] = append(tm.Tweets[tweet2.User], tweet2)
		tm.LastTweet = tweet2
		tm.iterateTweetForTTs(tweet2)
	}

	return tweet2.ID, err
}

//InitializeService aloca espacio
func (tm *TweetManager) InitializeService() {
	tm.Tweets = make(map[string][]*Tweet)
	tm.LastTweet = nil
	tm.Users = make([]*User, 0)
	tm.TTs = make(map[string]int)
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

//GetUserByName get user by name
func (tm *TweetManager) GetUserByName(name string) *User {
	var userFound *User
	for _, user := range tm.Users {
		if user.Name == name {
			userFound = user
		}
	}

	return userFound
}

//Follow follows
func (tm *TweetManager) Follow(follower, user string) error {
	var err error
	userFound := tm.GetUserByName(user)
	userFollower := tm.GetUserByName(follower)

	if userFound != nil && userFollower != nil {
		userFollower.Followeds = append(userFollower.Followeds, user)
	} else {
		err = fmt.Errorf("User doesnt exist")
	}
	return err
}

//GetTimeline returns followers published tweets
func (tm *TweetManager) GetTimeline(user string) []*Tweet {
	followedUsers := tm.GetUserByName(user).Followeds
	var listOfTweets []*Tweet
	for _, users := range followedUsers {
		listOfTweets = append(listOfTweets, tm.Tweets[users]...)
	}
	return listOfTweets
}

//SendMessage send a message
func (tm *TweetManager) SendMessage(from, to, message string) error {
	var err error

	if tm.GetUserByName(from) != nil && tm.GetUserByName(to) != nil {
		dm := NewDirectMessage(from, message)
		user := tm.GetUserByName(to)
		user.DirectMessages = append(user.DirectMessages, dm)
	} else {
		err = fmt.Errorf("User doesnt exist")
	}

	return err
}

//GetAllDirectMessages return direct messages from user
func (tm *TweetManager) GetAllDirectMessages(name string) []*DirectMessage {
	user := tm.GetUserByName(name)
	for _, dm := range user.DirectMessages {
		dm.Readed = true
	}
	return user.DirectMessages
}

//GetUnreadedMessages return direct messages unreadeds
func (tm *TweetManager) GetUnreadedMessages(name string) []*DirectMessage {

	var unreaded []*DirectMessage
	unreaded = make([]*DirectMessage, 0)
	user := tm.GetUserByName(name)
	for _, dm := range user.DirectMessages {
		if !dm.Readed {
			unreaded = append(unreaded, dm)
			dm.Readed = true
		}
	}

	return unreaded

}

//GetTrendingTopics get more used words
func (tm *TweetManager) GetTrendingTopics() []string {

	var listOfCounters []int
	listOfCounters = make([]int, 0)
	mayor := 0
	segundoMayor := 0
	var trendings []string
	trendings = make([]string, 2)

	for k, v := range tm.TTs {
		listOfCounters = append(listOfCounters, v)
		if v > mayor {
			mayor = v
			trendings[0] = k
		} else if v > segundoMayor {
			segundoMayor = v
			trendings[1] = k
		}
	}

	return trendings
}

//Retweetear the msg
func (tm *TweetManager) Retweetear(name string, id int) {

	tweet := tm.GetTweetByID(id)
	tm.Tweets[name] = append(tm.Tweets[name], tweet)
}

//Fav add to favourites
func (tm *TweetManager) Fav(name string, id int) {

	tweet := tm.GetTweetByID(id)
	user := tm.GetUserByName(name)
	user.FavouriteTweets = append(user.FavouriteTweets, tweet)
}

//GetTweetsFav list of favourite tweets
func (tm *TweetManager) GetTweetsFav(name string) []*Tweet {

	user := tm.GetUserByName(name)
	return user.FavouriteTweets
}
