package twitter

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/RyotaKatoh/oreppoid/config"
)

const (
	twitterAPIURL = ""
)

type TwitterClient struct {
	api *anaconda.TwitterApi
}

func NewTwitterClient() *TwitterClient {

	anaconda.SetConsumerKey(config.TwitterConsumerKey)
	anaconda.SetConsumerSecret(config.TwitterConsumerSecret)
	api := anaconda.NewTwitterApi(config.TwitterAccessToken, config.TwitterAccessTokenSecret)
	cli := TwitterClient{api}
	return &cli

}

func (cli *TwitterClient) Search(q string) []anaconda.Tweet {
	searchResult, _ := cli.api.GetSearch(q, nil)
	return searchResult.Statuses
}

func (cli *TwitterClient) GetTweets(v url.Values) []anaconda.Tweet {
	tweetResult, _ := cli.api.GetUserTimeline(v)
	return tweetResult
}
