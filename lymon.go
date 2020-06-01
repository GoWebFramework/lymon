package lymon

import (
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

// Context hold single lymon instance
type Global struct {
	Config Config

	Database string
	Mongo    *mongo.Client
	Redis    *redis.Client

	Path              map[string]route
	MiddlewareHandler []handler
	StatusCodeHandler map[int]handler
}
