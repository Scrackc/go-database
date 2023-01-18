package main

import (
	"log"

	"github.com/Scrackc/go-pg-database/pkg/invoiceheader"
	"github.com/Scrackc/go-pg-database/pkg/invoiceitem"
	"github.com/Scrackc/go-pg-database/pkg/product"
	"github.com/Scrackc/go-pg-database/storage"
)

func main() {
	// Migration DB
	migateDB()
	storage.NewPGDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	// * Crear un nuevo registro
	// m := &product.Model{
	// 	Name:         "Curso de db con GO",
	// 	Price:        70,
	// 	Observations: "On fire",
	// }
	// if err := serviceProduct.Create(m); err != nil {
	// 	log.Fatalf("product.Migrate: %v", err)
	// }
	// fmt.Printf("%+v\n", m)
	// * Para obtener un registro
	// ms, err := serviceProduct.GetByID(7)
	// switch {
	// case errors.Is(err, sql.ErrNoRows):
	// 	fmt.Println("No hay un producto con ese ID")
	// case err != nil:
	// 	log.Fatalf("product.GetAll: %v", err)
	// default:
	// 	fmt.Println(ms)

	// }
	// * Para obtener todos los productos
	// if err != nil {
	// 	log.Fatalf("product.GetAll: %v", err)
	// }
	// fmt.Println(ms)
	// * Para actualizar un producto
	// m := &product.Model{
	// 	ID:           20,
	// 	Name:         "DB with GO",
	// 	Observations: "Update this curse",
	// 	Price:        200,
	// }
	// err := serviceProduct.Upate(m)
	// if err != nil {
	// 	log.Fatalf("product.Update: %v", err)
	// }
	// * Para eliminar un producto
	err := serviceProduct.Delete(2)
	if err != nil {
		log.Fatalf("product.delete: %v", err)
	}

}

func migateDB() {
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
