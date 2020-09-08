package global

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
	"github.com/shyptr/archiveofourown/pkg/gormplugins"
	"github.com/shyptr/archiveofourown/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
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
	newLogger := gorm.Logger{
		LogWriter: log.New(io.MultiWriter(os.Stdout, &lumberjack.Logger{
			Filename: AppSetting.LogSavePath + "/sql" + AppSetting.LogFileExt,
			MaxSize:  600,
			MaxAge:   10,
		}), "\r\n", 0), // io writer
		//logger.Config{
		//	SlowThreshold: 5 * time.Second, // 慢 SQL 阈值
		//	LogLevel:      logger.Info,     // Log level
		//	Colorful:      false,           // 禁用彩色打印
		//},
	}
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		ds.Username, ds.Password, ds.Host, ds.DBName, ds.Charset, ds.ParseTime)
	db, err = gorm.Open(ds.DBType, dsn)
	if err != nil {
		return
	}

	// 数据库配置
	db.SetLogger(newLogger)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(ds.MaxIdConns)
	db.DB().SetMaxOpenConns(ds.MaxOpenConns)
	// tracer
	gormplugins.AddGormCallbacks(db)
	gormplugins.AddModelFieldToGorm(db)
	return
}
