package service

import (
	"fmt"
)

//Plugin interface
type Plugin interface {
	action(string, *TweetManager)
}

//FacebookPlugin implementation
type FacebookPlugin struct {
	urlAPI string
}

//GooglePlugin implementation
type GooglePlugin struct {
	urlAPI string
}

//TweeterPlugin implementation
type TweeterPlugin struct {
	urlAPI string
}

func (fp *FacebookPlugin) action(user string, tm *TweetManager) {
	fmt.Println("Your tweet has been plubished in facebook")
}

func (fp *GooglePlugin) action(user string, tm *TweetManager) {
	fmt.Println("Your tweet has been plubished in google")
}

func (fp *TweeterPlugin) action(user string, tm *TweetManager) {
	cuantity := tm.CountTweetsByUser(user)
	fmt.Printf("You have pulbished %d \n", cuantity)
}

func newFacebookPlugin() *FacebookPlugin {
	plug := FacebookPlugin{
		"facebook.com",
	}
	return &plug
}

func newGooglePlugin() *GooglePlugin {
	plug := GooglePlugin{
		"google.com",
	}
	return &plug
}

func newTweeterPlugin() *TweeterPlugin {
	plug := TweeterPlugin{
		"tweeter.com",
	}
	return &plug
}

//NewPlugin by ids
func NewPlugin(id int) Plugin {

	var plug Plugin

	switch id {
	case 1:
		plug = newFacebookPlugin()
	case 2:
		plug = newGooglePlugin()
	case 3:
		plug = newTweeterPlugin()
	}

	return plug
}
