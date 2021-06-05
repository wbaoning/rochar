package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"goginProject/setting"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func Init(cfg *setting.Redis) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: "",
		DB:       cfg.Pb,
		PoolSize: cfg.PoolSize,
	})

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	_, err = client.Ping().Result()

	return err
}

func Close() {
	client.Close()
}
