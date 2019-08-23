package dao

import (
	"crawler/douban/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var (
	db *sqlx.DB
)

func InitDB() {
	driverName := "mysql"
	dataSourceName := fmt.Sprintf(`%v:%v@tcp(%v:%v)/%v`, "root", "root", "127.0.0.1", "3306", "crawler")
	dbConn, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		log.Printf("InitDB error:%v\n", err)
		panic("InitDB error")
	}
	db = dbConn
}

func AddMovie(item *model.DouBanMovie) error {
	sql := fmt.Sprintf(`insert into douban_movie(title, subtitle, other, desc_one, desc_two, score, comment, quote) `)
	sql += fmt.Sprintf(`values(?, ?, ?, ?, ?, ?, ?, ?) `)
	fmt.Printf("item:#%v\n", item)

	if _, err := db.Exec(sql, item.Title, item.Subtitle, item.Other,
		item.DescOne, item.DescTwo, item.Score, item.Comment, item.Quote); err != nil {
		log.Printf("AddMovie error:%v\n", err)
		return err
	}
	return nil
}
