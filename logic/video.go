package logic

import (
	"fmt"
	"log"
	"mime/multipart"
	"strconv"
	"strings"
	"tiktok/common/ffmpeg"
	"tiktok/common/snowflake"
	"tiktok/dao/mysql"
	"tiktok/models"
	"tiktok/setting"
	"tiktok/types"
	"time"

	"github.com/gin-gonic/gin"
)

type VideoLogic struct{}

func NewVideoLogic() *VideoLogic {
	return &VideoLogic{}
}

func (logic *VideoLogic) SaveVideo(c *gin.Context, data *multipart.FileHeader, title string) (err error) {
	// 雪花算法生成视频名
	vidoName := strconv.FormatUint(snowflake.GetID(), 10)
	// 获取类型
	content_type := strings.Split(data.Header["Content-Type"][0], "/")
	video_type := content_type[1]
	// 拼接play_url
	dst := fmt.Sprintf("%s/%s.%s", setting.Conf.PlayUrlPrefix, vidoName, video_type)
	if err := c.SaveUploadedFile(data, dst); err != nil {
		log.Println("save err:", err)
		return err
	}
	// 获取封面截图
	coverName := strconv.FormatUint(snowflake.GetID(), 10)
	cover_url, err := ffmpeg.MakeCover(vidoName, coverName)
	if err != nil {
		log.Println("make cover err: ", err)
		return
	}
	// 保存（200ms）
	video := &models.Video{
		// UserID:    c.GetInt64("user_id"),
		UserID:    1,
		Play_url:  dst,
		Cover_url: cover_url,
		Titile:    title,
		TimeStamp: time.Now().Unix(),
	}
	if err := mysql.SaveVideo(video); err != nil {
		return err
	}
	return
}

func (logic *VideoLogic) VideoList(c *gin.Context, user_id int64) (list []types.Video, err error) {
	// 对官方给出的发布列表响应存疑
	// 不该每个video中都要求author信息，数据冗杂
	// 暂不查询author

	// videos := mysql.PublishList(user_id)
	mysql.ChekFollow(1, 1)
	return
}
