package lymon

import (
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

// Global hold Configuration, Database, Path and Handler data
type Global struct {
	Config Config

	Database string
	Mongo    *mongo.Client
	Redis    *redis.Client

	// MiddlewareHandler store middleware handler with ordered-array
	MiddlewareHandler []handler

	// Path store main route handler based on route+method that used as key
	Path map[string]route

	// StatusCodeHandler store handler based on http status code
	StatusCodeHandler map[int]handler
}
