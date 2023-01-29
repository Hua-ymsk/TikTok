package snowflake

import (
	"tiktok/setting"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake // 实例
	sonyMachineID uint16               // 机器ID
)

func getMachineID() (uint16, error) { // 返回全局定义的机器ID
	return sonyMachineID, nil
}

// 需传入当前的机器ID
func Init() (err error) {
	sonyMachineID = uint16(setting.Conf.MachineID)
	t, _ := time.Parse("2006-01-02", setting.Conf.StartTime) // 初始化一个开始的时间
	settings := sonyflake.Settings{                          // 生成全局配置
		StartTime: t,
		MachineID: getMachineID, // 指定机器ID
	}
	sonyFlake = sonyflake.NewSonyflake(settings) // 用配置生成sonyflake节点
	return
}

// GetID 返回生成的id值
func GetID() (id uint64) { // 拿到sonyflake节点生成id值
	id, _ = sonyFlake.NextID()
	return
}
