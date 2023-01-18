package main

import (
	"log"

	"github.com/Scrackc/go-database/pkg/invoice"
	"github.com/Scrackc/go-database/pkg/invoiceheader"
	"github.com/Scrackc/go-database/pkg/invoiceitem"
	"github.com/Scrackc/go-database/storage"
)

func main() {

	storage.NewMySqlDB()

	storageHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	storageItems := storage.NewMySQLInvoiceItem(storage.Pool())
	storageInvoice := storage.NewMySQLInvoice(storage.Pool(), storageHeader, storageItems)
	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Eduardo",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductId: 1},
			&invoiceitem.Model{ProductId: 3},
		},
	}
	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}

}
