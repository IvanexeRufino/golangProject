package main

import (
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/golangProject/src/domain"
	"github.com/golangProject/src/service"
)

func login(shelly *ishell.Shell, tm *service.TweetManager) *domain.User {
	shelly.Print("Enter your username: ")
	shelly.SetPrompt("Tweeter >> ")
	usuario := shelly.ReadLine()
	user := tm.GetUserByName(usuario)

	if user != nil {
		shelly.Print("You are already registered in the system, welcome back \n")
	} else {
		user = tm.AddUser(usuario)
		shelly.Print("You ahave been registered in our system, welcome \n")
	}

	return user
}

func main() {

	var user *domain.User

	shell := ishell.New()
	tm := service.NewTweetManager()

	user = login(shell, tm)

	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "logout",
		Help: "Exit session",
		Func: func(c *ishell.Context) {
			user = login(shell, tm)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			var tweet domain.Tweet
			c.Print("Write your tweet: ")
			text := c.ReadLine()
			tweet = domain.NewTextTweet(user.Name, text)
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
			defer c.ShowPrompt(true)
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
			c.Print("Enter a username you want to follow: ")
			usertoFollow := c.ReadLine()
			err := tm.Follow(user.Name, usertoFollow)
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

			defer c.ShowPrompt(true)
			listofTweets := tm.GetTimeline(user.Name)

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

	shell.AddCmd(&ishell.Cmd{
		Name: "showUsers",
		Help: "Shows registered users",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			listofUsers := tm.Users

			if len(listofUsers) != 0 {
				for i := 0; i < len(listofUsers); i++ {
					c.Println(listofUsers[i])
				}
			} else {
				c.Println("There isnt any user registered yet")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "sendMessage",
		Help: "Send a message from a user to another",
		Func: func(c *ishell.Context) {

			c.Print("Enter the reciever username: ")
			to := c.ReadLine()
			c.Print("Write your message: ")
			message := c.ReadLine()

			tm.SendMessage(user.Name, to, message)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showAllDirectMessage",
		Help: "Get all messages from a user",
		Func: func(c *ishell.Context) {

			listOfDMs := tm.GetAllDirectMessages(user.Name)

			if len(listOfDMs) != 0 {
				for i := 0; i < len(listOfDMs); i++ {
					c.Println(listOfDMs[i])
				}
			} else {
				c.Println("There isnt message registered yet")
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showUnreadedMessage",
		Help: "Get all unreaded messages from a user",
		Func: func(c *ishell.Context) {

			listOfDMs := tm.GetUnreadedMessages(user.Name)

			if len(listOfDMs) != 0 {
				for i := 0; i < len(listOfDMs); i++ {
					c.Println(listOfDMs[i])
				}
			} else {
				c.Println("There isnt any unreaded message registered yet")
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTrendingTopics",
		Help: "Retrieve most your works",
		Func: func(c *ishell.Context) {

			c.Print("the TrendingTopic  are the following: ")

			listOfDMs := tm.GetTrendingTopics()

			c.Println(listOfDMs)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "retweetear",
		Help: "retwitea",
		Func: func(c *ishell.Context) {
			c.Print("ingrese el id del tweet: ")
			id := c.ReadLine()
			idNum, _ := strconv.Atoi(id)
			tm.Retweetear(user.Name, idNum)
			c.Println("Ha retweeteado exitosamente")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "favouriteTweet",
		Help: "Adds to favourite",
		Func: func(c *ishell.Context) {
			c.Print("ingrese el id del tweet: ")
			id := c.ReadLine()
			idNum, _ := strconv.Atoi(id)
			tm.Fav(user.Name, idNum)
			c.Println("Tweet has been added to your favourit list")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getFavTweets",
		Help: "Show all favourite tweets",
		Func: func(c *ishell.Context) {
			ft := tm.GetTweetsFav(user.Name)

			if len(ft) != 0 {
				for i := 0; i < len(ft); i++ {
					c.Println(ft[i])
				}
			} else {
				c.Println("This user hasnt got any fav tweets")
			}
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "addPlugin",
		Help: "Adds a plugin",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Ingrese el id del plugin: ")
			id := c.ReadLine()

			ident, ok := strconv.Atoi(id)

			if (ident == 1 || ident == 2 || ident == 3) && ok == nil {

				plug := service.NewPlugin(ident)

				tm.AddPlugin(plug)

				c.Println("Plugin has been adde correctly")

			} else {
				c.Println("That id doesnt exist")
			}

			return
		},
	})

	shell.Run()
}
