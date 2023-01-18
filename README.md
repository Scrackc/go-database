# CRUD GO

1. Levantar db 
```
docker compose up -d
```
2. Ejecutar 
```
go run main.go
```

##### NOTA 
Para cada accion se debe de tener abierta la conexión a la base de datos Se debe de cambiar al metodo dependiendo la base de datos PG o MYSQL
```
    storage.NewPGDB() o storage.NewMySqlDB()
```
## Migrar tablas
```
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
```
## CRUD SIN DAO
### Crear un nuevo producto
```
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	m := &product.Model{
		Name:  "Curso de GO",
		Price: 10,
		// Observations: "On fire",
	}
	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
    // Visualizar el producto registrado
	fmt.Printf("%+v\n", m)
```
### Obtener todos los registros
```
    storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
    ms, err := serviceProduct.GetAll()
    if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}
	fmt.Println(ms)
```
### Obtener un registro
```
    storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
    ms, err := serviceProduct.GetByID(7)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("No hay un producto con ese ID")
	case err != nil:
		log.Fatalf("product.GetAll: %v", err)
	default:
		fmt.Println(ms)
```
### Actualizar un registro
```
    m := &product.Model{
		ID:           20,
		Name:         "DB with GO",
		Observations: "Update this curse",
		Price:        200,
	}
	err := serviceProduct.Upate(m)
	if err != nil {
		log.Fatalf("product.Update: %v", err)
	}
```
### Eliminar un registro
```
    err := serviceProduct.Delete(2)
	if err != nil {
		log.Fatalf("product.delete: %v", err)
	}
```
### Para crear una factura
```
    storageHeader := storage.NewPSQLInvoiceHeader(storage.Pool())
	storageItems := storage.NewPSQLInvoiceItem(storage.Pool())
	storageInvoice := storage.NewPsqlInvoice(storage.Pool(), storageHeader, storageItems)
	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Eduardo",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductId: 1},
			&invoiceitem.Model{ProductId: 99},
		},
	}
	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}
```

