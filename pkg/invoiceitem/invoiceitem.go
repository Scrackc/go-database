package invoiceitem

import (
	"database/sql"
	"time"
)

// Model of invoiceitem
type Model struct {
	ID              uint
	InvoiceHeaderId uint
	ProductId       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Models slice of Model
type Models []*Model

type Storage interface {
	Migrate() error
	CreateTX(*sql.Tx, uint, Models) error
}

// Service of invoiceitem
type Service struct {
	storage Storage
}

// NewService retorna un puntero de servicio
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate es utilizado para migrar invoiceitem
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
