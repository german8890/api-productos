package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/learnbasic/internal/adapters/usecases"
)

type ProductoHandler struct {
	UseCase *usecases.GetProductoUseCase
}

func (h *ProductoHandler) GetProducto(c *gin.Context) {
	id := c.Param("id")
	producto, err := h.UseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Producto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, producto)
}

func (h *ProductoHandler) UpdatePrecio(c *gin.Context) {
	id := c.Param("id")
	type PrecioRequest struct {
		Precio float64 `json:"precio,string"`
	}
	var req PrecioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos"})
		return
	}

	producto, err := h.UseCase.UpdateExecute(id, req.Precio)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Producto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, producto)
}

func (h *ProductoHandler) PostPrecio(c *gin.Context) {
	id := c.Param("id")
	type PrecioRequest struct {
		Precio float64 `json:"precio"`
		Nombre string  `json:"Nombre"`
	}
	var req PrecioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos"})
		return
	}

	producto, err := h.UseCase.PostExecute(id, req.Nombre, req.Precio)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Producto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, producto)
}
func (h *ProductoHandler) DeletePrecio(c *gin.Context) {
	id := c.Param("id")

	producto, err := h.UseCase.DeleteExecute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Producto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, producto)
}
