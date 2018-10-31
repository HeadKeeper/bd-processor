package stackoverflow

import (
	"../db"
	"../util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Full URL will contains SITE / API_VERSION / DATA_TYPE ? QUERY

const _SO_API_SITE = "https://api.stackexchange.com"
const _SO_API_VERSION = "2.2"
const _SO_DATA_TYPE_QUESTIONS = "question"

// PATTERN PARAMETERS:
// PAGE (int),
// PAGESIZE (int, Can't be bigger than 100),
// FROMDATE (long),
// TODATE (long),
// TAGGED (string),
// SITE (string)
const _SO_QUERY_PATTERN = "page=%d&pagesize=%d&fromdate=%d&todate=%d&order=desc&sort=activity&tagged=%s&site=%s"

//const _STACKOVERFLOW_API_PATTERN = "https://api.stackexchange.com/2.2/questions?page=1&pagesize=100&fromdate=1199145600&todate=1540512000&order=desc&sort=activity&tagged=java&site=stackoverflow"

const _PAGE_SIZE = 100
const _PAGES_COUNT = 100
const _FROM_DATE = 1220227200
const _START_PAGE = 1

const _SITE = "stackoverflow"
const _JAVA_TAG = "java"

const _DB = "mongo"
const _MONGO_URL = "localhost:27017"

var waitGroup *sync.WaitGroup
var connection db.DbConnection

func StartFetching(wg *sync.WaitGroup) {
	waitGroup = wg
	connectToDb()
	defer connection.Close()
	fetchQuestions()
	waitGroup.Done()
}

func connectToDb() {
	var err error
	log.Printf("Connecting to '%s'", _DB)
	connection, err = db.GetConnection(_DB, _MONGO_URL, "", "")
	if err != nil {
		log.Fatalf("Can't create connection to '%s' on address %s. Cause: %s", _DB, _MONGO_URL, err)
		waitGroup.Done()
	}
	log.Printf("Connection to '%s' was completly!", _DB)
}

func fetchQuestions() {
	pattern := util.JoinURL(
		_SO_API_SITE,
		_SO_API_VERSION,
		_SO_DATA_TYPE_QUESTIONS,
		_SO_QUERY_PATTERN,
	)
	for currentPage := _START_PAGE; currentPage < _PAGES_COUNT; currentPage++ {
		apiUrl := getNextApiURL(pattern, currentPage)
		err := processUrl(apiUrl)
		if err != nil {
			log.Printf("Skipped %s. Cause: %s", apiUrl, err)
		}
	}
}

func processUrl(apiUrl string) error {
	response, err := http.Get(apiUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	var body []byte
	response.Body.Read(body)
	return processQuestionBatch(body)
}

func processQuestionBatch(body []byte) error {
	data := &db.QuestionBatch{
		Items: []db.Question{},
	}
	json.Unmarshal([]byte(body), &data)
	for _, question := range data.Items {
		question.Id = bson.NewObjectId()
		err := connection.AddRecord(_SO_DATA_TYPE_QUESTIONS, question)
		if err != nil {
			log.Fatalf("Error on Mongo AddRecord. Cause: %s", err)
			waitGroup.Done()
		}
	}
	return nil
}

func getNextApiURL(pattern string, currentPage int) string {
	currentTimeMillis := util.GetCurrentTimeInMillis()
	tags := util.ConcatTags(_JAVA_TAG)
	return fmt.Sprintf(pattern, currentPage, _PAGE_SIZE, _FROM_DATE, currentTimeMillis, tags)
}

// ----------------
//var tags []string
//tags = append(tags, "aaa", "bbb", "ccc")
//a := db.Question{
//	Id: bson.NewObjectId(),
//	Tags: tags,
//	IsAnswered: true,
//	ViewCount: 1,
//	AcceptedAnswerId: 123,
//	AnswerCount: 500,
//	Score: 333,
//	LastActivityDate: 1540644304,
//	CreationDate: 1496221604,
//	LastEditDate: 1496221605,
//	QuestionId: 300,
//	Link: "test link",
//	Title: "Title",
//
//}
//err := connection.AddRecord(_SO_DATA_TYPE_QUESTIONS, a)
//if (err != nil) {
//	log.Fatal(err)
//	waitGroup.Done()
//}
// ----------

/*
batch := `{
	"items": [
	{
	"tags": [
	"java",
	"apache-spark",
	"apache-kafka",
	"apache-spark-sql",
	"spark-structured-streaming"
	],
	"owner": {
	"reputation": 1385,
	"user_id": 1870400,
	"user_type": "registered",
	"accept_rate": 42,
	"profile_image": "https://www.gravatar.com/avatar/86a511ab70d0d94c77746a4d27c66fdf?s=128&d=identicon&r=PG",
	"display_name": "user1870400",
	"link": "https://stackoverflow.com/users/1870400/user1870400"
	},
	"is_answered": true,
	"view_count": 774,
	"accepted_answer_id": 44281045,
	"answer_count": 1,
	"score": 1,
	"last_activity_date": 1540644304,
	"creation_date": 1496221604,
	"last_edit_date": 1540644304,
	"question_id": 44280360,
	"link": "https://stackoverflow.com/questions/44280360/how-to-convert-datasetrow-to-dataset-of-json-messages-to-write-to-kafka",
	"title": "How to Convert DataSet&lt;Row&gt; to DataSet of JSON messages to write to Kafka?"
	},
	{
	"tags": [
	"c#",
	"java",
	"c++",
	"memory",
	"garbage-collection"
	],
	"owner": {
	"reputation": 387,
	"user_id": 343841,
	"user_type": "registered",
	"accept_rate": 100,
	"profile_image": "https://www.gravatar.com/avatar/6e9a9f1fad9da5f9a79a99d67eb4fcdc?s=128&d=identicon&r=PG",
	"display_name": "EmbeddedProg",
	"link": "https://stackoverflow.com/users/343841/embeddedprog"
	},
	"is_answered": true,
	"view_count": 1756,
	"accepted_answer_id": 2983171,
	"answer_count": 5,
	"score": 19,
	"last_activity_date": 1540644279,
	"creation_date": 1275776267,
	"question_id": 2982325,
	"link": "https://stackoverflow.com/questions/2982325/quantifying-the-performance-of-garbage-collection-vs-explicit-memory-management",
	"title": "Quantifying the Performance of Garbage Collection vs. Explicit Memory Management"
	}
	]}`
	data := &db.QuestionBatch{
		Items: []db.Question{},
	}
	json.Unmarshal([]byte(batch), &data)
	for _, question := range data.Items {
		question.Id = bson.NewObjectId()
		err := connection.AddRecord(_SO_DATA_TYPE_QUESTIONS, question)
		if (err != nil) {
			log.Fatalf("Error on Mongo AddRecord. Cause: %s", err)
			waitGroup.Done()
		}
	}
*/
