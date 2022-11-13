package data

import "time"

type DistributedSurvey struct {
	Id               int64      `db:"id,omitempty,pk" json:"id,string"`
	CreatedAt        *time.Time `db:"created_at" json:"created_at"`
	Finished         bool       `db:"finished" json:"finished"`
	SurveyQuestionId int64      `db:"survey_question_id" json:"survey_question_id"`
}
