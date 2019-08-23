package model

type Car struct {
	City      string  `db:"city"`
	Title     string  `db:"title"`
	Price     float64 `db:"price"`
	Kilometer float64 `db:"kilometer"`
	Date      int     `db:"date"`
}

type Area struct {
	Name string  `json:"name"`
	City []*City `json:"city"`
}

type City struct {
	Name   string `json:"name"`
	Pinyin string `json:"pinyin"`
}
