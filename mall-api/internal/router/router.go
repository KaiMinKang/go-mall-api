package router

import (
	"net/http"

	"mall-api/internal/handler"
	"mall-api/internal/repository"
	"mall-api/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRouter(db *gorm.DB) *gin.Engine {
	// 注册路由
	r := gin.Default()

	// 依赖注入
	repo := repository.NewGoodsRepo(db)
	svc := service.NewGoodsService(repo)
	h := handler.NewGoodsHandler(svc)

	apiGroup := r.Group("/api/v1")
	{
		apiGroup.GET("/goods", h.GetGoodsList)
		apiGroup.GET("/goods/:id", h.GetGoodsByID)
	}

	// 健康检查在根路径：不属于业务 API 分组
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}
