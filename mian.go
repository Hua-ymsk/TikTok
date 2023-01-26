package main

import (
	"errors"
	"fmt"
	"tiktok/dao"
	"tiktok/dao/mysql"
	"tiktok/middleware"
	"tiktok/router"
	"tiktok/setting"
)

func main() {
	// 加载config
	if err := setting.Init(); err != nil {
		fmt.Println("load config err:", err)
	}
	// logger初始化
	if err := middleware.InitLogger(); err != nil {
		fmt.Println("init logger err:", err)
		return
	}
	// 原生sql
	if err := dao.Connect(); err != nil {
		fmt.Println(errors.New("connect error"))
		return
	}
	defer dao.Close()
	// gorm初始化
	if err := mysql.Init(setting.Conf.MysqlConfig); err != nil {
		fmt.Println("init mysql err:", err)
		return
	}
	// router初始化
	r := router.InitRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Println("init router err:", err)
		return
	}

	// r := gin.Default()
	// router.InitRouter()
	// r.Run(":9090") // listen and serve on 0.0.0.0:9090
}
