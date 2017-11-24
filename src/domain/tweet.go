package domain

import "time"
import "fmt"

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

//PrintableTweet nice tweet
func (t *Tweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", t.User, t.Text)
}

func (t *Tweet) String() string {
	return t.PrintableTweet()
}
