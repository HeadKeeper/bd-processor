package stackoverflow

import (
	"../db"
	"fmt"
	"sort"
	"sync"
)

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
