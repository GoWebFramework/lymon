package rubi

import (
	"net/http"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Context struct {
	Database *mongo.Database
	Mongo    *mongo.Client
	Redis    *redis.Client

	Path map[string]route
}

type Config struct {
	MongoURI string
	RedisURI string
}

type route struct {
	Handler func(http.ResponseWriter, *http.Request, Context)
	Method  string
}

type M = bson.M
