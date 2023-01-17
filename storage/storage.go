package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPGDB() {
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

// Pool retorna una unica instancia de DB
func Pool() *sql.DB {
	return db
}
