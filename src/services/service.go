package services

import (
	"go-learn/components/cache"
	"go-learn/repository/sql"
)

type Service struct {
	SqlTemplate *sql.Template

	RedisTemplate *cache.RedisTemplate
}

func New() (service *Service) {

	service = &Service{
		SqlTemplate:   sql.New(),
		RedisTemplate: cache.New(),
	}
	return

}
