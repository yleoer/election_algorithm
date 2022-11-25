package global

import (
	"fmt"
	"log"

	"election_algorithm/pkg/config"
	"github.com/go-redis/redis/v8"
)

var internalMgr *mgr

type mgr struct {
	config config.Config
	redis  *redis.Client
}

func init() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("[Error] new config failed: %v", err)
	}

	redisCli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.SvcCfg.DbFqdn, cfg.SvcCfg.DbPort),
	})

	internalMgr = &mgr{
		config: cfg,
		redis:  redisCli,
	}
}

func GetConfig() config.Config {
	return internalMgr.config
}

func GetRedisCli() *redis.Client {
	return internalMgr.redis
}
