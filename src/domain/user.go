package domain

//User estructura
type User struct {
	Name            string
	Followeds       []string
	DirectMessages  []*DirectMessage
	FavouriteTweets []*Tweet
}

//NewUser crea un tweet
func NewUser(name string) *User {

	user := User{
		name,
		make([]string, 0),
		make([]*DirectMessage, 0),
		make([]*Tweet, 0),
	}

	return &user
}
