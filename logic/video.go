package logic

import (
	"fmt"
	"log"
	"mime/multipart"
	"strconv"
	"tiktok/common/ffmpeg"
	"tiktok/common/snowflake"
	"tiktok/dao/mysql"
	"tiktok/models"
	"tiktok/setting"
	"tiktok/types"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type VideoLogic struct{}

func NewVideoLogic() *VideoLogic {
	return &VideoLogic{}
}

func (logic *VideoLogic) Feed(latest_time int64, sender_id int64) (list []types.Video, next_time int64, err error) {
	videos, err := mysql.GetVideosByLatestTime(latest_time)
	if err != nil {
		return nil, time.Now().Unix(), err
	}
	if err := copier.Copy(&list, &videos); err != nil {
		fmt.Println("copy err:", err)
		return nil, time.Now().Unix(), err
	}
	next_time = videos[len(videos)-1].TimeStamp
	// send_id为0，用户未登录，不查点赞信息
	if sender_id != 0 {
		for index, video := range list {
			list[index].IsFavorite, err = mysql.CheckFavorite(sender_id, video.ID)
			if err != nil {
				break
			}
		}
	}

	// 用户信息
	for index, _ := range videos {
		author, err := mysql.GetUserById(videos[index].UserID)
		if err != nil {
			continue
		}
		if err := copier.Copy(&list[index].Author, &author); err != nil {
			continue
		}
		// 是否关注
		if sender_id != 0 {
			list[index].Author.IsFollow, _ = mysql.ChekFollow(sender_id, author.ID)
		}
	}

	return
}

func (logic *VideoLogic) SaveVideo(c *gin.Context, data *multipart.FileHeader, title string) (err error) {
	// 雪花算法生成视频名
	vidoName := strconv.FormatUint(snowflake.GetID(), 10)
	// 获取类型
	// content_type := strings.Split(data.Header["Content-Type"][0], "/")
	// video_type := content_type[1]
	// 拼接play_url
	play_url := fmt.Sprintf("%s/%s.%s", setting.Conf.PlayUrlPrefix, vidoName, "mp4")
	play_dst := fmt.Sprintf("%s/%s.%s", setting.Conf.PlayStaticPrefix, vidoName, "mp4")
	if err := c.SaveUploadedFile(data, play_dst); err != nil {
		log.Println("save err:", err)
		return err
	}
	// 获取封面截图(优化：截图用jpg格式)
	coverName := strconv.FormatUint(snowflake.GetID(), 10)
	cover_url, err := ffmpeg.MakeCover(vidoName, coverName)
	if err != nil {
		log.Println("make covers err: ", err)
		return
	}
	// 保存（200ms）
	video := &models.Video{
		UserID:    c.GetInt64("user_id"),
		PlayURL:   play_url,
		CoverURL:  cover_url,
		Title:     title,
		TimeStamp: time.Now().Unix(),
	}
	if err := mysql.SaveVideo(video); err != nil {
		return err
	}
	return
}

func (logic *VideoLogic) VideoList(user_id, sender_id int64) (list []types.Video, err error) {
	// 对官方给出的发布列表响应存疑
	// 不该每个video中都要求author信息，数据冗杂
	// 暂不查询author

	// 视频信息
	videos, err := mysql.GetPublishList(user_id)
	if err := copier.Copy(&list, &videos); err != nil {
		fmt.Println("copy err:", err)
		return nil, err
	}
	// 发起请求的用户是否点赞
	for index, video := range list {
		list[index].IsFavorite, err = mysql.CheckFavorite(sender_id, video.ID)
		if err != nil {
			break
		}
	}

	// 作者信息(等jack写完)
	// 用户信息
	for index, _ := range videos {
		author, err := mysql.GetUserById(videos[index].UserID)
		if err != nil {
			continue
		}
		if err := copier.Copy(&list[index].Author, &author); err != nil {
			continue
		}
	}

	return
}
