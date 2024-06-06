package repository

import (
	"fmt"

	"github.com/learnbasic/internal/adapters/domain"
)

// InMemoryProductoRepository es un repositorio en memoria para productos
type InMemoryProductoRepository struct {
	productos map[string]*domain.Producto
}

func NewInMemoryProductoRepository() *InMemoryProductoRepository {
	return &InMemoryProductoRepository{
		productos: map[string]*domain.Producto{
			"1": {ID: "1", Nombre: "Producto 1", Precio: 100.0},
			"2": {ID: "2", Nombre: "Producto 2", Precio: 200.0},
		},
	}
}

func (repo *InMemoryProductoRepository) GetProductoByID(id string) (*domain.Producto, error) {
	if producto, ok := repo.productos[id]; ok {
		return producto, nil
	}
	return nil, fmt.Errorf("producto no encontrado")
}

func (repo *InMemoryProductoRepository) PachtProductoByID(id string, precio float64) (*domain.Producto, error) {
	if producto, ok := repo.productos[id]; ok {
		producto.Precio = precio
		return producto, nil
	}
	return nil, fmt.Errorf("producto no encontrado")
}
func (repo *InMemoryProductoRepository) PostProductoByID(id string, nombre string, precio float64) (*domain.Producto, error) {
	if _, ok := repo.productos[id]; ok {
		return nil, fmt.Errorf("El producto ya existe")
	}

	producto := &domain.Producto{
		ID:     id,
		Nombre: nombre,
		Precio: precio,
	}
	repo.productos[id] = producto
	return producto, nil
}
func (repo *InMemoryProductoRepository) DeleteProductoByID(id string) (*domain.Producto, error) {
	if _, ok := repo.productos[id]; ok {
		delete(repo.productos, id)

		producto := &domain.Producto{
			ID:     id,
			Nombre: "",
			Precio: 0,
		}

		return producto, nil
	}
	return nil, fmt.Errorf("El producto no existe, por lo tanto no se puede eliminar")

}
