package repository

import (
	"errors"
	"fajar/apiMvc/config"
	"fajar/apiMvc/models"
)

func GetAllUser() ([]models.User, error) {
	var users []models.User
	tx := config.DB.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func InsertUser(user models.User) error {

	tx := config.DB.Create(&user) // proses insert data

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("Insert failed")
	}
	return nil
}
func GetAllUserId(id int) ([]models.User, error) {
	var users []models.User

	tx := config.DB.First(&users, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}
func UpdateUserId(id int, user models.User) (any, error) {

	// tx := config.DB.Where("id= ?", id).First(user, id).Updates(user)
	tx := config.DB.Model(user).Where("id = ?", id).Updates(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("updates failed")
	}
	return user, nil
}

func DeleteUser(id int) error {
	var user models.User
	tx := config.DB.Where("id= ?", id).Delete(&user)

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}
