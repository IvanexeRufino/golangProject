package domain

import "time"
import "fmt"

var id int

//Tweet interface
type Tweet interface {
	String() string
	PrintableTweet() string
	GetUser() string
	GetText() string
	GetID() int
	GetDate() *time.Time
}

//TextTweet estructura
type TextTweet struct {
	ID   int
	User string
	Text string
	Date *time.Time
}

//ImageTweet estructura
type ImageTweet struct {
	ID   int
	User string
	Text string
	URL  string
	Date *time.Time
}

//QuoteTweet estructura
type QuoteTweet struct {
	ID    int
	User  string
	Text  string
	Tweet Tweet
	Date  *time.Time
}

//NewTextTweet crea un tweet
func NewTextTweet(user, text string) *TextTweet {
	date := time.Now()
	id++

	tweet := TextTweet{
		id,
		user,
		text,
		&date,
	}

	return &tweet
}

//PrintableTweet nice tweet
func (t *TextTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", t.User, t.Text)
}

func (t *TextTweet) String() string {
	return t.PrintableTweet()
}

//GetUser getter user
func (t *TextTweet) GetUser() string {
	return t.User
}

//GetText getter text
func (t *TextTweet) GetText() string {
	return t.Text
}

//GetID getter id
func (t *TextTweet) GetID() int {
	return t.ID
}

//GetDate getter date
func (t *TextTweet) GetDate() *time.Time {
	return t.Date
}

//NewImageTweet crea imagen
func NewImageTweet(user, text, url string) *ImageTweet {

	date := time.Now()
	id++

	tweet := ImageTweet{
		id,
		user,
		text,
		url,
		&date,
	}

	return &tweet

}

//PrintableTweet nice tweet
func (t *ImageTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s %s", t.User, t.Text, t.URL)
}

func (t *ImageTweet) String() string {
	return t.PrintableTweet()
}

//GetUser getter user
func (t *ImageTweet) GetUser() string {
	return t.User
}

//GetText getter text
func (t *ImageTweet) GetText() string {
	return t.Text
}

//GetID getter id
func (t *ImageTweet) GetID() int {
	return t.ID
}

//GetDate getter date
func (t *ImageTweet) GetDate() *time.Time {
	return t.Date
}

//NewQuoteTweet crea imagen
func NewQuoteTweet(user, text string, tweetQuoted Tweet) *QuoteTweet {

	date := time.Now()
	id++

	tweet := QuoteTweet{
		id,
		user,
		text,
		tweetQuoted,
		&date,
	}

	return &tweet

}

//PrintableTweet nice tweet
func (t *QuoteTweet) PrintableTweet() string {
	return fmt.Sprintf(`@%s: %s "%s"`, t.User, t.Text, t.Tweet.PrintableTweet())
}

func (t *QuoteTweet) String() string {
	return t.PrintableTweet()
}

//GetUser getter user
func (t *QuoteTweet) GetUser() string {
	return t.User
}

//GetText getter text
func (t *QuoteTweet) GetText() string {
	return t.Text
}

//GetID getter id
func (t *QuoteTweet) GetID() int {
	return t.ID
}

//GetDate getter date
func (t *QuoteTweet) GetDate() *time.Time {
	return t.Date
}
