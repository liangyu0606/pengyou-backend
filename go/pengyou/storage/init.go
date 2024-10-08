package storage

import (
	"pengyou/global/config"
	db "pengyou/storage/database"
	rds "pengyou/storage/redis"
)

func Init(cfg *config.Config) {
	// init redis
	rds.InitRedis(&cfg.Redis)

	// init mysql
	db.InitMySQL(cfg)

	// InitFile(cfg)

	InitElasticSearch(cfg)
}
