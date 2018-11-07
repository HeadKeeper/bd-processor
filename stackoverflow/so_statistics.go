package stackoverflow

type SO_Statistics struct {
	Tag                 string             `json:"tag"`
	AnsweredPercent     float64            `json:"answered_percent"`
	CoTags              map[string]int     `json:"co_tags"`
	QuestionsByYear     map[string]int     `json:"questions_by_year"`
	QuestionsByMonth    map[string]int     `json:"questions_by_month"`
	AnswersByMonth      map[string]float64 `json:"answers_by_month"`
	AnswersByYear       map[string]float64 `json:"answers_by_year"`
	ViewCountByCoTag    map[string]int     `json:"view_count_by_co_tag"`
	AnswersCountByCoTag map[string]int     `json:"answers_count_by_co_tag"`
	TotalScoreByCoTag   map[string]int     `json:"total_score_by_co_tag"`
}
