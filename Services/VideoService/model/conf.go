package model

import (
	"VideoService/utils"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/micro/go-micro/v2/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

var db *gorm.DB
var redisDB *redis.Client
var err error

func InitDB() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		os.Exit(1)
	}
	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}

//func InitRedis() redis.Conn {
//	c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", utils.RedisHost, utils.RedisPort))
//	if err != nil {
//		fmt.Printf("redis.Dial() error:%v", err)
//		return nil
//	}
//	return c
//}

func InitRedis() {
	log.Info(utils.RedisAddr[0], utils.Pwd)
	redisDB = redis.NewClient(&redis.Options{
		Addr:     utils.RedisAddr[0],
		Password: utils.Pwd, // no password set
		DB:       0,         // use default DB
	})
}
