package handler

import (
	"VideoService/model"
	"VideoService/utils"
	"context"
	"github.com/gomodule/redigo/redis"
	log "github.com/micro/go-micro/v2/logger"
	"time"
)
import videoService "VideoService/proto/VideoService"

type VideoService struct{}

func (video *VideoService) PublishAction(c context.Context, req *videoService.DouyinPublishActionRequest, rsp *videoService.DouyinPublishActionResponse) error {
	var token string
	log.Info("PublishAction called")
	var v model.Video
	model.InitVideo(&v)

	var data []byte
	log.Infof("stream_size：%d", len(req.Data))
	token = req.Token
	data = req.Data

	r := model.InitRedis()
	if r != nil {
		userId, err := redis.Int64(r.Do("HGET", token, "UserId"))
		if err != nil {
			log.Errorf("redis query error:%s, %d", err.Error(), token)
			rsp.StatusCode = 1
			rsp.StatusMsg = "token query failed"
			return nil
		}
		v.AuthorId = userId
	} else {
		log.Error("init redis error")
		rsp.StatusCode = 1
		rsp.StatusMsg = "init redis error"
		return nil
	}

	log.Info(len(data))
	//go utils.UploadQiniu(data)
	res := utils.UploadQiniu(data)
	//res := ""
	//if res == "" {
	//	rsp.StatusCode = 1
	//	rsp.StatusMsg = "upload failed"
	//	log.Error("upload failed")
	//	return nil
	//}
	v.PlayUrl = "http://rbtdate4z.hn-bkt.clouddn.com/" + res
	model.InsertVideo(&v)

	rsp.StatusCode = 0
	rsp.StatusMsg = "upload success"
	return nil
}

func PaserModel(date string, video []*model.Video) (int64, []*videoService.Video) {
	var VideoList []*model.Video
	if date == "" && video != nil {
		VideoList = video
	} else if date != "" && video == nil {
		VideoList = model.QueryVideo(&date, &utils.VideoNumLimit)
	}
	var LatestTime = time.Now().Unix()
	var Videos []*videoService.Video
	for _, v := range VideoList {
		if v.CreateDate.Unix() < LatestTime {
			LatestTime = v.CreateDate.Unix()
		}
		var tmpV videoService.Video
		tmpV.Author = &videoService.User{
			Id:            0,
			Name:          "",
			Fol1OwCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}
		tmpV.Id = v.Id
		tmpV.CoverUrl = v.CoverUrl
		tmpV.PlayUrl = v.PlayUrl
		tmpV.IsFavorite = false
		tmpV.CommentCount = v.CommentCount
		tmpV.FavoriteCount = v.FavoriteCount
		u := model.QueryUserById(v.AuthorId)
		tmpV.Author.Id = u.UserId
		tmpV.Author.Name = u.Username
		tmpV.Author.FollowerCount = u.FollowerCount
		tmpV.Author.Fol1OwCount = u.FollowCount
		tmpV.Author.IsFollow = model.IsFavorite(u.UserId, v.Id)

		Videos = append(Videos, &tmpV)
	}
	return LatestTime, Videos
}

func (video *VideoService) Feed(c context.Context, req *videoService.DouyinFeedRequest, rep *videoService.DouyinFeedResponse) error {
	LatestTime := req.LatestTime
	format := "2006-01-02 15:04:05"
	t := time.Unix(LatestTime, 0)
	date := t.Format(format)
	log.Info(date)

	rep.StatusCode = 0
	rep.StatusMsg = "success"
	rep.NextTime, rep.VideoList = PaserModel(date, nil)
	return nil
}

func (video *VideoService) PublishList(c context.Context, req *videoService.DouyinPublishListRequest, rep *videoService.DouyinPublishListResponse) error {
	token := req.Token
	if token == "" {
		log.Error("token is empty")
		rep.StatusCode = 1
		rep.StatusMsg = "token error"
		rep.VideoList = nil
		return nil
	}

	r := model.InitRedis()
	var userId int64
	if r != nil {
		uId, err := redis.Int64(r.Do("HGET", token, "UserId"))
		if err != nil {
			log.Errorf("redis query error:%s, %d", err.Error(), token)
			rep.StatusCode = 1
			rep.StatusMsg = "token query failed"
			return nil
		}
		userId = uId
	} else {
		log.Error("init redis error")
		rep.StatusCode = 1
		rep.StatusMsg = "init redis error"
		return nil
	}

	videoList := model.QueryVideoByUserId(userId)
	_, rep.VideoList = PaserModel("", videoList)
	rep.StatusCode = 0
	rep.StatusMsg = "success"
	return nil
}
