package storage

import (
	"database/sql"
	"fmt"

	"github.com/Scrackc/go-database/pkg/invoice"
	"github.com/Scrackc/go-database/pkg/invoiceheader"
	"github.com/Scrackc/go-database/pkg/invoiceitem"
)

// MySQLInvoice usado para trabajar con mysql - invoice
type MySQLInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

// MySQLsqlInvoice retorna un puntero de MySQLInvoice
func NewMySQLInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *MySQLInvoice {
	return &MySQLInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

// Create implemta la interface invoice.storage
func (p *MySQLInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if err := p.storageHeader.CreateTX(tx, m.Header); err != nil {
		tx.Rollback()
		return err
	}
	fmt.Printf("Factura creada con id: %d\n", m.Header.ID)
	if err = p.storageItems.CreateTX(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return err
	}
	fmt.Printf("Items añadidos a la factura: %d\n", len(m.Items))

	return tx.Commit()
}
