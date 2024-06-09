package usecases

import "internal\adapters\domain\producto.go"

type LibroRepository interface {
	GetProductoByID(nombre string) (*domain.Libro, error)
	PachtProductoByID(nombre string, editorial string) (*domain.Libro, error)
	PostProductoByID(nombre string, editorial string, autor string) (*domain.Libro, error)
	DeleteProductoByID(nombre string) (*domain.Libro, error)
}

type GetProductoUseCase struct {
	Repo ProductoRepository
}

func (uc *GetProductoUseCase) Execute(id string) (*domain.Libro, error) {
	return uc.Repo.GetProductoByID(id)
}

func (uc *GetProductoUseCase) UpdateExecute(id string, precio float64) (*domain.Libro, error) {
	return uc.Repo.PachtProductoByID(id, precio)
}

func (uc *GetProductoUseCase) PostExecute(id string, Nombre string, Precio float64) (*domain.Libro, error) {
	return uc.Repo.PostProductoByID(id, Nombre, Precio)
}
func (uc *GetProductoUseCase) DeleteExecute(id string) (*domain.Libro, error) {
	return uc.Repo.DeleteProductoByID(id)
}
