package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id)
	)`
)

// PSQLProduct usado par atrabajar con PG y el paquete product
type PSQLProduct struct {
	db *sql.DB
}

// NewPsqlProduct Retorna un nuevo puntero de PSQLProduct
func NewPsqlProduct(db *sql.DB) *PSQLProduct {
	return &PSQLProduct{db}
}

// Migrate implemneta la interfaz product.Storage
func (p *PSQLProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Migracion de producto ejecutada correctamente")
	return nil
}
