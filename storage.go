package main

import (
	"database/sql"
	"fmt"
	"os"
)

type MysqlStorage struct {
	db *sql.DB
}

type Storage interface {
	GetAllCities() ([]City, error)
	// GetCityById(id int) (City, error)
	// GetAllDistricts() ([]District, error)
	// GetDistrictById(id int) (District, error)
	// GetAllNeighboorhoods() ([]Neighboorhood, error)
	// GetNeighboorhoodsByDistrictId(id int) ([]Neighboorhood, error)
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

var (
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME     = os.Getenv("DB_NAME")
	DB_HOST     = os.Getenv("DB_HOST")
	DB_PORT     = os.Getenv("DB_PORT")
)

func NewMysqlStorage() (*MysqlStorage, error) {

	db, err := sql.Open("mysql", DB_USER+":"+DB_PASSWORD+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME+"?charset=utf8")
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
	_, _, err := s.createCityTable(), s.createDistrictTable(), s.ReadFromFile()
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

func (s *MysqlStorage) ReadFromFile() error {
	data, err := os.ReadFile("data.sql")
	fmt.Println(string(data))
	return err
}

func (s *MysqlStorage) insertDataSql() error {
	return nil
}

// GetAllCities implements Storage.
func (*MysqlStorage) GetAllCities() ([]City, error) {
	panic("unimplemented")
}
