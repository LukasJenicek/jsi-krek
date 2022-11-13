package data

type SurveyAnswer struct {
	Id               int64  `db:"id,omitempty,pk" json:"id,string"`
	Answer           string `db:"answer" json:"answer"`
	SurveyQuestionId int64  `db:"survey_question_id" json:"survey_question_id"`
}
