package main

import (
	"database/sql"
	"os"
)

type MysqlStorage struct {
	db *sql.DB
}

// cfg := mysql.Config{
// 	User:   os.Getenv("DBUSER"),
// 	Passwd: os.Getenv("DBPASS"),
// 	Net:    "tcp",
// 	Addr:   "127.0.0.1:3306",
// 	DBName: "recordings",
// }
// // Get a database handle.
// var err error
// db, err = sql.Open("mysql", cfg.FormatDSN())
// if err != nil {
// 	log.Fatal(err)
// }

func NewMysqlStorage() (*MysqlStorage, error) {
	connStr := "user=ubo dbname=turkey-api-db password=Ubovic-mysql-passwd361. sslmode=disabled"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &MysqlStorage{
		db: db,
	}, nil
}

func (s *MysqlStorage) Init() error {
	_, err := s.createCityTable(), s.createDistrictTable()
	return err
}

func (s *MysqlStorage) createCityTable() error {
	query := `
		create table if not exists city(
			id primary key,
			city_name varchar(55)
		)
	`
	_, err := s.db.Exec(query)
	return err
}

func (s *MysqlStorage) createDistrictTable() error {
	query := `
		create table if not exists city(
			id primary key,
			district_name varchar(55),
			city_id int 
		)
	`
	_, err := s.db.Exec(query)
	return err
}

func (s *MysqlStorage) ReadFromFile() (string, error) {
	data, err := os.ReadFile("/data.sql")
	if err != nil {
		return string(data), err
	}
	return string(data), err
}

func (s *MysqlStorage) insertDataSql() error {
	return nil
}
