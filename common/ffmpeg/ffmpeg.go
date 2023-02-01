package ffmpeg

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"tiktok/setting"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func MakeCover(videoName, coverName string) (cover_url string, err error) {
	filePath := fmt.Sprintf("%s/%s.%s", setting.Conf.PlayStaticPrefix, videoName, "mp4")
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Println("get cover err:", err)
		return
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		log.Println("decode err: ", err)
		return
	}
	cover_url = fmt.Sprintf("%s/%s.%s", setting.Conf.CoverUrlPrefix, coverName, "jpg")
	cover_dst := fmt.Sprintf("%s/%s.%s", setting.Conf.CoverStaticPrefix, coverName, "jpg")
	if err = imaging.Save(img, cover_dst); err != nil {
		log.Println("save png err:", err)
		return
	}
	return
}

func ExecCover(videoName, coverName string) (cover_url string, err error) {
	inputFile := "E:\\Project\\GoProject\\tiktok-static\\static\\videos\\" + videoName + ".mp4"
	outputFile := "E:\\Project\\GoProject\\tiktok-static\\static\\covers\\" + coverName + ".png"

	args := []string{"-ss", "0.1", "-i", inputFile, "-vframes", "1", outputFile}
	cmd := exec.Command("ffmpeg", args...)
	if err := cmd.Run(); err != nil {
		log.Println("cmd run err:", err)
	}

	return
}
