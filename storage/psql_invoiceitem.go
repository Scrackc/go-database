package storage

import (
	"database/sql"
	"fmt"

	"github.com/Scrackc/go-database/pkg/invoiceitem"
)

const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id SERIAL NOT NULL,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_id_pk PRIMARY KEY (id),
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers(id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`
	psqlCreateInvoiceItem = `INSERT INTO invoice_items(invoice_header_id, product_id) VALUES ($1, $2) RETURNING id, created_at`
)

// PSQLInvoiceItem usado par atrabajar con PG y el paquete invoiceheader
type PSQLInvoiceItem struct {
	db *sql.DB
}

// NewPSQLInvoiceItem Retorna un nuevo puntero de PSQLInvoiceItem
func NewPSQLInvoiceItem(db *sql.DB) *PSQLInvoiceItem {
	return &PSQLInvoiceItem{db}
}

// Migrate implemneta la interfaz invoiceheader.Storage
func (p *PSQLInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Migracion de invoiceitem ejecutada correctamente")
	return nil
}

// CreateTX implemneta la interfaz invoiceheader.Storage
func (p *PSQLInvoiceItem) CreateTX(tx *sql.Tx, headerID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range ms {
		err = stmt.QueryRow(headerID, item.ProductId).Scan(
			&item.ID,
			&item.CreatedAt,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
