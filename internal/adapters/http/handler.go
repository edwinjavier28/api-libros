package http

import (
	"net/http"

	"github.com/api-libros/internal/adapters/domain"
	usecases "github.com/api-libros/internal/adapters/use_cases"
	"github.com/gin-gonic/gin"
)

type LibroHandler struct {
	UseCase *usecases.LibroUseCase
}

func (h *LibroHandler) GetLibro(c *gin.Context) {
	id := c.Param("id")
	Libro, err := h.UseCase.Execute(id, domain.LibroRequest{})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Libro no encontrado"})
		return
	}
	c.JSON(http.StatusOK, Libro)
}

func (h *LibroHandler) PostLibro(c *gin.Context) {
	id := c.Param("id")

	var req domain.LibroRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos"})
		return
	}

	libro, err := h.UseCase.PostExecute(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Libro no encontrado"})
		return
	}
	c.JSON(http.StatusOK, libro)
}
func (h *LibroHandler) DeletePrecio(c *gin.Context) {
	id := c.Param("id")

	libro, err := h.UseCase.DeleteExecute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Libro no encontrado"})
		return
	}
	c.JSON(http.StatusOK, libro)
}
