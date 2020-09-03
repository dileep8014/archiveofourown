package global

import (
	"fmt"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
	"github.com/shyptr/archiveofourown/pkg/setting"
	"github.com/shyptr/archiveofourown/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

var Engine *gorm.DB

func SetupDBEngine() (err error) {
	defer errwrap.Wrap(&err, "init.SetupDBEngine")

	Engine, err = newDBEngine(DatabaseSetting)
	return
}

func newDBEngine(ds *setting.DatabaseSetting) (db *gorm.DB, err error) {
	defer errwrap.Wrap(&err, "database connect")
	// gorm日志配置
	newLogger := logger.New(
		log.New(io.MultiWriter(os.Stdout, &lumberjack.Logger{
			Filename: AppSetting.LogSavePath + "/sql" + AppSetting.LogFileExt,
			MaxSize:  600,
			MaxAge:   10,
		}), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 5 * time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info,     // Log level
			Colorful:      false,           // 禁用彩色打印
		},
	)
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		ds.Username, ds.Password, ds.Host, ds.DBName, ds.Charset, ds.ParseTime)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		return
	}
	DB, err := db.DB()
	if err != nil {
		return
	}
	// 数据库配置
	DB.SetMaxIdleConns(ds.MaxIdConns)
	DB.SetMaxOpenConns(ds.MaxOpenConns)
	// tracer
	tracer.AddGormCallbacks(db)
	return
}
