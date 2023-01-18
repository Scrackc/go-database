package storage

import (
	"database/sql"
	"fmt"

	"github.com/Scrackc/go-database/pkg/invoiceitem"
)

const (
	mysqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers(id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`
	mysqlCreateInvoiceItem = `INSERT INTO invoice_items(invoice_header_id, product_id) VALUES (?, ?)`
)

// MySQLInvoiceItem usado par atrabajar con MYSQL y el paquete invoiceitem
type MySQLInvoiceItem struct {
	db *sql.DB
}

// NewMySQLInvoiceItem Retorna un nuevo puntero de MySQLInvoiceItem
func NewMySQLInvoiceItem(db *sql.DB) *MySQLInvoiceItem {
	return &MySQLInvoiceItem{db}
}

// Migrate implemneta la interfaz invoiceitem.Storage
func (p *MySQLInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(mysqlMigrateInvoiceItem)
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

// CreateTX implemneta la interfaz invoiceitem.Storage
func (p *MySQLInvoiceItem) CreateTX(tx *sql.Tx, headerID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(mysqlCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range ms {

		result, err := stmt.Exec(headerID, item.ProductId)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		item.ID = uint(id)

	}
	return nil
}
