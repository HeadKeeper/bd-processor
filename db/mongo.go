package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const _DEFAULT_USERNAME = "undefined"
const _DEFAULT_PASSWORD = "undefined"
const _DEFAULT_DATABASE = "bd-processor"

type MongoConnection struct {
	DbConnection
	session  *mgo.Session
	url      string
	username string
	password string
	database string
}

func (connection *MongoConnection) connectToDb(url, username, password string) (DbConnection, error) {
	session, err := mgo.Dial(url)
	connection.session = session
	connection.url = url
	connection.password = _DEFAULT_PASSWORD
	connection.username = _DEFAULT_USERNAME
	connection.database = _DEFAULT_DATABASE
	return connection, err
}

func (connection *MongoConnection) AddRecord(target string, record interface{}) error {
	return connection.session.DB(_DEFAULT_DATABASE).C(target).Insert(record)
}

func (connection *MongoConnection) ReplaceRecord(target string, record interface{}, id bson.ObjectId) error {
	return nil
}

func (connection *MongoConnection) DeleteRecord(target string, id bson.ObjectId) error {
	return nil
}

func (connection *MongoConnection) Close() {
	if connection.session != nil {
		log.Printf("Connection %s was closed.", connection.url)
	}
}

func (connection *MongoConnection) FindByQuery(collection string, query bson.M) (*QuestionBatch, error) {
	result := new(QuestionBatch)
	err := connection.session.DB(_DEFAULT_DATABASE).C(collection).Find(query).All(&result.Items)
	return result, err
}

func (connection *MongoConnection) CountByQuery(collection string, query bson.M) (int, error) {
	result, err := connection.session.DB(_DEFAULT_DATABASE).C(collection).Find(query).Count()
	return result, err
}
