package repository

import (
	"product-rest-api/features/user/entity"
	"product-rest-api/features/user/model"
	bcrypt "product-rest-api/utils/bcrypt"
	utils "product-rest-api/utils/jwt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) entity.UserDataInterface {
	return &userRepository{
		db: db,
	}
}

// DeleteUser implements entity.UserDataInterface.
func (userRepo *userRepository) DeleteUser(id string) (err error) {
	var chekId model.Users

	errData := userRepo.db.Where("id = ?", id).Delete(&chekId)
	if errData != nil {
		return errData.Error
	}

	return nil
}

// Login implements entity.UserDataInterface.
func (userRepo *userRepository) Login(email string, password string) (entity.UserCore, string, error) {
	var data model.Users

	tx := userRepo.db.Where("email=? AND password=?", email, password).First(&data)
	if tx.Error != nil {
		return entity.UserCore{}, "", tx.Error
	}

	var token string

	if tx.RowsAffected > 0 {
		var errToken error
		token, errToken = utils.CreateToken(data.ID)
		if errToken != nil {
			return entity.UserCore{}, "", errToken
		}
	}

	var resp = entity.UserCore{
		ID:       data.ID.String(),
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}

	return resp, token, nil
}

// Register implements entity.UserDataInterface.
func (userRepo *userRepository) Register(data entity.UserCore) (row int, err error) {

	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return 0, UUIDerr
	}

	hashPassword, err := bcrypt.HashPassword(data.Password)
	if err != nil {
		return 0, err
	}

	var input = model.Users{
		ID:       newUUID,
		Username: data.Username,
		Email:    data.Email,
		Password: hashPassword,
	}

	erruser := userRepo.db.Save(&input)
	if erruser.Error != nil {
		return 0, erruser.Error
	}

	return 1, nil

}

// ReadSpecificUser implements entity.UserDataInterface.
func (userRepo *userRepository) ReadSpecificUser(id string) (user entity.UserCore, err error) {
	var data model.Users
	errData := userRepo.db.Where("id=?", id).First(&data).Error
	if errData != nil{
		return entity.UserCore{}, errData
	}

	userCore := entity.UserCore{
		ID: data.ID.String(),
		Username: data.Username,
		Email: data.Email,
	}

	return userCore, nil
}
