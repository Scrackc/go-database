package storage

import (
	"database/sql"
	"fmt"

	"github.com/Scrackc/go-database/pkg/invoiceheader"
)

const (
	mysqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`
	mysqlCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES (?)`
)

// MySQLInvoiceHeader usado par atrabajar con MYSQL y el paquete invoiceheader
type MySQLInvoiceHeader struct {
	db *sql.DB
}

// NewMySQLInvoiceHeader Retorna un nuevo puntero de MySQLInvoiceHeader
func NewMySQLInvoiceHeader(db *sql.DB) *MySQLInvoiceHeader {
	return &MySQLInvoiceHeader{db}
}

// Migrate implemneta la interfaz invoiceheader.Storage
func (p *MySQLInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(mysqlMigrateInvoiceHeader)
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

// Migrate implemneta la interfaz invoiceheader.Storage
func (p *MySQLInvoiceHeader) CreateTX(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(mysqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(m.Client)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = uint(id)

	return nil
}
