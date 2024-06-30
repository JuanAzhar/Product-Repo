package entity

import "mime/multipart"

type ProductsDataInterface interface {
	PostProduct(data ProductsCore, image *multipart.FileHeader) (row int, err error)
	ReadAllProduct() ([]ProductsCore, error)
	ReadSpecificProduct(id string) (product ProductsCore, err error)
	UpdateProduct(id string, data ProductsCore, image *multipart.FileHeader) (product ProductsCore, err error)
	DeleteProduct(id string) (err error)
}

type ProductsUseCaseInterface interface {
	PostProduct(data ProductsCore, image *multipart.FileHeader) (row int, err error)
	ReadAllProduct() ([]ProductsCore, error)
	ReadSpecificProduct(id string) (product ProductsCore, err error)
	UpdateProduct(id string, data ProductsCore, image *multipart.FileHeader) (product ProductsCore, err error)
	DeleteProduct(id string) (err error)
}
