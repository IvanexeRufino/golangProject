package domain

import "time"

var id int

//Tweet estructura
type Tweet struct {
	ID   int
	User string
	Text string
	Date *time.Time
}

//NewTweet crea un tweet
func NewTweet(user, text string) *Tweet {
	date := time.Now()
	id++

	tweet := Tweet{
		id,
		user,
		text,
		&date,
	}

	return &tweet
}
