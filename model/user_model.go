package model

import (
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Fullname string `json:"fullname" gorm:"type:varchar(30);not null;unique"`
	Email    string `json:"email" gorm:"type:varchar(100);not null;unique"`
	Password string `json:"password,omitempty" gorm:"size:255;not null"`
	Role     string `json:"role" gorm:"size:20;not null"`
	Balance  int    `json:"balance" gorm:"not null"`
}

type UserRegisterRequest struct {
	FullName string `json:"full_name" valid:"required~Full name is required"`
	Email    string `json:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

type UserLoginRequest struct {
	Email    string `json:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
}

type UserBalanceRequest struct {
	Balance int `json:"balance" valid:"required~Your balance is required,range(0|50000000)~balance has reach the limit"`
}
type UserRegisterResponse struct {
	GormModel
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Passowrd string `json:"password"`
	Balance  int    `json:"balance"`
}

func HashPass(p string) string {
	salt := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = HashPass(u.Password)
	err = nil
	return
}
