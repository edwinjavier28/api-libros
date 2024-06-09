package usecases

import (
	"github.com/api-libros/internal/adapters/domain"
	respositoy "github.com/api-libros/internal/adapters/repository"
)

type LibroUseCase struct {
	Repo respositoy.LibroRepository
}

func (uc *LibroUseCase) Execute(id string, req domain.LibroRequest) (*domain.ResponseLibro, error) {
	return uc.Repo.GetLibroByID(id)
}

func (uc *LibroUseCase) PostExecute(id string, req domain.LibroRequest) (*domain.ResponseLibro, error) {
	return uc.Repo.PostLibroByID(id, req)
}
func (uc *LibroUseCase) DeleteExecute(id string) (*domain.ResponseLibro, error) {
	return uc.Repo.DeleteLibroByID(id)
}
