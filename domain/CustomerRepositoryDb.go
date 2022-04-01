package domain

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	// registering database driver

	_ "github.com/lib/pq"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Print("Error while querying customer table" + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)

		if err != nil {
			log.Print("Errir while scanning customer" + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	var psqlInfo string

	if os.Getenv("pgHost") == "" {
		fmt.Println("Cant get env fron os.Getenv, get env from config.json")
		filename := "./config.json"

		jsonFile, err := os.Open(filename)
		if err != nil {
			fmt.Printf("failed to open json file: %s, error: %v", filename, err)
			panic(err)
		}
		defer jsonFile.Close()

		jsonData, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			fmt.Printf("failed to read json file, error: %v", err)
			panic(err)
		}

		data := Data{}
		if err := json.Unmarshal(jsonData, &data); err != nil {
			fmt.Printf("failed to unmarshal json file, error: %v", err)
			panic(err)
		}

		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			data.Host, data.Port, data.User, data.Password, data.DBName)
	} else {
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("pgHost"), os.Getenv("pgPort"), os.Getenv("pgUser"), os.Getenv("pgPassword"), os.Getenv("pgDbName"))
	}

	client, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxIdleTime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxOpenConns(10)

	err = MakeMigration(client)
	if err != nil {
		log.Panic(err)
	}

	return CustomerRepositoryDb{client}
}

func MakeMigration(db *sql.DB) error {
	b, err := ioutil.ReadFile("./models.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}

type Data struct {
	Host     string `json:"pgHost"`
	Port     string `json:"pgPort"`
	User     string `json:"pgUser"`
	Password string `json:"pgPassword"`
	DBName   string `json:"pgDBname"`
}
