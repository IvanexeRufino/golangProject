package domain

import "time"

//Tweet estructura
type Tweet struct {
	User string
	Text string
	Date *time.Time
}

//NewTweet crea un tweet
func NewTweet(user, text string) *Tweet {
	date := time.Now()

	tweet := Tweet{
		user,
		text,
		&date,
	}

	return &tweet
}
