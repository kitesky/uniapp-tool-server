package boot

import (
	"strings"

	"github.com/hibiken/asynq"
)

type AsynqType struct {
	Client *asynq.Client
	Server *asynq.Server
}

func SetupAsynq() {
	addr := strings.Join([]string{Config.Redis.Addr, Config.Redis.Port}, ":")

	redisOption := asynq.RedisClientOpt{
		Addr:     addr,
		Password: Config.Redis.Password,
		DB:       Config.Redis.DB,
	}

	// 创建一个客户端
	client := asynq.NewClient(redisOption)

	// 创建一个服务器
	srv := asynq.NewServer(
		redisOption,
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	Asynq = AsynqType{
		Client: client,
		Server: srv,
	}
}
