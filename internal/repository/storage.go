package repository

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/ubo-dev/turkey-address-api/internal/model"
)

type MysqlRepository struct {
	db *sql.DB
}

type Repository interface {
	GetAllCities() ([]model.City, error)
	// GetCityById(id int) (City, error)
	// GetAllDistricts() ([]District, error)
	// GetDistrictById(id int) (District, error)
	// GetAllNeighboorhoods() ([]Neighboorhood, error)
	// GetNeighboorhoodsByDistrictId(id int) ([]Neighboorhood, error)
}

//	cfg := mysql.Config{
//		User:   os.Getenv("DBUSER"),
//		Passwd: os.Getenv("DBPASS"),
//		Net:    "tcp",
//		Addr:   "127.0.0.1:3306",
//		DBName: "recordings",
//	}
//
// // Get a database handle.
// var err error
// db, err = sql.Open("mysql", cfg.FormatDSN())
//
//	if err != nil {
//		log.Fatal(err)
//	}
func NewMysqlStorage() (*MysqlRepository, error) {

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = godotenv.Load(filepath.Join(pwd, ".env"))
	if err != nil {
		return nil, err
	}

	var (
		DB_USER     = os.Getenv("MYSQL_USER")
		DB_PASSWORD = os.Getenv("MYSQL_PASSWORD")
		DB_NAME     = os.Getenv("MYSQL_DATABASE")
		DB_PORT     = os.Getenv("MYSQL_PORT")
		DB_HOST     = os.Getenv("MYSQL_HOST")
	)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	fmt.Println(connectionString)
	db, err := sql.Open("mysql", connectionString)
	db.SetMaxIdleConns(0)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &MysqlRepository{
		db: db,
	}, nil
}

func (s *MysqlRepository) Init() error {
	_, _, _, err := s.createCityTable(), s.createDistrictTable(), s.createNeighbourhoodable(), s.ReadFromFile()
	return err
}

func (s *MysqlRepository) createCityTable() error {
	query := `
		create table if not exists city(
			id int primary key,
			city_name varchar(55)
		)
	`
	_, err := s.db.Exec(query)
	fmt.Print("city table created")
	return err
}

func (s *MysqlRepository) createDistrictTable() error {
	query := `
		create table if not exists district(
			id int primary key,
			district_name varchar(55),
			city_id int,
			foreign key (city_id) references city(id)
		)
	`
	_, err := s.db.Exec(query)
	fmt.Print("district table created")
	return err
}

func (s *MysqlRepository) createNeighbourhoodable() error {
	query := `
		create table if not exists neighbourhood(
			id int primary key,
			postal_code int unique,
			neighbourhood_name varchar(55),
			district_id int,
			foreign key (district_id) references district(id)
		)
	`
	_, err := s.db.Exec(query)
	fmt.Print("neighbourhood table created")
	return err
}

func (s *MysqlRepository) ReadFromFile() error {
	data, err := os.ReadFile("data.sql")
	if err != nil {
		fmt.Println(err)
	}
	_, err = s.db.Exec(string(data))

	fmt.Println(err)
	return err
}

func (s *MysqlRepository) GetAllCities() ([]model.City, error) {
	query := `
		select * from city
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var cities []model.City
	for rows.Next() {
		var city model.City
		if err := rows.Scan(&city.Id, &city.CityName); err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}
	return cities, nil
}
