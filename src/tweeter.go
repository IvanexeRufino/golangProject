package main

import (
	"github.com/abiosoft/ishell"
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
			defer c.ShowPrompt(true)
			c.Print("Write your tweet: ")
			tweet := c.ReadLine()
			service.PublishTweet(tweet)
			c.Print("Tweet sent\n")
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "CleanTweet",
		Help: "Clean Tweet",
		Func: func(c *ishell.Context) {
			service.CleanTweet()
			c.Print("Tweet remove\n")
			return
		},
	})

	shell.Run()
}
