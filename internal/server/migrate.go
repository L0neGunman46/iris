package server

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type USERTABLE struct {
	ID       uint       `gorm:"primaryKey" json:"id"`
	Email    string     `gorm:"size:255;uniqueIndex;not null" json:"email"`
	UserName string     `gorm:"size:255;uniqueIndex;not null" json:"userName"`
	Password string     `gorm:"size:255;not null" json:"password"`
	URLs     []URLTABLE `gorm:"foreignKey:UserID"`
}

type URLTABLE struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	UserID          uint      `gorm:"not null" json:"userId"`
	User            USERTABLE `gorm:"foreignKey:UserID" json:"user"`
	CustomShort     string    `gorm:"size:255;uniqueIndex" json:"short"`
	XRateRemaining  uint      `gorm:"not null" json:"rate_limit"`
	XRateLimitReset uint      `gorm:"not null" json:"rate_limit_reset"`
}

func (u *USERTABLE) BeforeCreate(tx *gorm.DB) (err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

func MigrateURLs(db *gorm.DB) error {
	err := db.AutoMigrate(&USERTABLE{}, &URLTABLE{})
	return err
}
