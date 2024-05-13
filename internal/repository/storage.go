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
	GetCityById(id int) (model.City, error)
	GetAllDistricts() ([]model.District, error)
	GetDistrictById(id int) (model.District, error)
	GetDistrictByCityId(id int) ([]model.District, error)
	GetAllNeighbourhoods() ([]model.Neighbourhood, error)
	GetNeighbourhoodsByDistrictId(id int) ([]model.Neighbourhood, error)
	GetNeighbourhoodsByDistrictName(districtName string) ([]model.Neighbourhood, error)
	GetNeighbourhoodsByZipCode(zipCode string) ([]model.Neighbourhood, error)
}

func NewMysqlStorage() (*MysqlRepository, error) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = godotenv.Load(filepath.Join(pwd, ".env"))
	// err = godotenv.Load(filepath.Join(pwd, "../.env"))
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

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)
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
	_, _, _, err := s.createCityTable(), s.createDistrictTable(), s.createNeighbourhoodTable(), s.ReadFromFile()
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

func (s *MysqlRepository) createNeighbourhoodTable() error {
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
	// data, err := os.ReadFile("../data.sql")
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

func (s *MysqlRepository) GetCityById(id int) (model.City, error) {
	query := `
		select * from city where id = ?
	`
	row := s.db.QueryRow(query, id)

	var city model.City
	if err := row.Scan(&city.Id, &city.CityName); err != nil {
		return model.City{}, err
	}
	return city, nil
}

func (s *MysqlRepository) GetAllDistricts() ([]model.District, error) {
	query := `
		select * from district
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var districts []model.District
	for rows.Next() {
		var district model.District
		if err := rows.Scan(&district.Id, &district.DistrictName, &district.DistrictName); err != nil {
			return nil, err
		}
		districts = append(districts, district)
	}
	return districts, nil
}

func (s *MysqlRepository) GetDistrictById(id int) (model.District, error) {
	query := `
		select * from district where id = ?
	`

	row := s.db.QueryRow(query, id)
	var district model.District
	if err := row.Scan(&district.Id, &district.DistrictName, &district.CityId); err != nil {
		return model.District{}, err
	}
	return district, nil
}

func (s *MysqlRepository) GetDistrictByCityId(id int) ([]model.District, error) {
	query := `
		select * from district where city_id = ?
	`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var districts []model.District
	for rows.Next() {
		var district model.District
		if err := rows.Scan(&district.Id, &district.DistrictName, &district.CityId); err != nil {
			return nil, err
		}
		districts = append(districts, district)
	}
	return districts, nil
}

func (s *MysqlRepository) GetAllNeighbourhoods() ([]model.Neighbourhood, error) {
	query := `
		select * from neighbourhood
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var neighbourhoods []model.Neighbourhood
	for rows.Next() {
		var neighbourhood model.Neighbourhood
		if err := rows.Scan(&neighbourhood.Id, &neighbourhood.NeighbourhoodName, &neighbourhood.DistrictId); err != nil {
			return nil, err
		}
		neighbourhoods = append(neighbourhoods, neighbourhood)
	}
	return neighbourhoods, nil
}

func (s *MysqlRepository) GetNeighboorhoodsByDistrictId(id int) ([]model.Neighbourhood, error) {
	query := `
		select * from neighbourhood where district_id = ?
	`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var neighbourhoods []model.Neighbourhood
	for rows.Next() {
		var neighbourhood model.Neighbourhood
		if err := rows.Scan(&neighbourhood.Id, &neighbourhood.NeighbourhoodName, &neighbourhood.DistrictId); err != nil {
			return nil, err
		}
		neighbourhoods = append(neighbourhoods, neighbourhood)
	}
	return neighbourhoods, nil
}

func (s *MysqlRepository) GetNeighboorhoodsByDistrictName(
	districtName string,
) ([]model.Neighbourhood, error) {
	query := `
		select * from neighbourhood where district_id = (select id from district where district_name = ?)
	`
	rows, err := s.db.Query(query, districtName)
	if err != nil {
		return nil, err
	}

	var neighbourhoods []model.Neighbourhood
	for rows.Next() {
		var neighbourhood model.Neighbourhood
		if err := rows.Scan(&neighbourhood.Id, &neighbourhood.NeighbourhoodName, &neighbourhood.DistrictId); err != nil {
			return nil, err
		}
		neighbourhoods = append(neighbourhoods, neighbourhood)
	}
	return neighbourhoods, nil
}

func (s *MysqlRepository) GetNeighboorhoodsByZipCode(
	zipCode string,
) ([]model.Neighbourhood, error) {
	query := `
		select * from neighbourhood where postal_code = ?
	`
	rows, err := s.db.Query(query, zipCode)
	if err != nil {
		return nil, err
	}

	var neighbourhoods []model.Neighbourhood
	for rows.Next() {
		var neighbourhood model.Neighbourhood
		if err := rows.Scan(&neighbourhood.Id, &neighbourhood.NeighbourhoodName, &neighbourhood.DistrictId); err != nil {
			return nil, err
		}
		neighbourhoods = append(neighbourhoods, neighbourhood)
	}
	return neighbourhoods, nil
}
