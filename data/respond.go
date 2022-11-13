package data

import "time"

type Respond struct {
	Id        int64      `db:"id" json:"id,string"`
	SurveyId  int64      `db:"survey_id" json:"survey_id"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	Who       string     `db:"who" json:"who"` // TODO: Probably ip address ?
	Answer    string     `db:"answer" json:"answer"`
}
