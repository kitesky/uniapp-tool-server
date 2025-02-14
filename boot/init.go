package boot

import (
	"app-api/config"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Asynq  AsynqType
	Config *config.Config
	DB     *gorm.DB
	Redis  *redis.Client
)

func Init() {
	SetupConfig() // 初始化配置信息
	SetupRedis()  // 初始化redis
	SetupMySQL()  // 初始化数据库
	SetupAsynq()  // 初始化任务队列
}
