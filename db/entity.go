package db

import "gopkg.in/mgo.v2/bson"

type QuestionBatch struct {
	Items []Question `json:"items"`
}

type Question struct {
	Id               bson.ObjectId `bson:"_id" json:"id"`
	Tags             []string      `bson:"tags" json:"tags"`
	IsAnswered       bool          `bson:"is_answered" json:"is_answered"`
	ViewCount        int           `bson:"view_count" json:"view_count"`
	AcceptedAnswerId int           `bson:"accepted_answer_id" json:"accepted_answer_id"`
	AnswerCount      int           `bson:"answer_count" json:"answer_count"`
	Score            int           `bson:"score" json:"score"`
	LastActivityDate uint64        `bson:"last_activity_date" json:"last_activity_date"`
	CreationDate     uint64        `bson:"creation_date" json:"creation_date"`
	LastEditDate     uint64        `bson:"last_edit_date" json:"last_edit_date"`
	QuestionId       int           `bson:"question_id" json:"question_id"`
	Link             string        `bson:"link" json:"link"`
	Title            string        `bson:"title" json:"title"`
}
