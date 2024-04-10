package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"ginl/config"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var Db *sql.DB
var GormDb *gorm.DB

func InitDb() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.CustomConfig.Mysql.User, config.CustomConfig.Mysql.Password, config.CustomConfig.Mysql.Hostname, config.CustomConfig.Mysql.Port, config.CustomConfig.Mysql.Database)
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	GormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	err = Db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// 格式化时间，相当于yyyy-MM-dd HH:mm:ss
const localTimeFormat string = "2006-01-02 15:04:05"

// LocalTime 自定义一个类型，本质上是time.Time，但是重写该方法的MarshalJSON方法来改变返回值
type LocalTime time.Time

// MarshalJSON 自定义成序列化JSON内容
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	t2 := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", t2.Format(localTimeFormat))), nil
}

// Value 存储调用
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

// Scan 查询读库操作
func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
