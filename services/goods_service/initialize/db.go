package initialize

import (
	"fmt"

	"github.com/zhaohaihang/goods_service/global"
	"github.com/zhaohaihang/goods_service/model"
	"go.uber.org/zap"

	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB
// @Description: 初始化DB
func InitDB() {
	var err error
	MySqlInfo := global.ServiceConfig.MySqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", MySqlInfo.User, MySqlInfo.Password, MySqlInfo.Host, MySqlInfo.Port, MySqlInfo.Name)
	// 创建日志文件
	newLogger := logger.New(log.New(logFileWriter, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      false,
	})

	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		zap.S().Fatalw("gorm open dsn failed: %s", "err", err.Error())
	}
	err = global.DB.AutoMigrate(&model.Category{}, &model.Brand{}, &model.GoodsCategoryBrand{}, &model.Banner{}, &model.Goods{})
	if err != nil {
		zap.S().Fatalw("db  AutoMigrate : %s", "err", err.Error())
	}
	zap.S().Infow("init goods db conn success")
}
