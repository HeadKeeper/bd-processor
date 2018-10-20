package db

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

const _DEFAULT_USERNAME = "undefined"
const _DEFAULT_PASSWORD = "undefined"

type MongoConnection struct {
	DbConnection
	session *mgo.Session
	url string
	username string
	password string
}

func (connection *MongoConnection) connectToDb(url, username, password string) (DbConnection, error) {
	session, err := mgo.Dial(url)
	connection.session = session
	connection.url = url
	connection.password = _DEFAULT_PASSWORD
	connection.username = _DEFAULT_USERNAME
	return connection, err
}
