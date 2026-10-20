package seed

import (
	"mall-api/internal/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.Goods{})
}

func Seed(db *gorm.DB) error {
	// 第一步永远是"判空"：Seed 会在每次启动时被调用，
	// 不判空的话，每重启一次就多插 3 条，表里很快堆满重复数据
	var count int64
	if err := db.Model(&model.Goods{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	// 第二步是"插入数据"
	// 传切片 → GORM 一次生成多条 INSERT（批量插入）
	goods := []model.Goods{
		{Name: "测试商品A", Price: 1999, Stock: 100},
		{Name: "测试商品B", Price: 2999, Stock: 50},
		{Name: "测试商品C", Price: 9900, Stock: 0},
	}
	return db.Create(&goods).Error
}
