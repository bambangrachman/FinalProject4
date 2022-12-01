package faker

import (
	"finalproject4/model"
	"time"

	"gorm.io/gorm"
)

func Admin(db *gorm.DB) *model.User {
	return &model.User{
		GormModel: model.GormModel{
			ID:        0,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		Fullname: "Admin",
		Email:    "admin@admin.com",
		Password: "$2a$08$6e4EqvAc.Xjt1pYq0UwRpu7kA4R4xxggC8afu2u2boD2ld2K9Dkxu", // password123
		Role:     "admin",
	}
}
