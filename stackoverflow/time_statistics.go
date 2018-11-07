package stackoverflow

import (
	"../db"
	"../util"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strconv"
	"sync"
	"time"
)

func collectPercentageByMonths(statistics *SO_Statistics, connection db.DbConnection, collectionName string, group *sync.WaitGroup) {
	monthsStatistics := make(map[string]float64)
	monthsBorders := util.GetMonthsTimestamps()
	index := 0
	for index < len(monthsBorders)-1 {
		timeInDate := time.Unix(monthsBorders[index], 0)
		timeName := timeInDate.Month().String() + " " + strconv.Itoa(timeInDate.Year())
		lessBorder := monthsBorders[index]
		highBorder := monthsBorders[index+1]
		isAnsweredQuery := bson.M{
			"creation_date": bson.M{
				"$lt": highBorder,
				"$gt": lessBorder,
			},
			"is_answered": true,
		}
		isNotAnsweredQuery := bson.M{
			"creation_date": bson.M{
				"$lt": highBorder,
				"$gt": lessBorder,
			},
			"is_answered": false,
		}
		answered, err := connection.CountByQuery(collectionName, isAnsweredQuery)
		notAnswered, err := connection.CountByQuery(collectionName, isNotAnsweredQuery)
		if err == nil {
			monthsStatistics[timeName] = (float64(answered) / (float64(notAnswered) + float64(answered)))
		} else {
			monthsStatistics[timeName] = -1
		}
		index++
	}
	statistics.AnswersByMonth = monthsStatistics
	fmt.Println(statistics.AnswersByMonth)
	group.Done()
}

func collectPercentageByYears(statistics *SO_Statistics, connection db.DbConnection, collectionName string, group *sync.WaitGroup) {
	yearsStatistics := make(map[string]float64)
	yearsBorders := util.InitializeYears()
	for year, _ := range yearsBorders {
		nextYearInt, err := strconv.Atoi(year)
		nextYear := strconv.Itoa(nextYearInt + 1)
		lessBorder := yearsBorders[year]
		highBorder := yearsBorders[nextYear]
		isAnsweredQuery := bson.M{
			"creation_date": bson.M{
				"$lt": highBorder,
				"$gt": lessBorder,
			},
			"is_answered": true,
		}
		isNotAnsweredQuery := bson.M{
			"creation_date": bson.M{
				"$lt": highBorder,
				"$gt": lessBorder,
			},
			"is_answered": false,
		}
		answered, err := connection.CountByQuery(collectionName, isAnsweredQuery)
		notAnswered, err := connection.CountByQuery(collectionName, isNotAnsweredQuery)
		if err == nil {
			yearsStatistics[year] = (float64(answered) / (float64(notAnswered) + float64(answered)))
		} else {
			yearsStatistics[year] = -1
		}
	}
	statistics.AnswersByYear = yearsStatistics
	group.Done()
}

func collectQuestionsByYears(statistics *SO_Statistics, connection db.DbConnection, collectionName string, group *sync.WaitGroup) {
	yearsStatistics := make(map[string]int)
	yearsBorders := util.InitializeYears()
	for year, _ := range yearsBorders {
		nextYearInt, err := strconv.Atoi(year)
		nextYear := strconv.Itoa(nextYearInt + 1)
		lessBorder := yearsBorders[year]
		highBorder := yearsBorders[nextYear]
		query := bson.M{"creation_date": bson.M{
			"$lt": highBorder,
			"$gt": lessBorder,
		}}
		count, err := connection.CountByQuery(collectionName, query)
		if err != nil {
			log.Fatal(err)
		}
		yearsStatistics[year] = count
	}
	statistics.QuestionsByYear = yearsStatistics
	group.Done()
}
