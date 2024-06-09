package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Crear el repositorio
	productoRepo := repository.NewInMemoryProductoRepository()

	// Crear el caso de uso
	getProductoUseCase := &usecases.GetProductoUseCase{
		Repo: productoRepo,
	}

	// Crear el handler
	productoHandler := &http.ProductoHandler{
		UseCase: getProductoUseCase,
	}

	// Definir la ruta
	r.GET("/producto/:id", productoHandler.GetProducto)
	r.PATCH("/producto/:id", productoHandler.UpdatePrecio)
	r.POST("/producto/:id", productoHandler.PostPrecio)
	r.DELETE("/producto/:id", productoHandler.DeletePrecio)

	// Iniciar el servidor
	r.Run(":8086")
}
