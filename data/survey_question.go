package data

import "github.com/upper/db/v4"

type SurveyQuestion struct {
	Id       int64  `db:"id,omitempty,pk" json:"id,string"`
	Question string `db:"question" json:"question"`
	SurveyId int64  `db:"survey_id" json:"survey_id"`
}

func (*SurveyQuestion) Store(sess db.Session) db.Store {
	return GetSurveyQuestionStore(sess)
}

func GetSurveyQuestionStore(sess db.Session) *SurveyQuestionStore {
	return &SurveyQuestionStore{sess.Collection("survey_questions")}
}

type SurveyQuestionStore struct {
	db.Collection
}
