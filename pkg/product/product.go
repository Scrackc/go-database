package product

import "time"

// MOdel of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Models slice of Model
type Models []*Model

type Storage interface {
	Migrate() error
	// Create(*Model) error
	// Update(*Model) error
	// GetAll() (Models, error)
	// GetByID(uint) (*Model, error)
	// Delete(uint) error
}

// Service of product
type Service struct {
	storage Storage
}

// NewService retorna un puntero de servicio
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate es utilizado para migrar producto
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
