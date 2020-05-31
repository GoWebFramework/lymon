package lymon

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Database export
type Database struct {
	Database   *mongo.Database
	Collection *mongo.Collection
}

// DB mongo database
func (c Context) DB(dbname ...interface{}) Database {

	if len(dbname) > 0 {
		switch dbname[0].(type) {
		case string:
			return Database{
				Database: c.Mongo.Database(dbname[0].(string)),
			}
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
