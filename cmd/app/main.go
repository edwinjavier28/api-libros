package main

import (
	"github.com/api-libros/internal/adapters/http"
	"github.com/api-libros/internal/adapters/repository"
	usecases "github.com/api-libros/internal/adapters/use_cases"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Crear el repositorio
	productoRepo := repository.NewInMemoryProductoRepository()

	// Crear el caso de uso
	getProductoUseCase := &usecases.LibroUseCase{
		Repo: productoRepo,
	}

	// Crear el handler
	LibrosHandler := &http.LibroHandler{
		UseCase: getProductoUseCase,
	}

	// Definir la ruta

	r.POST("/Libros/:id", LibrosHandler.PostLibro)
	r.GET("/Libros/:id", LibrosHandler.GetLibro)
	r.DELETE("/Libros/:id", LibrosHandler.DeleteLibro)
	// Iniciar el servidor
	r.Run(":8088")
}
