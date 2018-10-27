package db

import "gopkg.in/mgo.v2/bson"

type Question struct {
	Id               bson.ObjectId `bson:"_id"`
	Tags             []string      `bson:"tags"`
	IsAnswered       bool          `bson:"is_answered"`
	ViewCount        int           `bson:"view_count"`
	AcceptedAnswerId int           `bson:"accepted_answer_id"`
	AnswerCount      int           `bson:"answer_count"`
	Score            int           `bson:"score"`
	LastActivityDate int64         `bson:"last_activity_date"`
	CreationDate     int64         `bson:"creation_date"`
	LastEditDate     int64         `bson:"last_edit_date"`
	QuestionId       int           `bson:"question_id"`
	Link             string        `bson:"link"`
	Title            string        `bson:"title"`
}
