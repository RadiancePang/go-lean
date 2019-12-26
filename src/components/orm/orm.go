package orm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-learn/components/logger"
	"go-learn/config"
	"go.uber.org/zap"
	"time"
)

func NewMySQL() (db *gorm.DB) {

	dataSouce := config.WebMetaConfig.Web.DataSource

	db, err := gorm.Open("mysql", dataSouce.Url)

	db.SingularTable(true)

	if err != nil {
		logger.EngineLogger.Error("数据库连接失败", zap.String("dsn", dataSouce.Url), zap.String("errorDetail", err.Error()))
		panic(err)
	}
	db.DB().SetMaxIdleConns(dataSouce.Idle)
	db.DB().SetMaxOpenConns(dataSouce.Active)
	db.DB().SetConnMaxLifetime(time.Duration(dataSouce.IdleTimeout) * time.Second)

	return

}
