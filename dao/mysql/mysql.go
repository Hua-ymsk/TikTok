// gorm

package mysql

import (
	"fmt"
	"tiktok/setting"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(cfg *setting.MysqlConfig) (err error){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",cfg.User,cfg.Pwd,cfg.Host,cfg.Port,cfg.DB)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式（MySQL 5.7之后）
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	}), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("connect to mysql err:%+v", err))
	}

	return
}