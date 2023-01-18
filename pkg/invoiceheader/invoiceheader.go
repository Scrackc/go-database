package invoiceheader

import (
	"database/sql"
	"time"
)

// Model of invoiceheader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Storage interface {
	Migrate() error
	CreateTX(*sql.Tx, *Model) error
}

// Service of invoiceheader
type Service struct {
	storage Storage
}

// NewService retorna un puntero de servicio
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate es utilizado para migrar invoiceheader
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
