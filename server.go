package main

import (
	"net/url"
	"os"

	"github.com/RyotaKatoh/oreppoid/app/mongo"
	"github.com/RyotaKatoh/oreppoid/app/twitter"
	"github.com/RyotaKatoh/oreppoid/config"
	"github.com/RyotaKatoh/oreppoid/lib/server-helper/logging"
	"github.com/Sirupsen/logrus"
	"github.com/justinas/alice"
)

func getEnvVar(key, defaultVal string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		val = defaultVal
	}

	return val
}

func main() {

	logrus.SetLevel(logrus.DebugLevel)
	serverVersion := config.ServerVersion

	mws := []alice.Constructor{
		logging.NewHandler,
		//		func(handler http.Handler) http.Handler{return filteres.RenderSetupHandler},
	}

	//	appHandler := alice.New(mws...).Then(app.)

	twitterClient := twitter.NewTwitterClient()
	v := url.Values{
		"count":       {"200"},
		"include_rts": {"false"},
	}

	tweets := twitterClient.GetTweets(v)

	mongoClient := &mongo.MongoClient{}
	mongoClient.SaveTweet(tweets)
}
