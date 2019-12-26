package cache

import (
	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
	"go-learn/common/model"
	"go-learn/components/logger"
	"go-learn/config"
	"time"
)

type RedisTemplate struct {
	cluster     redisc.Cluster
	redisConfig model.RedisConfig
}

type poolTemplate struct {
	redisConfig model.RedisConfig
}

func New() (redis *RedisTemplate) {

	redisConfig := config.WebMetaConfig.Redis

	poolTemplate := poolTemplate{redisConfig: redisConfig}

	cluster := redisc.Cluster{

		StartupNodes: redisConfig.Cluster.Nodes,

		DialOptions: initDialOptions(redisConfig),

		CreatePool: poolTemplate.creatPool,
	}

	if err := cluster.Refresh(); err != nil {

		logger.EngineLogger.Error(err.Error())

	} else {

		redis = &RedisTemplate{redisConfig: redisConfig, cluster: cluster}

	}

	return

}

func initDialOptions(redisConfig model.RedisConfig) []redis.DialOption {

	conntectTimeout := time.Duration(redisConfig.Connection.ConnectTimeout) * time.Second

	writTimeout := time.Duration(redisConfig.Connection.WriteTimeout) * time.Second

	readTimeout := time.Duration(redisConfig.Connection.ReadTimeout) * time.Second

	return []redis.DialOption{
		redis.DialConnectTimeout(conntectTimeout),
		redis.DialWriteTimeout(writTimeout),
		redis.DialReadTimeout(readTimeout),
	}

}

func (poolTemplate poolTemplate) creatPool(address string, options ...redis.DialOption) (*redis.Pool, error) {

	redispool := &redis.Pool{

		MaxActive: poolTemplate.redisConfig.Pool.MaxActive,

		MaxIdle: poolTemplate.redisConfig.Pool.MaxIdle,

		IdleTimeout: time.Duration(poolTemplate.redisConfig.Pool.IdleTimeout) * time.Second,

		MaxConnLifetime: time.Duration(poolTemplate.redisConfig.Pool.MaxConnLifetime) * time.Second,

		Dial: func() (conn redis.Conn, e error) {

			c, err := redis.Dial("tcp", address, options...)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("SELECT", 0); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil

		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {

			if time.Since(t) < time.Duration(poolTemplate.redisConfig.Pool.TestOnBorrow)*time.Second {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	return redispool, nil

}

func (template *RedisTemplate) Get(key string) (string, error) {

	return redis.String(template.cluster.Get().Do("GET", key))

}

func (template *RedisTemplate) Set(key string, value interface{}) (string, error) {

	return redis.String(template.cluster.Get().Do("SET", key, value))

}
