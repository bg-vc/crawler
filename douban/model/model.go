package model

type DouBanMovie struct {
	Title    string  `db:"title"`
	Subtitle string  `db:"subtitle"`
	Other    string  `db:"order"`
	DescOne  string  `db:"desc_one"`
	DescTwo  string  `db:"desc_two"`
	Score    float64 `db:"score"`
	Comment  int64   `db:"comment"`
	Quote    string  `db:"quote"`
}
