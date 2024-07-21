package db

import (
	"Anvarjon-33/Nuxt_Go.port/classic"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofor-little/env"
)

var db *sql.DB

func init() {
	err := env.Load(".env")
	if err != nil {
		panic(err)
	}
}

func GetUser(wild string) ([]classic.Customers, error) {
	var res []classic.Customers
	db, err := sql.Open("mysql", env.Get("DB_URL", ""))
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(`select customerName, contactLastName, contactFirstName from customers where customerName LIKE ?`, "%"+wild+"%")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var t classic.Customers
		if err := rows.Scan(&t.ContactLastName, &t.CustomerName, &t.ContactFirstName); err != nil {
			panic(err)
		}
		res = append(res, t)
	}
	return res, err
}
func GetUser_() ([]classic.Customers, error) {
	var res []classic.Customers
	db, err := sql.Open("mysql", env.Get("DB_URL", ""))
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("select customerName, contactLastName, contactFirstName from customers;")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var t classic.Customers
		if err := rows.Scan(&t.ContactLastName, &t.CustomerName, &t.ContactFirstName); err != nil {
			panic(err)
		}
		res = append(res, t)
	}
	return res, err
}
