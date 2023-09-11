package models

import (
	"gorm.io/gorm"
)

// Table-Model
type Leaderboard struct {
	gorm.Model
	Name  string
	Flags uint64
}

func ImportTable(db *gorm.DB) error {
	err := db.AutoMigrate(&Leaderboard{})
	return err
}
