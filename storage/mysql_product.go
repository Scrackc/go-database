package storage

import (
	"database/sql"
	"fmt"

	"github.com/Scrackc/go-database/pkg/product"
)

const (
	mysqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`
	mysqlCreateProduct  = `INSERT INTO products (name, observations, price, created_at) VALUES (?, ?, ?, ?)`
	mysqlGetAllProducts = `SELECT id, name, observations, price, created_at, updated_at FROM products`
	mysqlGetProductById = mysqlGetAllProducts + " WHERE id = ?"
	mysqlUpdateProduct  = `UPDATE products SET name = ?, observations = ?, price = ?, updated_at = ? WHERE id = ?`
	mysqlDeleteProduct  = `DELETE FROM products WHERE id = ?`
)

// mySQLProduct usado par atrabajar con MYSQL y el paquete product
type mySQLProduct struct {
	db *sql.DB
}

// NewMySQLProduct Retorna un nuevo puntero de MySQLProduct
func newMySQLProduct(db *sql.DB) *mySQLProduct {
	return &mySQLProduct{db}
}

// Migrate implemneta la interfaz product.Storage
func (p *mySQLProduct) Migrate() error {
	stmt, err := p.db.Prepare(mysqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Migracion de producto ejecutada correctamente")
	return nil
}

// Create implemneta la interfaz product.Storage
func (p *mySQLProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(mysqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(m.Name, stringToNull(m.Observations), m.Price, m.CreatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = uint(id)
	fmt.Println("Se creo el producto correctamente con id: ", id)
	return nil
}

// GetAll implemneta la interfaz product.Storage
func (p *mySQLProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(mysqlGetAllProducts)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ms := make(product.Models, 0)
	fmt.Println("AQUI")
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil
}

// GetByID implemneta la interfaz product.Storage
func (p *mySQLProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(mysqlGetProductById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

// Upate implemneta la interfaz product.Storage
func (p *mySQLProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(mysqlUpdateProduct)
	if err != nil {
		return err
	}

	defer stmt.Close()
	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNULL(m.UpdatedAt),
		m.ID,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no existe el producto con id: %d", m.ID)
	}
	fmt.Println("Se actualizo el producto correctamente")
	return nil
}

// Delete implemneta la interfaz product.Storage
func (p *mySQLProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(mysqlDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no existe el producto con id: %d", id)
	}
	fmt.Println("Se elimino el producto")
	return nil
}
