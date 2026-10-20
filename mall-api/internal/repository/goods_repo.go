package repository

import (
	"mall-api/internal/model"

	"gorm.io/gorm"
)

type GoodsRepo struct {
	db *gorm.DB
}

func NewGoodsRepo(db *gorm.DB) *GoodsRepo {
	return &GoodsRepo{db: db}
}

func (r *GoodsRepo) GetGoods() ([]model.Goods, error) {
	var goods []model.Goods = []model.Goods{}
	err := r.db.Find(&goods).Error
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func (r *GoodsRepo) GetGoodsByID(id int) (*model.Goods, error) {
	var goods model.Goods
	err := r.db.First(&goods, id).Error
	if err != nil {
		return nil, err
	}
	return &goods, nil
}
