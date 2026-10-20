package model

import (
	"time"
)

type Goods struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Price     int64     `gorm:"not null" json:"price"`
	Stock     int       `gorm:"not null;default:0"  json:"stock"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名，否则GORM默认去找 goods
func (Goods) TableName() string {
	return "goods"
}
