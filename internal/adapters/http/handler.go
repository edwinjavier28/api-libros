package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LibroHandler struct {
	UseCase *usecases.GetProductoUseCase
}

func (h *LibroHandler) GetLibro(c *gin.Context) {
	id := c.Param("id")
	Libro, err := h.UseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Libro no encontrado"})
		return
	}
	c.JSON(http.StatusOK, Libro)
}

func (h *LibroHandler) PostLibro(c *gin.Context) {
	id := c.Param("id")
	type LibroRequest struct {
		Nombre     string  `json:"Nombre"`
		Editorial  string  `json:"Editorial"`
		Autor      string  `json:"Autor"`
		Precio     float64 `json:"Precio"`
		Cantidades int     `json:"Cantidades"`
		Edicion    int     `json:"Edicion"`
		Bestseller bool    `json:"Bestseller"`
	}
	var req LibroRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos"})
		return
	}

	libro, err := h.UseCase.PostExecute(id, req.Nombre, req.Precio)
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
