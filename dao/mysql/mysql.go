package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"goginProject/setting"
)

var db *sqlx.DB

func Init(cfg *setting.Mysql) (err error) {
	//数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	//校验账号是否正确
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return
	}

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	return
}

func Close() {
	_ = db.Close()
}
