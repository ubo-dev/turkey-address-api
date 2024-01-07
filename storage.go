package main

import (
	"database/sql"
	"os"
)

type MysqlStorage struct {
	db *sql.DB
}

type Storage interface {
	GetAllCities() ([]City, error)
	GetCityById(id int) (City, error)
	GetAllDistricts() ([]District, error)
	GetDistrictById(id int) (District, error)
	GetAllNeighboorhoods() ([]Neighboorhood, error)
	GetNeighboorhoodsByDistrictId(id int) ([]Neighboorhood, error)
}

type City struct {
	Id       int
	CityName string
}

type District struct {
	Id           int
	DistrictName string
	CityId       int
}

type Neighboorhood struct {
	Id                int
	NeighboorhoodName string
	DistrictId        int
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

	db, err := sql.Open("mysql", "ubo:Ubovic-mysql-passwd361.:@tcp(172.19.0.1:3306)/turkey-api-db")
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
	data, err := os.ReadFile("data.sql")
	if err != nil {
		return string(data), err
	}
	return string(data), err
}

func (s *MysqlStorage) insertDataSql() error {
	return nil
}
