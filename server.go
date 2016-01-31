package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RyotaKatoh/oreppoid/app"
	"github.com/RyotaKatoh/oreppoid/app/filters"
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

	mws := []alice.Constructor{
		logging.NewHandler,
		func(handler http.Handler) http.Handler { return filters.RenderSetupHandler(handler) },
	}

	appHandler := alice.New(mws...).Then(app.BuildRouter())
	http.Handle("/", appHandler)

	port := getEnvVar("PORT", config.DefaultServerPort)

	addr := fmt.Sprintf(":%s", port)
	logrus.Infof("Server listening on port %s", port)
	http.ListenAndServe(addr, nil)

	/*
		twitterClient := twitter.NewTwitterClient()
		v := url.Values{
			"count":       {"200"},
			"include_rts": {"false"},
		}

		tweets := twitterClient.GetTweets(v)

		mongoClient := &mongo.MongoClient{}
		mongoClient.SaveTweet(tweets)

	*/
}
