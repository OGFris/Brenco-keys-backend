package database

import (
	"github.com/OGFris/Brenco-keys-backend/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"time"
)

type DB struct {
	*sqlx.DB
}

const queryCreateKeysTable = `CREATE TABLE IF NOT EXISTS "keys" (
  "id" SERIAL UNIQUE PRIMARY KEY,
  "created_at" datetime,
  "updated_at" datetime,
  "name" varchar UNIQUE,
  "key" varchar UNIQUE
)`

type Key struct {
	Id        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Name      string    `db:"name"`
	Key       string    `db:"key"`
}

// New connects to the postgres database and returns the database instance.
func New() (*DB, error) {
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {

		return nil, err
	}

	db.MustExec(queryCreateKeysTable)

	return &DB{db}, nil
}

var queryInsetIntoKeys = `INSERT INTO keys (created_at, updated_at, name, key) VALUES (:created_at, :updated_at, :name, :key)`

// CreateKey takes a name and creates for it a new row in the keys table.
func (db *DB) CreateKey(name string) (string, error) {
	key := utils.GenerateKey()

	_, err := db.NamedExec(queryInsetIntoKeys, &Key{
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Key:       key,
	})

	if err != nil {

		return "", err
	}

	return key, nil
}

var queryRemoveKey = `DELETE FROM keys WHERE name=$1`

// RemoveKey takes a name and remove the row associated to it in the keys table.
func (db *DB) RemoveKey(name string) error {
	_, err := db.Exec(queryRemoveKey, name)

	if err != nil {

		return err
	}

	return nil
}

var querySelectKeys = `SELECT * FROM keys`

// GetKeys returns all the keys from the keys table.
func (db *DB) GetKeys() ([]Key, error) {
	var keys []Key

	err := db.Select(&keys, querySelectKeys)

	if err != nil {

		return nil, err
	}

	return keys, nil
}
