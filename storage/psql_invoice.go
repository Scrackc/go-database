package storage

import (
	"database/sql"
	"fmt"

	"github.com/Scrackc/go-database/pkg/invoice"
	"github.com/Scrackc/go-database/pkg/invoiceheader"
	"github.com/Scrackc/go-database/pkg/invoiceitem"
)

// PsqlInvoice usado para trabajar con pg - invoice
type PsqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

// NewPsqlInvoice retorna un puntero de PsqlInvoice
func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *PsqlInvoice {
	return &PsqlInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

// Create implemta la interface invoice.storage
func (p *PsqlInvoice) Create(m *invoice.Model) error {
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
