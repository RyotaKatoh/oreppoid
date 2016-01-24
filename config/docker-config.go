package config

type dockerConfig struct {
}

func (c dockerConfig) MongoURI() string {
	return "mongodb://0.0.0.0:27017/oreppoid"
}

func (c dockerConfig) MongoDatabase() string {
	return "oreppoid"
}
