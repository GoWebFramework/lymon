package lymon

import (
	"context"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UseDefaultConfig no comment
func (h *Context) UseDefaultConfig() {
	conf := Config{
		MongoURI: "mongodb://127.0.0.1:27017",
		RedisURI: "redis://127.0.0.1:6379",
		Listen:   "127.0.0.1:8080",
	}
	h.Config = conf

	h.UseConfig(conf)
}

// UseConfig use config from user
func (h *Context) UseConfig(conf Config) {

	h.Config = conf

	if conf.MongoURI != "" {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		m1, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.MongoURI))
		if err != nil {
			panic(err)
		}
		h.Mongo = m1
	}

	if conf.RedisURI != "" {

		url, err := url.Parse(conf.RedisURI) // "redis://username:password@localhost:6379/0"
		if err != nil {
			panic(err)
		}

		cleanup := strings.ReplaceAll(url.Path, "/", "")
		redisdb, err := strconv.Atoi(cleanup)
		if err != nil {
			panic(err)
		}

		// start redis connection
		r1 := redis.NewClient(&redis.Options{
			Addr: url.Host,
			// redis doesn't have a username, so take username position as password
			Password: url.User.Username(),
			DB:       redisdb, // use default DB
		})
		h.Redis = r1
	}
}
