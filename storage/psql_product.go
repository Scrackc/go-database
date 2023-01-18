package storage

import (
	"database/sql"
	"fmt"

	"github.com/Scrackc/go-pg-database/pkg/product"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id)
	)`
	psqlCreateProduct  = `INSERT INTO products (name, observations, price, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	psqlGetAllProducts = `SELECT id, name, observations, price, created_at, updated_at FROM products`
	psqlGetProductById = psqlGetAllProducts + " WHERE id = $1"
	psqlUpdateProduct  = `UPDATE products SET name = $1, observations = $2, price = $3, updated_at = $4 WHERE id = $5`
	psqlDeleteProduct  = `DELETE FROM products WHERE id = $1`
)

// PSQLProduct usado par atrabajar con PG y el paquete product
type PSQLProduct struct {
	db *sql.DB
}

// NewPsqlProduct Retorna un nuevo puntero de PSQLProduct
func NewPsqlProduct(db *sql.DB) *PSQLProduct {
	return &PSQLProduct{db}
}

// Migrate implemneta la interfaz product.Storage
func (p *PSQLProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
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
func (p *PSQLProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(m.Name, stringToNull(m.Observations), m.Price, m.CreatedAt).Scan(&m.ID)
	if err != nil {
		return err
	}
	fmt.Println("Se creo el producto correctamente")
	return nil
}

// GetAll implemneta la interfaz product.Storage
func (p *PSQLProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProducts)
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
func (p *PSQLProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

// Upate implemneta la interfaz product.Storage
func (p *PSQLProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
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
func (p *PSQLProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)
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

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationsNul := sql.NullString{}
	updatedAtNul := sql.NullTime{}
	err := s.Scan(
		&m.ID, &m.Name, &observationsNul, &m.Price, &m.CreatedAt, &updatedAtNul,
	)
	if err != nil {
		return nil, err
	}
	fmt.Println(m.Name)
	m.Observations = observationsNul.String
	m.UpdatedAt = updatedAtNul.Time

	return m, nil
}
