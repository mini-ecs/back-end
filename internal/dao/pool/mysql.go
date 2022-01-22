package pool

import (
	"database/sql"
	"fmt"
	"github.com/mini-ecs/back-end/internal/model"
	"github.com/mini-ecs/back-end/pkg/config"
	"github.com/mini-ecs/back-end/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB

func init() {
	username := config.GetConfig().MySQL.User     //账号
	password := config.GetConfig().MySQL.Password //密码
	host := config.GetConfig().MySQL.Host         //数据库地址，可以是Ip或者域名
	port := config.GetConfig().MySQL.Port         //数据库端口
	Dbname := config.GetConfig().MySQL.Name       //数据库名
	timeout := "10s"                              //连接超时，10秒

	var dsn string
	if config.GetConfig().Debug {
		log.GetGlobalLogger().Infof("DebugMode: true")
		//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	} else {
		log.GetGlobalLogger().Infof("DebugMode: false")
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(db:%d)/", username, password, port))
		if err != nil {
			panic(err)
		}
		_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + Dbname)
		if err != nil {
			panic(err)
		}
		db.Close()

		dsn = fmt.Sprintf("%s:%s@tcp(db:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, port, Dbname, timeout)
	}
	var err error
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	sqlDB, _ := _db.DB()
	err = _db.AutoMigrate(&model.VM{}, &model.User{}, &model.UserType{}, &model.ImageOrSnapshot{}, &model.Snapshot{})
	if err != nil {
		panic(err)
	}
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

func GetDB() *gorm.DB {
	return _db
}
