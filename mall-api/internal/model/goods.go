package model

import (
	"time"
)

type Goods struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"size:100;not null" json:"name"`
	// Go 的 int 会被 MySQL 驱动映射成 bigint，只有 int32（位宽 32）才建出 day02 约定的 INT
	Price     int32     `gorm:"not null" json:"price"`
	Stock     int32     `gorm:"not null;default:0" json:"stock"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名，否则GORM默认去找 goods,此处防命名策略变更显式锁定
func (Goods) TableName() string {
	return "goods"
}
