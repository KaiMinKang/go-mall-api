package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"mall-api/internal/config"
)

func main() {
	// 1. 读配置
	cfg := config.Load()

	// 2. 拼 DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	// 3. 连 MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err) // 打印 + 退出，一步到位
	}

	// 4. Ping（GORM → *sql.DB → PingContext，带超时）
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取底层连接失败: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		log.Fatalf("数据库 Ping 失败: %v", err)
	}
	log.Println("mysql connected") // ← 验收标准要的日志

	// 5. 起 Gin
	r := gin.Default()

	// 6. 注册 GET /health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 7. Run
	if err := r.Run(fmt.Sprintf(":%d", cfg.ServerPort)); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
