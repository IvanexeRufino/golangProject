package rest

import (
	"net/http"

	"github.com/golangProject/src/domain"

	"github.com/gin-gonic/gin"
	"github.com/golangProject/src/service"
)

//GinTweet DTO
type GinTweet struct {
	User string
	Text string
}

//GinServer struct
type GinServer struct {
	tweetManager *service.TweetManager
}

//NewGinServer constructor
func NewGinServer(tweetManager *service.TweetManager) *GinServer {
	return &GinServer{tweetManager}
}

//StartGinServer starts
func (server *GinServer) StartGinServer() {
	router := gin.Default()

	router.GET("listTweets", server.listTweets)
	router.GET("listTweets/:user", server.getTweetsByUser)
	router.POST("publishTweet", server.publishTweet)

	go router.Run()
}

func (server *GinServer) listTweets(c *gin.Context) {

	c.JSON(http.StatusOK, server.tweetManager.GetTweets())
}

func (server *GinServer) getTweetsByUser(c *gin.Context) {

	user := c.Param("user")
	c.JSON(http.StatusOK, server.tweetManager.GetTweetsByUser(user))
}

func (server *GinServer) publishTweet(c *gin.Context) {

	var tweetdata GinTweet
	c.Bind(&tweetdata)

	tweetToPublish := domain.NewTextTweet(tweetdata.User, tweetdata.Text)

	id, err := server.tweetManager.PublishTweet(tweetToPublish)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error publishing tweet")
	} else {
		c.JSON(http.StatusOK, struct{ ID int }{id})
	}

}
