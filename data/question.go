package data

type Question struct {
	Id    int64  `db:"id" json:"id,string"`
	Value string `db:"value" json:"value"`
}
