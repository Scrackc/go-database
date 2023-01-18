package product

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrIDNotFound = errors.New("el producto no contiene un id")
)

// MOdel of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s \n",
		m.ID, m.Name, m.Observations, m.Price, m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02"))
}

// Models slice of Model
type Models []*Model

type Storage interface {
	Migrate() error
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error
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

// Create es usaso para crear un producto
func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}

// GetAll es usado para obtener todos los productos
func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

// GetByID es usado para obtener todos los productos
func (s *Service) GetByID(id uint) (*Model, error) {
	return s.storage.GetByID(id)
}

// Upate es usado para actualizar un producto
func (s *Service) Upate(m *Model) error {
	if m.ID == 0 {
		return ErrIDNotFound
	}

	m.UpdatedAt = time.Now()

	return s.storage.Update(m)
}

// Delete es usado para eliminar un producto
func (s *Service) Delete(id uint) error {
	return s.storage.Delete(id)
}
