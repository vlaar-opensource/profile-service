package db

import "github.com/jinzhu/gorm"

//UserAccount blabla
type UserAccount struct {
	gorm.Model
	Username         string `gorm:"type:varchar(100);unique_index;not null"`
	Password         string `gorm:"not null"`
	Name             string
	Email            string `gorm:"type:varchar(100)"`
	Phone            string
	PhoneConfirmed   bool `gorm:"default:false"`
	EmailConfirmed   bool `gorm:"default:false"`
	ProfileCompleted bool `gorm:"default:false"`
	AddressFilled    bool `gorm:"default:false"`
	DateOfBirth      int
}

type address struct {
	gorm.Model
	Country   string
	City      string
	Lat       int
	Lng       int
	Text      string
	IsPrimary bool
}
