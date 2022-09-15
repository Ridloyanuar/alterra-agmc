package database

import (
	"advance-golang/config"
	"advance-golang/middlewares"
	"advance-golang/models"
)

func GetUsers() (interface{}, error) {
	var users []models.Users

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUserByID(id string) (interface{}, error) {
	var user models.Users

	if e := config.DB.First(&user, id).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func DeleteUserByID(id string) (interface{}, error) {
	var user models.Users

	if e := config.DB.Delete(&user, id).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func UserLogin(user *models.Users) (interface{}, error) {

	var err error
	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DetailUser(id int) (interface{}, error) {

	var user models.Users

	if e := config.DB.Find(&user, id).Error; e != nil {
		return nil, e
	}

	return user, nil
}
