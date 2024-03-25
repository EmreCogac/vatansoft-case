package models

import (
	"app/app/database"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

type Posts struct {
	gorm.Model
	Header string `json:"header" gorm:"type:varchar(30);not null" binding:"required"`
	State  string `json:"state" gorm:"type:varchar(2);not null" binding:"required"`
	Title  string `json:"title"  gorm:"type:varchar(40);not null" binding:"required"`
	Tarih  string `json:"tarih" binding:"required" gorm:"unique"`
	Uid    uint
	User   *User `gorm:"foreignKey:Uid"`
}

func (now *Posts) UpdatePost(updateted *Posts, id int) error {
	db := database.GlobalDB
	now.UpdatedAt = time.Now()
	now.Title = updateted.Title
	now.Header = updateted.Header
	now.Tarih = updateted.Tarih
	err := db.Where("id= ?", id).Updates(&now)
	if err != nil {
		return err.Error
	}
	return nil

}

func (now *Posts) UpdateState(updated *Posts, id uint) error {
	db := database.GlobalDB
	now.UpdatedAt = time.Now()
	now.State = updated.State
	err := db.Where("id= ?", id).Updates(&now)
	if err != nil {
		return err.Error
	}
	return nil

}

func (deleted *Posts) DeletePost(id int) error {
	db := database.GlobalDB
	err := db.Unscoped().Where("id =? ", id).Delete(&deleted)
	if err != nil {
		return err.Error
	}
	return nil
}

func (post *Posts) CreatePost() error {
	result := database.GlobalDB.Create(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (user *User) CreateUserRecord() error {
	result := database.GlobalDB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
