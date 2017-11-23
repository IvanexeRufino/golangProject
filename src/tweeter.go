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
			id, err := service.PublishTweet(tweet)
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
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			i := 0
			tweet := service.GetTweets()

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
		Name: "cleanTweet",
		Help: "Removes tweet",
		Func: func(c *ishell.Context) {
			service.CleanTweet()
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
			count := service.CountTweetsByUser(user)
			c.Print("User ", user, " has sent ", count, " tweets \n")
			return
		},
	})

	shell.Run()
}
