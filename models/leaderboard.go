package models

import (
	"gorm.io/gorm"
)

// Table-Model
type Leaderboard struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255);not null"`
	Flags        uint64 `gorm:"type:uint;not null"`
	Email        string `gorm:"type:varchar(255);not null"`
	MobileNumber string `gorm:"type:varchar(255);not null"`
}

func ImportTable(db *gorm.DB) error {
	err := db.AutoMigrate(&Leaderboard{})
	return err
}
