package cache

import (
	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
	"go-learn/common/model"
	"go-learn/components/logger"
	"go-learn/config"
	"time"
)

var cluster redisc.Cluster

var redisConfig model.RedisConfig

func New() {

	redisConfig = config.WebMetaConfig.Redis

	cluster = redisc.Cluster{

		StartupNodes: redisConfig.Cluster.Nodes,

		DialOptions: initDialOptions(),

		CreatePool: creatPool,
	}

	if err := cluster.Refresh(); err != nil {

		logger.EngineLogger.Error(err.Error())

	}

}

func initDialOptions() []redis.DialOption {

	conntectTimeout := time.Duration(redisConfig.Connection.ConnectTimeout) * time.Second

	writTimeout := time.Duration(redisConfig.Connection.WriteTimeout) * time.Second

	readTimeout := time.Duration(redisConfig.Connection.ReadTimeout) * time.Second

	return []redis.DialOption{
		redis.DialConnectTimeout(conntectTimeout),
		redis.DialWriteTimeout(writTimeout),
		redis.DialReadTimeout(readTimeout),
	}

}

func creatPool(address string, options ...redis.DialOption) (*redis.Pool, error) {

	redispool := &redis.Pool{

		MaxActive: redisConfig.Pool.MaxActive,

		MaxIdle: redisConfig.Pool.MaxIdle,

		IdleTimeout: time.Duration(redisConfig.Pool.IdleTimeout) * time.Second,

		MaxConnLifetime: time.Duration(redisConfig.Pool.MaxConnLifetime) * time.Second,

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

			if time.Since(t) < time.Duration(redisConfig.Pool.TestOnBorrow)*time.Second {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	return redispool, nil

}

func Get(key string) (string, error) {

	return redis.String(cluster.Get().Do("GET", key))

}

func Set(key string, value interface{}) (string, error) {

	return redis.String(cluster.Get().Do("SET", key, value))

}
