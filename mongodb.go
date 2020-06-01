package lymon

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// M called as lymon.M aka bson.M wrapper
type M = bson.M
type D = bson.D

// Database export
type Database struct {
	Database   *mongo.Database
	Collection *mongo.Collection
}

// DB mongo database
func (c Global) DB(dbname ...string) Database {
	if len(dbname) > 0 {
		return Database{
			Database: c.Mongo.Database(dbname[0]),
		}
	}

	return Database{
		Database: c.Mongo.Database(c.Database),
	}
}

// C mongo collection
func (db Database) C(collection string) *mongo.Collection {
	return db.Database.Collection(collection)
}
