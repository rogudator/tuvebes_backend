package entity

type Tuvebe struct {
	Date     string `json:"date" db:"date"`
	TimeFrom string `json:"time_from" db:"time_from"`
	TimeTo   string `json:"time_to" db:"time_to"`
	Title    string `json:"title" db:"title"`
}
