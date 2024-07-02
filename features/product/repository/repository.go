package repository

import (
	"errors"
	"mime/multipart"
	"product-rest-api/features/product/entity"
	"product-rest-api/features/product/model"
	"product-rest-api/utils/cloudinary"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productsRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) entity.ProductsDataInterface {
	return &productsRepository{
		db: db,
	}
}

// ReadSpecificEvent implements entity.ProductsDataInterface.
func (productRepo *productsRepository) ReadSpecificProduct(id string) (product entity.ProductsCore, err error) {

	var productData model.Products
	errData := productRepo.db.Where("id = ?", id).First(&productData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return entity.ProductsCore{}, errors.New("product not found")
		}
		return entity.ProductsCore{}, errData
	}

	eventCore := entity.ProductsCore{
		ID:            productData.ID.String(),
		Product_image: productData.Product_image,
		Product_name:  productData.Product_name,
		Description:   productData.Description,
		Price:         productData.Price,
		Quantity:      productData.Quantity,
		CreatedAt:     productData.CreatedAt,
		UpdatedAt:     productData.UpdatedAt,
	}

	return eventCore, nil
}

// PostEvent implements entity.ProductsDataInterface.
func (productRepo *productsRepository) PostProduct(data entity.ProductsCore, image *multipart.FileHeader) (row int, err error) {
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return 0, UUIDerr
	}

	// imageURL, uploadErr := storage.UploadPoster(image)
	// if uploadErr != nil {
	// 	return 0, uploadErr
	// }

	// if uploadErr != nil {
	// 	return 0, uploadErr
	// }
	file, err := image.Open()
	if err != nil {
		return 0, err
	}
	defer file.Close()

	imageURL, err := cloudinary.UploadToCloudinary(file, image.Filename)
	if err != nil {
		return 0, err
	}
	data.Product_image = imageURL

	var input = model.Products{
		ID:            newUUID,
		Product_image: data.Product_image,
		Product_name:  data.Product_name,
		Description:   data.Description,
		Price:         data.Price,
		Quantity:      data.Quantity,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
	}

	errData := productRepo.db.Save(&input)
	if errData != nil {
		return 0, errData.Error
	}

	return 1, nil
}

// ReadAllEvent implements entity.ProductsDataInterface.
func (productRepo *productsRepository) ReadAllProduct() ([]entity.ProductsCore, error) {
	var dataProduct []model.Products

	errData := productRepo.db.Find(&dataProduct).Error
	if errData != nil {
		return nil, errData
	}

	mapData := make([]entity.ProductsCore, len(dataProduct))
	for i, value := range dataProduct {
		mapData[i] = entity.ProductsCore{
			ID:            value.ID.String(),
			Product_image: value.Product_image,
			Product_name:  value.Product_name,
			Description:   value.Description,
			Price:         value.Price,
			Quantity:      value.Quantity,
			CreatedAt:     value.CreatedAt,
			UpdatedAt:     value.UpdatedAt,
		}
	}
	return mapData, nil
}

// UpdateProduct implements entity.ProductsDataInterface.
func (productRepo *productsRepository) UpdateProduct(id string, data entity.ProductsCore, image *multipart.FileHeader) (product entity.ProductsCore, err error) {
	var productData model.Products
	errData := productRepo.db.Where("id = ?", id).First(&productData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return entity.ProductsCore{}, errors.New("product not found")
		}
		return entity.ProductsCore{}, errData
	}

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return entity.ProductsCore{}, err
	}

	var imageURL string
	if image != nil {
		// Open the file
		file, err := image.Open()
		if err != nil {
			return entity.ProductsCore{}, err
		}
		defer file.Close()

		// Upload the image to Cloudinary
		imageURL, err = cloudinary.UploadToCloudinary(file, image.Filename)
		if err != nil {
			return entity.ProductsCore{}, err
		}
	} else {
		// Use existing image if no new image is uploaded
		imageURL = productData.Product_image
	}

	// Set the new product data
	productData.ID = uuidID
	productData.Product_image = imageURL
	productData.Product_name = data.Product_name
	productData.Description = data.Description
	productData.Price = data.Price
	productData.Quantity = data.Quantity
	productData.UpdatedAt = data.UpdatedAt

	// Update the product in the database
	errSave := productRepo.db.Save(&productData).Error
	if errSave != nil {
		return entity.ProductsCore{}, errSave
	}

	// Create the response product core
	productCore := entity.ProductsCore{
		ID:            productData.ID.String(),
		Product_image: productData.Product_image,
		Product_name:  productData.Product_name,
		Description:   productData.Description,
		Price:         productData.Price,
		Quantity:      productData.Quantity,
		CreatedAt:     productData.CreatedAt,
		UpdatedAt:     productData.UpdatedAt,
	}

	return productCore, nil
}


// DeleteEvents implements entity.ProductsDataInterface.
func (productRepo *productsRepository) DeleteProduct(id string) (err error) {
	var checkId model.Products
	println("id dari repo " + id)

	// dataId := productRepo.db.Where("id = ?", id).First(&checkId)
	// if dataId != nil {
	// 	return errors.New("product not found")
	// }

	errData := productRepo.db.Where("id = ?", id).Delete(&checkId)
	
	if errData.RowsAffected == 0 {
		return errors.New("data not found")
	}

	if errData != nil {
		return errData.Error
	}

	
	return nil
}
