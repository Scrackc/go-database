package main

import (
	"log"

	"github.com/Scrackc/go-pg-database/pkg/invoiceheader"
	"github.com/Scrackc/go-pg-database/pkg/invoiceitem"
	"github.com/Scrackc/go-pg-database/pkg/product"
	"github.com/Scrackc/go-pg-database/storage"
)

func main() {
	storage.NewPGDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

	storageInvoiceHeader := storage.NewPSQLInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)
	if err := serviceInvoiceHeader.Migrate(); err != nil {
		log.Fatalf("invoiceheader.Migrate: %v", err)
	}

	storageInvoiceItem := storage.NewPSQLInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceitem.Migrate: %v", err)
	}

}
