package service

import (
	"errors"
	"mall-api/internal/model"
	"mall-api/internal/repository"

	"gorm.io/gorm"
)

type GoodsService struct {
	repo *repository.GoodsRepo
}

func NewGoodsService(repo *repository.GoodsRepo) *GoodsService {
	return &GoodsService{repo: repo}
}

// 查列表
func (s *GoodsService) GetGoods() ([]model.Goods, error) {
	return s.repo.GetGoods()
}

// 查询单条商品
var ErrGoodsNotFound = errors.New("商品不存在")

func (s *GoodsService) GetGoodsByID(id int) (*model.Goods, error) {
	goods, err := s.repo.GetGoodsByID(id)

	if err != nil {
		// 1. 只有明确是"记录没找到"，才翻译成业务错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGoodsNotFound
		}

		// 2. 其它错误（比如网络断了、表被删了），原样往上抛！
		return nil, err
	}
	return goods, nil
}
