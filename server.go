package main

import (
	"net/url"

	"github.com/RyotaKatoh/oreppoid/app/mongo"
	"github.com/RyotaKatoh/oreppoid/app/twitter"
)

func main() {

	twitterClient := twitter.NewTwitterClient()
	v := url.Values{
		"count":       {"200"},
		"include_rts": {"false"},
	}

	tweets := twitterClient.GetTweets(v)

	mongoClient := &mongo.MongoClient{}
	mongoClient.SaveTweet(tweets)
}
