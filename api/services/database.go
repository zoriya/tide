package services

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/zoriya/tide/api/models"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Database struct {
	Connection *sql.DB
}

func NewDatabase() (*Database, error) {
	d := new(Database)
	con := fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_SERVER"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	d.Connection = db
	return d, nil
}

func (db *Database) Migrate() error {
	driver, err := postgres.WithInstance(db.Connection, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}
	m.Up()
	return nil
}

func (d *Database) AddItem(item *models.Item) (*models.Item, error) {
	_, err := d.Connection.Exec(
		"INSERT INTO items (id, name, path, size, files) VALUES (?, ?, ?, ?, ?)",
		item.Id,
		item.Name,
		item.Path,
		item.Size,
		item.Files,
	)
	if err != nil {
		return nil, err
	}
	return item, nil
}
