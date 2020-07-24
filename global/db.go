package global

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shyptr/archiveofourown/pkg/errwrap"
	"github.com/shyptr/archiveofourown/pkg/setting"
)

var DBEngine *sql.DB

func SetupDBEngine() (err error) {
	defer errwrap.Wrap(&err, "init.SetupDBEngine")

	DBEngine, err = newDBEngine(DatabaseSetting)
	return
}

func newDBEngine(ds *setting.DatabaseSetting) (db *sql.DB, err error) {
	defer errwrap.Wrap(&err, "database connect")

	db, err = sql.Open(ds.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		ds.Username, ds.Password, ds.Host, ds.DBName, ds.Charset, ds.ParseTime))
	return
}
