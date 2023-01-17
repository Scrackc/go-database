package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id SERIAL NOT NULL,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id)
	)`
)

// PSQLInvoiceHeader usado par atrabajar con PG y el paquete invoiceheader
type PSQLInvoiceHeader struct {
	db *sql.DB
}

// NewPSQLInvoiceHeader Retorna un nuevo puntero de PSQLInvoiceHeader
func NewPSQLInvoiceHeader(db *sql.DB) *PSQLInvoiceHeader {
	return &PSQLInvoiceHeader{db}
}

// Migrate implemneta la interfaz invoiceheader.Storage
func (p *PSQLInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Migracion de invoiceheader ejecutada correctamente")
	return nil
}
