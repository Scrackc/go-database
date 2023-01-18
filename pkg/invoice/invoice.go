package invoice

import (
	"github.com/Scrackc/go-database/pkg/invoiceheader"
	"github.com/Scrackc/go-database/pkg/invoiceitem"
)

// Model of invoice
type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}

// Storage interface que implementa storage db
type Storage interface {
	Create(*Model) error
}

// Service of invoice
type Service struct {
	storage Storage
}

// NewService retorna un puntero de servicio
func NewService(s Storage) *Service {
	return &Service{s}
}

// Create una nueva factura
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
