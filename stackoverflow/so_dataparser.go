package stackoverflow

import (
	"../db"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

var collectionName string

// Collectable information:
// 1) Percent of answered questions

func ParseData(tag string, wg *sync.WaitGroup) error {
	connection, err := db.GetConnection(_DB, _MONGO_URL, "", "")
	if err != nil {
		wg.Done()
		return err
	}
	defer connection.Close()

	collectionName = getCollectionName(tag)
	percentOfAnswered := getPercentage(connection, collectionName)

	fmt.Println(percentOfAnswered)
	wg.Done()
	return nil
}

func getPercentage(connection db.DbConnection, collectionName string) float64 {
	isAnsweredQuery := bson.M{"is_answered": true}
	isNotAnsweredQuery := bson.M{"is_answered": false}
	answered, err := connection.CountByQuery(collectionName, isAnsweredQuery)
	notAnswered, err := connection.CountByQuery(collectionName, isNotAnsweredQuery)
	if err == nil {
		return (float64(answered) / (float64(notAnswered) + float64(answered)))
	} else {
		return -1
	}
}

func getCollectionName(tag string) string {
	return _SO_DATA_TYPE_QUESTIONS + "_" + getTagId(tag)
}
