package lymon

import (
	"net/http"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Context struct {
	Config Config

	Database string
	Mongo    *mongo.Client
	Redis    *redis.Client

	Path map[string]route
}

type Config struct {
	MongoURI string
	RedisURI string

	Listen string
}

type route struct {
	Handler func(http.ResponseWriter, *http.Request, Context)
	Method  string
}

type M = bson.M
