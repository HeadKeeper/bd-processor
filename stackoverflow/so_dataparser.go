package stackoverflow

import (
	"../db"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"sort"
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

	statistics := new(SO_Statistics)
	statistics.AnsweredPercent = getPercentage(connection, collectionName)
	parseCoTags(connection, collectionName, statistics)
	wg.Done()
	return nil
}

func parseCoTags(connection db.DbConnection, collectionName string, statistics *SO_Statistics) {
	questions, _ := connection.FindByQuery(collectionName, nil)
	var parseGroup sync.WaitGroup
	parseGroup.Add(1)
	//go collectCoTags(statistics, questions.Items, &parseGroup)
	//go collectViewByCoTags(statistics, questions.Items, &parseGroup)
	//go collectAnswersCountByCoTags(statistics, questions.Items, &parseGroup)
	go collectTotalScoreByCoTag(statistics, questions.Items, &parseGroup)
	parseGroup.Wait()
}

func collectTotalScoreByCoTag(statistics *SO_Statistics, questions []db.Question, group *sync.WaitGroup) {
	coTags := make(map[string]int)
	for index, question := range questions {
		fmt.Printf("CO-TAG (Total Score): #%d is started...\n", index)
		for _, tag := range question.Tags {
			coTags[tag] += question.Score
		}
	}
	statistics.TotalScoreByCoTag = getSortedMap(coTags, 16)
	fmt.Println(statistics.TotalScoreByCoTag)
	group.Done()
}

func collectAnswersCountByCoTags(statistics *SO_Statistics, questions []db.Question, group *sync.WaitGroup) {
	coTags := make(map[string]int)
	for index, question := range questions {
		fmt.Printf("CO-TAG (Answers Count): #%d is started...\n", index)
		for _, tag := range question.Tags {
			coTags[tag] += question.AnswerCount
		}
	}
	statistics.AnswersCountByCoTag = getSortedMap(coTags, 16)
	fmt.Println(statistics.AnswersCountByCoTag)
	group.Done()
}

func collectViewByCoTags(statistics *SO_Statistics, questions []db.Question, group *sync.WaitGroup) {
	coTags := make(map[string]int)
	for index, question := range questions {
		fmt.Printf("CO-TAG (View Count): #%d is started...\n", index)
		for _, tag := range question.Tags {
			coTags[tag] += question.ViewCount
		}
	}
	statistics.ViewCountByCoTag = getSortedMap(coTags, 16)
	fmt.Println(statistics.ViewCountByCoTag)
	group.Done()
}

func collectCoTags(statistics *SO_Statistics, questions []db.Question, group *sync.WaitGroup) {
	coTags := make(map[string]int)
	for index, question := range questions {
		fmt.Printf("CO-TAG: #%d is started...\n", index)
		for _, tag := range question.Tags {
			coTags[tag]++
		}
	}
	statistics.CoTags = getSortedMap(coTags, 16) // Top 15, excluded tag herself
	fmt.Println(statistics.CoTags)
	group.Done()
}

// Get top N topics in map
func getSortedMap(unsorted map[string]int, topCount int) map[string]int {
	tmp := map[int][]string{}
	result := make(map[string]int)
	var a []int
	for k, v := range unsorted {
		tmp[v] = append(tmp[v], k)
	}
	for k := range tmp {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	border := 0
	for _, k := range a {
		for _, s := range tmp[k] {
			if border < topCount {
				result[s] = k
			}
			border++
		}
	}
	return result
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
