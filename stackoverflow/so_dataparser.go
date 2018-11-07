package stackoverflow

import (
	"../db"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

var collectionName string

func ParseData(tag string, wg *sync.WaitGroup) error {
	connection, err := db.GetConnection(_DB, _MONGO_URL, "", "")
	if err != nil {
		wg.Done()
		return err
	}
	defer connection.Close()

	collectionName = getCollectionName(tag)

	statistics := new(SO_Statistics)
	statistics.AnsweredPercent = getPercentage(connection, collectionName)
	//parseCoTags(connection, collectionName, statistics)
	parseTimeStatistics(connection, collectionName, statistics)
	wg.Done()
	return nil
}

func parseTimeStatistics(connection db.DbConnection, collectionName string, statistics *SO_Statistics) {
	var parseGroup sync.WaitGroup
	parseGroup.Add(1)
	go collectPercentageByYears(statistics, connection, collectionName, &parseGroup)
	//go collectQuestionsByYears(statistics, connection, collectionName, &parseGroup)
	parseGroup.Wait()
}

func parseCoTags(connection db.DbConnection, collectionName string, statistics *SO_Statistics) {
	questions, _ := connection.FindByQuery(collectionName, nil)
	var parseGroup sync.WaitGroup
	parseGroup.Add(4)
	go collectCoTags(statistics, questions.Items, &parseGroup)
	go collectViewByCoTags(statistics, questions.Items, &parseGroup)
	go collectAnswersCountByCoTags(statistics, questions.Items, &parseGroup)
	go collectTotalScoreByCoTag(statistics, questions.Items, &parseGroup)
	parseGroup.Wait()
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
