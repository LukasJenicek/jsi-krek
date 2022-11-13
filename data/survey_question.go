package data

type SurveyQuestion struct {
	Id       int64  `db:"id" json:"id,string"`
	Question string `db:"question" json:"question"`
	SurveyId int64  `db:"survey_id" json:"survey_id"`
}
