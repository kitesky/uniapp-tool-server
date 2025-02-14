package boot

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func SetupMySQL() *gorm.DB {
	cfg := Config.Database
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,     // user
		cfg.Password, // password
		cfg.Host,     // host
		cfg.Port,     // port
		cfg.Database, // DB name
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Info),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		PrepareStmt:            false,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
			TablePrefix:   cfg.Prefix,
		},
	})

	if err != nil {
		log.Fatal("连接数据库失败，请检查参数：", err)
		os.Exit(1)
	}

	conn, _ := DB.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	conn.SetMaxIdleConns(cfg.MaxIdleConns)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	conn.SetMaxOpenConns(cfg.MaxOpenConns)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	conn.SetConnMaxLifetime(1 * time.Hour)

	return DB
}
