package main

import (
	"flag"
	"go-learn/components/cache"
	"go-learn/components/orm"
	"go-learn/components/tools"
	"go-learn/config"
	"go-learn/controller"
)

func main() {

	configInfo := flag.String("config", "", "运行环境配置信息")

	flag.Parse()

	configName := *configInfo

	if !tools.IsNotEmpty(configName) {

		panic("运行环境配置信息参数缺失")

	}

	//加载配置信息
	config.InitMeta(configName)

	//初始化连接池
	orm.Db = orm.NewMySQL()

	//初始化redis
	cache.New()

	controller.Start()

}
