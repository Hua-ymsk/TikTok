package ffmpeg

import (
	"bytes"
	"fmt"
	"os"
	"tiktok/setting"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func MakeCover(videoName, coverName string) (cover_url string, err error) {
	conf := setting.Conf.VideoConfig
	filePath := conf.PlayUrlPrefix + videoName + ".mp4"
	fmt.Println(filePath)
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		fmt.Println("get cover err:", err)
		return
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		fmt.Println("decode err: ", err)
		return
	}
	cover_url = conf.CoverUrlPrefix + coverName + ".png"
	if err = imaging.Save(img, cover_url); err != nil {
		fmt.Println("save png err:", err)
		return
	}

	return
}
