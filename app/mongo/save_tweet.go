package mongo

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/RyotaKatoh/oreppoid/config"
	"github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

type MongoClient struct {
}

type Tweet struct {
	Text      string `bson: "text"`
	Name      string `bson: "user"`
	UserId    string `bson: "user_id"`
	CreatedAt string `bson: "created_at"`
}

func (cli *MongoClient) SaveTweet(tweets []anaconda.Tweet) {
	session, err := mgo.Dial(config.MongoDBServerHost)
	if err != nil {
		logrus.Errorln(err)
	}
	defer session.Close()
	db := session.DB("oreppoid")
	col := db.C("tweet")

	for _, tweet := range tweets {
		logrus.Infoln("saving...", tweet.Text)
		insertTweet := Tweet{
			Text:      tweet.Text,
			Name:      tweet.User.Name,
			UserId:    tweet.User.IdStr,
			CreatedAt: tweet.CreatedAt,
		}
		err := col.Insert(insertTweet)
		if err != nil {
			logrus.Fatalln(err)
		}
	}

	logrus.Infoln("tweets are saved in mongo")

}
