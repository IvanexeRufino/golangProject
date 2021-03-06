package service

import (
	"fmt"
	"strings"

	"github.com/golangProject/src/domain"
)

//TweetManager struct
type TweetManager struct {
	LastTweet domain.Tweet
	Users     []*domain.User
	TTs       map[string]int
	Plugins   []Plugin
	Writer    *ChannelTweetWriter
}

//NewTweetManager constructor
func NewTweetManager() *TweetManager {
	tweetMan := TweetManager{}
	tweetMan.InitializeService()

	return &tweetMan
}

func (tm *TweetManager) iterateTweetForTTs(tweet domain.Tweet) {

	listOfWords := strings.Fields(tweet.GetText())
	for _, word := range listOfWords {
		tm.TTs[word] = tm.TTs[word] + 1
	}

}

//AddPlugin adds
func (tm *TweetManager) AddPlugin(plug Plugin) {
	tm.Plugins = append(tm.Plugins, plug)
}

//ExecuteActions obersevers
func (tm *TweetManager) ExecuteActions(user string) {
	for _, plug := range tm.Plugins {
		plug.action(user, tm)
	}
}

//AddUser adds one
func (tm *TweetManager) AddUser(user string) *domain.User {
	usuarioNuevo := domain.NewUser(user)
	tm.Users = append(tm.Users, usuarioNuevo)
	return usuarioNuevo
}

//PublishTweet qe hace nada
func (tm *TweetManager) PublishTweet(tweet2 domain.Tweet) (int, error) {
	var err error
	if tweet2.GetUser() == "" {
		err = fmt.Errorf("user is required")
	} else if tweet2.GetText() == "" {
		err = fmt.Errorf("text is required")
	} else if len(tweet2.GetText()) > 140 {
		err = fmt.Errorf("text exceeds 140 characters")
	} else {
		tweetsToWrite := make(chan domain.Tweet)
		quit := make(chan bool)

		go tm.Writer.WriteTweet(tweetsToWrite, quit)
		tweetsToWrite <- tweet2
		close(tweetsToWrite)

		<-quit
		tm.LastTweet = tweet2
		tm.iterateTweetForTTs(tweet2)
	}

	tm.ExecuteActions(tweet2.GetUser())

	return tweet2.GetID(), err
}

//InitializeService aloca espacio
func (tm *TweetManager) InitializeService() {
	tm.LastTweet = nil
	tm.Users = make([]*domain.User, 0)
	tm.TTs = make(map[string]int)
	tm.Plugins = make([]Plugin, 0)
	tm.Writer = NewChannelTweetWriter(NewMemoryTweetWriter())
}

//GetTweets getter
func (tm *TweetManager) GetTweets() []domain.Tweet {
	var listOfTweets []domain.Tweet
	for _, listTweet := range tm.Writer.GetTweets() {
		listOfTweets = append(listOfTweets, listTweet...)
	}
	return listOfTweets
}

//GetLastTweet return last tweet
func (tm *TweetManager) GetLastTweet() domain.Tweet {
	return tm.LastTweet
}

//CleanTweet limpia el texto
func (tm *TweetManager) CleanTweet() {
	tm.InitializeService()
}

//GetTweetByID recibe
func (tm *TweetManager) GetTweetByID(id int) domain.Tweet {
	for _, tweet := range tm.GetTweets() {
		if tweet.GetID() == id {
			return tweet
		}
	}
	return nil
}

//CountTweetsByUser cuenta twees por usuario
func (tm *TweetManager) CountTweetsByUser(user string) int {
	return len(tm.Writer.GetTweets()[user])
}

//GetTweetsByUser return tweets by user
func (tm *TweetManager) GetTweetsByUser(user string) []domain.Tweet {
	return tm.Writer.GetTweets()[user]
}

//GetUserByName get user by name
func (tm *TweetManager) GetUserByName(name string) *domain.User {
	var userFound *domain.User
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
func (tm *TweetManager) GetTimeline(user string) []domain.Tweet {
	followedUsers := tm.GetUserByName(user).Followeds
	var listOfTweets []domain.Tweet
	for _, users := range followedUsers {
		listOfTweets = append(listOfTweets, tm.Writer.GetTweets()[users]...)
	}
	return listOfTweets
}

//SendMessage send a message
func (tm *TweetManager) SendMessage(from, to, message string) error {
	var err error

	if tm.GetUserByName(from) != nil && tm.GetUserByName(to) != nil {
		dm := domain.NewDirectMessage(from, message)
		user := tm.GetUserByName(to)
		user.DirectMessages = append(user.DirectMessages, dm)
	} else {
		err = fmt.Errorf("User doesnt exist")
	}

	return err
}

//GetAllDirectMessages return direct messages from user
func (tm *TweetManager) GetAllDirectMessages(name string) []*domain.DirectMessage {
	user := tm.GetUserByName(name)
	for _, dm := range user.DirectMessages {
		dm.Readed = true
	}
	return user.DirectMessages
}

//GetUnreadedMessages return direct messages unreadeds
func (tm *TweetManager) GetUnreadedMessages(name string) []*domain.DirectMessage {

	var unreaded []*domain.DirectMessage
	unreaded = make([]*domain.DirectMessage, 0)
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
	tm.Writer.GetTweets()[name] = append(tm.Writer.GetTweets()[name], tweet)
}

//Fav add to favourites
func (tm *TweetManager) Fav(name string, id int) {

	tweet := tm.GetTweetByID(id)
	user := tm.GetUserByName(name)
	user.FavouriteTweets = append(user.FavouriteTweets, tweet)
}

//GetTweetsFav list of favourite tweets
func (tm *TweetManager) GetTweetsFav(name string) []domain.Tweet {

	user := tm.GetUserByName(name)
	return user.FavouriteTweets
}
