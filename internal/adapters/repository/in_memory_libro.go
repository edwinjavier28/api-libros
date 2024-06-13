package repository

import (
	"fmt"

	"github.com/api-libros/internal/adapters/domain"
)

type LibroRepository interface {
	GetLibroByID(id string) (*domain.ResponseLibro, error)
	PostLibroByID(id string, req domain.LibroRequest) (*domain.ResponseLibro, error)
	DeleteLibroByID(id string) (*domain.ResponseLibro, error)
}

// InMemoryProductoRepository es un repositorio en memoria para productos
type InMemoryLibroRepository struct {
	libro map[string]*domain.ResponseLibro
}

func NewInMemoryProductoRepository() *InMemoryLibroRepository {
	return &InMemoryLibroRepository{
		libro: map[string]*domain.ResponseLibro{
			"1": {Nombre: "Lavoragine", Editorial: "Larouse", Autor: "Jose", Precio: 100.0, Cantidades: 3, Edicion: 2, BestSeller: true},
			"2": {Nombre: "El Coronel", Editorial: "Mamba", Autor: "Gabriel", Precio: 300.0, Cantidades: 2, Edicion: 1, BestSeller: true},
		},
	}
}

func (repo *InMemoryLibroRepository) GetLibroByID(id string) (*domain.ResponseLibro, error) {
	if libro, ok := repo.libro[id]; ok {
		return libro, nil
	}
	return nil, fmt.Errorf("libro no encontrado")
}

func (repo *InMemoryLibroRepository) PostLibroByID(id string, req domain.LibroRequest) (*domain.ResponseLibro, error) {
	if _, ok := repo.libro[id]; ok {
		return nil, fmt.Errorf("El producto ya existe")
	}

	libro := &domain.ResponseLibro{
		ID:         id,
		Nombre:     req.Nombre,
		Editorial:  req.Editorial,
		Autor:      req.Autor,
		Precio:     req.Precio,
		Cantidades: req.Cantidades,
		Edicion:    req.Edicion,
		BestSeller: req.Bestseller,
	}
	repo.libro[id] = libro
	return libro, nil
}
func (repo *InMemoryLibroRepository) DeleteLibroByID(id string) (*domain.ResponseLibro, error) {
	if _, ok := repo.libro[id]; ok {
		delete(repo.libro, id)

		libro := &domain.ResponseLibro{
			Nombre:     id,
			Editorial:  "",
			Autor:      "",
			Precio:     0,
			Cantidades: 0,
			Edicion:    0,
			BestSeller: false,
		}

		return libro, nil
	}
	return nil, fmt.Errorf("El libro no existe, por lo tanto no se puede eliminar")

}
