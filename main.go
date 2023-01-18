package main

import (
	"fmt"
	"log"

	"github.com/Scrackc/go-database/storage"
)

func main() {

	driver := storage.MySQl
	storage.New(driver)
	mystorage, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("daoproduct %v", err)
	}
	ms, err := mystorage.GetAll()
	if err != nil {
		log.Fatalf("getall %v", err)
	}
	fmt.Println(ms)

}
