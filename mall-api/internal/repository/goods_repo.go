package repository

import (
	"errors"
	"mall-api/internal/model"

	"gorm.io/gorm"
)

var ErrNotFound = errors.New("记录不存在")

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

func (r *GoodsRepo) GetGoodsByID(id uint) (*model.Goods, error) {
	var goods model.Goods
	err := r.db.First(&goods, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &goods, nil
}
