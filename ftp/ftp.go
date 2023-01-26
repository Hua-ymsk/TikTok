package ftp

import (
	"fmt"
	"tiktok/setting"

	"github.com/dutchcoders/goftp"
)

var FTPServer *goftp.FTP

func Init(cfg *setting.FtpConfig) (err error) {
	// 连接ftp服务器
	FTPServer, err = goftp.Connect(cfg.ServerAddr)
	if err != nil {
		panic(fmt.Errorf("connect to ftp server err:%v", err))
	}
	// 登录
	if err = FTPServer.Login(cfg.Name, cfg.Pwd); err != nil {
		panic(fmt.Errorf("login ftp server err:%v", err))
	}


	return
}
