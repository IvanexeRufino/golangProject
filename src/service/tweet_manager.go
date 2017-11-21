package service

//Tweet es un tweet
var tweet string

//PublishTweet qe hace nada
func PublishTweet(ttweet string) {
	tweet = ttweet
}

//GetTweet getter
func GetTweet() string {
	return tweet
}
func CleanTweet() {
	tweet = ""

}
