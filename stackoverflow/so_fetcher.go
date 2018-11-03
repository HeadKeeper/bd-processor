package stackoverflow

import (
	"../db"
	"../util"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// Full URL will contains SITE / API_VERSION / DATA_TYPE ? QUERY

// Parts of URL
const (
	_SO_API_SITE            = "https://api.stackexchange.com"
	_SO_API_VERSION         = "2.2"
	_SO_DATA_TYPE_QUESTIONS = "questions"
	_SO_QUERY_PATTERN       = "key=%s&page=%d&pagesize=%d&fromdate=%d&todate=%d&order=desc&sort=activity&tagged=%s&site=%s"
)

// PATTERN PARAMETERS:
// PAGE (int),
// PAGESIZE (int, Can't be bigger than 100),
// FROMDATE (long),
// TODATE (long),
// TAGGED (string),
// SITE (string)

// Network
const (
	_STATUS_OK         = "200"
	_DB                = "mongo"
	_MONGO_URL         = "localhost:27017"
	_STACKOVERFLOW_KEY = "tAPXnERpI2HiPYtx)BKGvQ(("
)

// Parameters
const (
	_PAGE_SIZE   = 100
	_PAGES_COUNT = 14756
	_FROM_DATE   = 1220227200
	_START_PAGE  = 10501
	_SITE        = "stackoverflow"
	_JAVA_TAG    = "java"
	_C_SHARP_TAG = "c#"
	_RUBY_TAG    = "ruby"
	_C_PLUS_TAG  = "c++"
)

// Util
const (
	_PROCESSING_GROUP_SIZE = 1
	_TAG_CONCATENATOR      = ";"
)

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
	var processingGroup sync.WaitGroup
	processingGroup.Add(_PROCESSING_GROUP_SIZE)
	//go processTags(pattern, _JAVA_TAG, &processingGroup)
	//go processTags(pattern, _C_SHARP_TAG, &processingGroup)
	//go processTags(pattern, _RUBY_TAG, &processingGroup)
	//go processTags(pattern, _C_PLUS_TAG, &processingGroup)
	processingGroup.Wait()

}

func processTags(pattern string, tag string, processingGroup *sync.WaitGroup) {
	for currentPage := _START_PAGE; currentPage <= _PAGES_COUNT; currentPage++ {
		apiUrl := getNextApiURL(pattern, currentPage, tag)
		err := processUrl(apiUrl, tag)
		if err != nil {
			log.Printf("Skipped %s. Cause: %s", apiUrl, err)
		}
	}
	processingGroup.Done()
}

func processUrl(apiUrl, tag string) error {
	log.Println("Start processing ---> " + apiUrl)
	response, err := http.Get(apiUrl)
	if err != nil {
		return err
	}
	if strings.Compare(response.Status, _STATUS_OK) == 0 {
		return errors.New("Response to " + apiUrl + " was crushed. STATUS = " + response.Status)
	}
	defer response.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(response.Body)
	return processQuestionBatch(body, tag)
}

func processQuestionBatch(body []byte, tag string) error {
	data := &db.QuestionBatch{
		Items: []db.Question{},
	}
	json.Unmarshal([]byte(body), &data)
	for _, question := range data.Items {
		question.Id = bson.NewObjectId()
		err := connection.AddRecord(_SO_DATA_TYPE_QUESTIONS+"_"+getTagId(tag), question)
		if err != nil {
			log.Fatalf("Error on Mongo AddRecord. Cause: %s", err)
			waitGroup.Done()
		}
	}
	return nil
}

func getNextApiURL(pattern string, currentPage int, tagsPieces ...string) string {
	currentTimeMillis := util.GetCurrentTimeInMillis()
	tags := strings.Join(tagsPieces, _TAG_CONCATENATOR)
	return fmt.Sprintf(pattern, _STACKOVERFLOW_KEY, currentPage, _PAGE_SIZE, _FROM_DATE, currentTimeMillis, url.QueryEscape(tags), _SITE)
}

func getTagId(tag string) string {
	switch tag {
	case _JAVA_TAG:
		return "java"
	case _RUBY_TAG:
		return "ruby"
	case _C_SHARP_TAG:
		return "c_sharp"
	case _C_PLUS_TAG:
		return "c_plus"
	}
	return "undef"
}
