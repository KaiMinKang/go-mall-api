package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"mall-api/internal/router"
	"mall-api/internal/seed"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"mall-api/internal/config"
)

func main() {
	// 1. 读配置
	cfg := config.Load()

	// 2. 拼 DSN
	dsn := config.GetDsn(cfg)

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

	// 5. 建表 + 写入种子数据
	if err := seed.Migrate(db); err != nil {
		log.Fatalf("建表失败: %v", err)
	}

	if err := seed.Seed(db); err != nil {
		log.Fatalf("写入种子数据失败: %v", err)
	}

	// 6. 组装路由
	r := router.RegisterRouter(db)

	// 7. 启动
	if err := r.Run(fmt.Sprintf(":%d", cfg.ServerPort)); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
