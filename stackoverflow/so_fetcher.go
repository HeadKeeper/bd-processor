package stackoverflow

import (
	"sync"
)

// Full URL will contains SITE / API_VERSION / DATA_TYPE ? QUERY

const _SO_SITE = "https://api.stackexchange.com"
const _SO_API_VERSION = "2.2"
const _SO_DATA_TYPE_QUESTIONS = "question"

// PATTERN PARAMETERS:
// PAGE (int),
// PAGESIZE (int, Can't be bigger than 100),
// FROMDATE (long),
// TODATA (long),
// TAGGED (string),
// SITE (string)
const _SO_QUERY_PATTERN = "page=%d&pagesize=%d&fromdate=%d&todate=%d&order=desc&sort=activity&tagged=%s&site=%s"

//const _STACKOVERFLOW_API_PATTERN = "https://api.stackexchange.com/2.2/questions?page=1&pagesize=100&fromdate=1199145600&todate=1540512000&order=desc&sort=activity&tagged=java&site=stackoverflow"

func StartFetching(waitGroup *sync.WaitGroup) {

}
