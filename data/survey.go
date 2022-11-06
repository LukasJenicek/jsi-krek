package data

type Survey struct {
	Id int64 `db:"id" json:"id,string"`
}

type SurveyAnswer struct {
	// TODO: Probably ip address ?
	Who    string `db:"who" json:"who"`
	Answer string `db:"answer" json:"answer"`
}
