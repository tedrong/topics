package models

import (
	"errors"
	"log"
	"strings"

	"github.com/topics/database"
	"github.com/topics/forms"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct{}

var authModel = new(AuthModel)

// Login, check database and compare the password
func (m UserModel) Login(form forms.LoginForm) (user database.User, token Token, err error) {
	result := database.GetPG(database.DBContent).Where("email = ?", strings.ToLower(form.Email)).Find(&user)
	if result.Error != nil {
		log.Panic(result.Error)
		return user, token, result.Error
	}

	//Compare the password form and database if match
	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
	if err != nil {
		return user, token, err
	}

	//Generate the JWT auth token
	tokenDetails, err := authModel.CreateToken(int64(user.Model.ID))
	if err != nil {
		return user, token, err
	}

	saveErr := authModel.CreateAuth(int64(user.Model.ID), tokenDetails)
	if saveErr == nil {
		token.AccessToken = tokenDetails.AccessToken
		token.RefreshToken = tokenDetails.RefreshToken
	}

	return user, token, nil
}

// Create record in database if there is a new user
func (m UserModel) Register(form forms.RegisterForm) (user database.User, err error) {
	db := database.GetPG(database.DBContent)

	//Check if the user exists in database
	result := db.Where("email = ?", strings.ToLower(form.Email)).Find(&user)
	if result.Error != nil {
		return user, errors.New("something went wrong, please try again later")
	}
	if result.RowsAffected > 0 {
		return user, errors.New("email already exists")
	}

	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	//Create the user and return back the user ID
	user.Name = form.Name
	user.Email = form.Email
	user.Password = string(hashedPassword)

	result = db.Create(&user)
	if result.Error != nil {
		return user, errors.New("something went wrong, please try again later")
	}
	return user, err
}
