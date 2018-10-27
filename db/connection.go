package db

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

/*
	Defined connection types:
	- Mongo
*/

const MONGO = "mongo"

type DbConnection interface {
	connectToDb(url, username, password string) (DbConnection, error)
	AddRecord(target string, record interface{}) error
	ReplaceRecord(target string, record interface{}, id bson.ObjectId) error
	DeleteRecord(target string, id bson.ObjectId) error
	Close()
}

func GetConnection(connectionType, url, username, password string) (DbConnection, error) {
	var connection DbConnection
	switch connectionType {
	case MONGO:
		connection = new(MongoConnection)
	default:
		return nil, errors.New("Undefined connection type.")
	}
	connection.connectToDb(url, username, password)
	return connection, nil
}
