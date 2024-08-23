package initialize

import (
	"fmt"
	"goods_srv/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func InitDB() {
	log.Println("Database connected")
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	c := global.ServerConfig.MysqlInfo
	//"root", "888888", "127.0.0.1", 3306, "mxshop"
	dsm := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?loc=Local&charset=utf8mb4,utf8&parseTime=true", c.User, c.Password, c.Host, c.Port, c.Name)

	var err error
	global.Db, err = gorm.Open(mysql.Open(dsm), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "mx_",
			SingularTable: true,
		},
		Logger:      newLogger,
		PrepareStmt: true, // 启用预编译语句以提高性能
	})

	if err != nil {
		log.Println("err:", err.Error())
	}

	// 迁移 schema
	//_ = Db.AutoMigrate(&model.User{})
}
