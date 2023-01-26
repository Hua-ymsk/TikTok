package main

import (
	"errors"
	"fmt"
	"testing"
	"tiktok/dao"
)

func TestUserExist(t *testing.T) {
	defer dao.Close()
	if err := dao.Connect(); err != nil {
		fmt.Println(errors.New("connect error"))
		return
	}
	var (
		input   = "1"
		exected = "1"
	)

	actual, _ := dao.UserExist(input)
	if actual != exected {
		t.Errorf("错误")
	}
}
