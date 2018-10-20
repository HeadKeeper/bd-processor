package db

import "errors"

/*
	Defined connection types:
	- Mongo
*/

const MONGO = "mongo"

type DbConnection interface {
	connectToDb(url, username, password string) (DbConnection, error)
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
