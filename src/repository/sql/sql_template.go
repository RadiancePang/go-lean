package sql

import (
	"github.com/jinzhu/gorm"
	"go-learn/components/orm"
)

type Template struct {
	DB *gorm.DB
}

func New() (tmp *Template) {
	tmp = &Template{DB: orm.NewMySQL()}
	return
}
