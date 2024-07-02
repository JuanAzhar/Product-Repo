package service

import (
	"errors"
	"mime/multipart"
	"product-rest-api/features/product/entity"
)

type productUsecase struct {
	productRepository entity.ProductsUseCaseInterface
}

func New(ProductUC entity.ProductsDataInterface) entity.ProductsUseCaseInterface {
	return &productUsecase{
		productRepository: ProductUC,
	}
}

// ReadSpecificProduct implements entity.ProductsUseCaseInterface.
func (productUC *productUsecase) ReadSpecificProduct(id string) (product entity.ProductsCore, err error) {
	if id == "" {
		return entity.ProductsCore{}, errors.New("product ID is required")
	}

	// Call the productRepository's ReadSpecificEvent method
	product, err = productUC.productRepository.ReadSpecificProduct(id)
	if err != nil {
		return entity.ProductsCore{}, err
	}

	// Check if the event is found in the repository, if not return an error
	if product.ID == "" {
		return entity.ProductsCore{}, errors.New("product not found")
	}

	return product, nil
}

// DeleteProduct implements entity.ProductsUseCaseInterface.
func (productUC *productUsecase) DeleteProduct(id string) (err error) {
	println("id dari service " + id)
	if id == "" {
		return errors.New("product not found")
	}

	errEvent := productUC.productRepository.DeleteProduct(id)
	if errEvent != nil {
		return errors.New("can't delete product " + errEvent.Error())
	}

	return nil
}

// PostProduct implements entity.ProductsUseCaseInterface.
func (productUC *productUsecase) PostProduct(data entity.ProductsCore, image *multipart.FileHeader) (row int, err error) {
	if data.Product_name == "" || data.Description == "" || data.Price == "" {
		return 0, errors.New("error, product name, description and price can't be empty")
	}

	if data.Quantity <= 0 {
		return 0, errors.New("error, quantity can't be empty")
	}

	if image != nil && image.Size > 10*1024*1024 {
		return 0, errors.New("image file size should be less than 10 MB")
	}

	errProduct, errPost := productUC.productRepository.PostProduct(data, image)
	if errPost != nil {
		return 0, errPost
	}
	return errProduct, nil
}

// ReadAllProduct implements entity.ProductsUseCaseInterface.
func (productUC *productUsecase) ReadAllProduct() ([]entity.ProductsCore, error) {
	products, err := productUC.productRepository.ReadAllProduct()
	if err != nil {
		return nil, errors.New("error get data")
	}

	return products, nil
}

// UpdateProduct implements entity.ProductsUseCaseInterface.
func (productUC *productUsecase) UpdateProduct(id string, data entity.ProductsCore, image *multipart.FileHeader) (product entity.ProductsCore, err error) {
	if id == "" {
		return entity.ProductsCore{}, errors.New("error, Product ID is required")
	}

	if data.Product_name == "" || data.Description == "" || data.Price == "" {
		return entity.ProductsCore{}, errors.New("error, product name, description and price can't be empty")
	}

	if data.Quantity < 0 {
		return entity.ProductsCore{}, errors.New("error, Quantity must be a positive integer")
	}

	if image != nil && image.Size > 10*1024*1024 {
		return entity.ProductsCore{}, errors.New("image file size should be less than 10 MB")
	}

	updatedproduct, err := productUC.productRepository.UpdateProduct(id, data, image)
	if err != nil {
		return entity.ProductsCore{}, err
	}

	return updatedproduct, nil
}
