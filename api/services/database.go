package services

import (
	"database/sql"
	"fmt"
	"os"
	"tide/api/models"

	_ "github.com/lib/pq"
)

type Database struct {
	Connection *sql.DB
}

func NewDatabase() (*Database, error) {
	d := new(Database)
	con := fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_SERVER"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := sql.Open("posgres", con)
	if err != nil {
		return nil, err
	}
	d.Connection = db
	return d, nil
}

func (d *Database) AddItem(item *models.Item) (*models.Item, error) {
	_, err := d.Connection.Exec("INSERT INTO items (id) VALUES (?)", item.Id)
	if err != nil {
		return nil, err
	}
	return item, nil
}
