package datasource

import (
	"main/config"
	"main/model"

	_ "github.com/go-sql-driver/mysql" //不能忘记导入
	"github.com/kataras/iris/v12"
	"xorm.io/xorm"
)

func NewMysqlEngine() *xorm.Engine {
	initConfig := config.InitConfig()
	if initConfig == nil {
		return nil
	}

	database := initConfig.DataBase

	dataSourceName := database.User + ":" + database.Pwd + "@tcp(" + database.Host + ")/" + database.Database + "?charset=utf8"

	engine, err := xorm.NewEngine(database.Drive, dataSourceName)
	iris.New().Logger().Info(database)

	err = engine.Sync2(new(model.User),
		new(model.Project),
		new(model.Estimate),
		new(model.EstimateDetail),
		new(model.Order),
		new(model.Delivery),
		new(model.Acceptance),
		new(model.Invoice),
		new(model.Customer),
		new(model.Product),
		new(model.Category),
		new(model.Timeline))

	if err != nil {
		panic(err.Error())
	}

	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)

	return engine
}
