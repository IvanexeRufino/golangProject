package main

import (
	"github.com/abiosoft/ishell"
	"github.com/golangProject/src/domain"
	"github.com/golangProject/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	tm := service.NewTweetManager()

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {
			var tweet *domain.Tweet
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			c.Print("Write your tweet: ")
			text := c.ReadLine()
			tweet = domain.NewTweet(user, text)
			id, err := tm.PublishTweet(tweet)
			if err != nil {
				c.Print("Your tweet has some error, empty text or greater than 140 characters or empty user \n")
			} else {
				c.Print("Tweet sent with id ", id, "\n")
				c.Print("Your tweet is like: ", tweet, "\n")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			i := 0
			tweet := tm.GetTweets()

			if tweet != nil {
				for ; i < len(tweet); i++ {
					c.Println(tweet[i])
				}
			} else {
				c.Println(tweet)
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getLastTweet",
		Help: "Shows last tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			tweet := tm.GetLastTweet()
			if tweet != nil {
				c.Println(tweet)
			} else {
				c.Print("There isnt any tweets published \n")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "cleanTweet",
		Help: "Removes tweet",
		Func: func(c *ishell.Context) {
			tm.CleanTweet()
			c.Print("Tweet remove\n")
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countTweetsByUser",
		Help: "Count tweets sent by user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			c.Print("Enter a username: ")
			user := c.ReadLine()
			count := tm.CountTweetsByUser(user)
			c.Print("User ", user, " has sent ", count, " tweets \n")
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetsByUser",
		Help: "Return tweets sent by user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			c.Print("Enter a username: ")
			user := c.ReadLine()
			tweet := tm.GetTweetsByUser(user)
			if len(tweet) != 0 {
				for i := 0; i < len(tweet); i++ {
					c.Println(tweet[i])
				}
			} else {
				c.Println("User hasnt got any tweets")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "followUser",
		Help: "Enter a user you want to follow",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			c.Print("Enter a username you want to follow: ")
			usertoFollow := c.ReadLine()
			err := tm.Follow(user, usertoFollow)
			if err == nil {
				c.Println("You are now following ", usertoFollow)
			} else {
				c.Println("That user doesnt exist")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTimeline",
		Help: "Shows published tweets that you  might be interested in",
		Func: func(c *ishell.Context) {

			c.Print("Enter your username: ")
			user := c.ReadLine()
			listofTweets := tm.GetTimeline(user)

			if len(listofTweets) != 0 {
				for i := 0; i < len(listofTweets); i++ {
					c.Println(listofTweets[i])
				}
			} else {
				c.Println("Users you follow havent published yet")
			}
			return
		},
	})

	shell.Run()
}
