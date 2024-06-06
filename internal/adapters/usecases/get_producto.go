package usecases

import "github.com/learnbasic/internal/adapters/domain"

type ProductoRepository interface {
	GetProductoByID(id string) (*domain.Producto, error)
	PachtProductoByID(id string, precio float64) (*domain.Producto, error)
	PostProductoByID(id string, Nomnre string, Precio float64) (*domain.Producto, error)
	DeleteProductoByID(id string) (*domain.Producto, error)
}

type GetProductoUseCase struct {
	Repo ProductoRepository
}

func (uc *GetProductoUseCase) Execute(id string) (*domain.Producto, error) {
	return uc.Repo.GetProductoByID(id)
}

func (uc *GetProductoUseCase) UpdateExecute(id string, precio float64) (*domain.Producto, error) {
	return uc.Repo.PachtProductoByID(id, precio)
}

func (uc *GetProductoUseCase) PostExecute(id string, Nombre string, Precio float64) (*domain.Producto, error) {
	return uc.Repo.PostProductoByID(id, Nombre, Precio)
}
func (uc *GetProductoUseCase) DeleteExecute(id string) (*domain.Producto, error) {
	return uc.Repo.DeleteProductoByID(id)
}
