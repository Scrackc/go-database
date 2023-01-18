package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Scrackc/go-database/pkg/product"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// Driver de storage
type Driver string

// Drivers
const (
	MySQl    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// Crea la conxión con la base de datos
func New(d Driver) {
	switch d {
	case MySQl:
		newMySqlDB()
	case Postgres:
		newPGDB()
	}
}

func newPGDB() {
	// Singleton solo se va a ejeuctar una vez aun que se llame muchas veces
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://postgres:postgresPassword@localhost:5432/go-db?sslmode=disable")
		if err != nil {
			// panic(err)
			log.Fatalf("No se puede conectar a la base de datos: %v", err)
		}
		// defer db.Close()
		if err = db.Ping(); err != nil {
			log.Fatalf("No se pudo comprobar la conexión a la base de datos: %v", err)
			// panic(err)
		}
		fmt.Println("Conectado a PG")
	})
}
func newMySqlDB() {
	// Singleton solo se va a ejeuctar una vez aun que se llame muchas veces
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:postgresPassword@/go-db?parseTime=true")
		if err != nil {
			// panic(err)
			log.Fatalf("No se puede conectar a la base de datos: %v", err)
		}
		// defer db.Close()
		if err = db.Ping(); err != nil {
			log.Fatalf("No se pudo comprobar la conexión a la base de datos: %v", err)
			// panic(err)
		}
		fmt.Println("Conectado a MYSQL")
	})
}

// Pool retorna una unica instancia de DB
func Pool() *sql.DB {
	return db
}

func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

func timeToNULL(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationsNul := sql.NullString{}
	updatedAtNul := sql.NullTime{}
	err := s.Scan(
		&m.ID, &m.Name, &observationsNul, &m.Price, &m.CreatedAt, &updatedAtNul,
	)
	if err != nil {
		return nil, err
	}
	m.Observations = observationsNul.String
	m.UpdatedAt = updatedAtNul.Time

	return m, nil
}

// DAO Product factory of prodcut.storage
func DAOProduct(driver Driver) (product.Storage, error) {
	switch driver {
	case Postgres:
		return newPsqlProduct(db), nil
	case MySQl:
		return newMySQLProduct(db), nil
	default:
		return nil, fmt.Errorf("driver not implemend")
	}
}
