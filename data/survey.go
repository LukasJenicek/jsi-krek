package data

type Survey struct {
	Id int64 `db:"id,omitempty,pk" json:"id,string"`
}
