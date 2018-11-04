package stackoverflow

type SO_Statistics struct {
	Tag                 string
	AnsweredPercent     float64
	CoTags              map[string]int
	QuestionsByYear     map[string]int
	AnswersByYear       map[string]float64
	ViewCountByCoTag    map[string]int
	AnswersCountByCoTag map[string]int
	TotalScoreByCoTag   map[string]int
}
