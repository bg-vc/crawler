package dao

import (
	"crawler/qczj/model"
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

func BatchAddCar(cars []*model.Car) error {
	sql := fmt.Sprintf(`insert into qczj_car(city, title, price, kilometer, date) `)
	sql += fmt.Sprintf(`values(?, ?, ?, ?, ?) `)
	tx, err := db.Beginx()
	if err != nil {
		log.Printf("db.Beginx error:%v\n", err)
	}
	for _, car := range cars {
		if _, err := tx.Exec(sql, car.City, car.Title, car.Price, car.Kilometer, car.Date); err != nil {
			log.Printf("tx.Exec error:%v\n", err)
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
