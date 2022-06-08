package handler

import (
	"VideoService/model"
	"VideoService/utils"
	"context"
	"fmt"
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
	var err error
	v.AuthorId, err = model.QueryUserIdByToken(c, token)
	if err != nil {
		log.Errorf("redis query error:%s, %d", err.Error(), token)
		rsp.StatusCode = 1
		rsp.StatusMsg = err.Error()
		return nil
	}

	log.Info(len(data))
	//go utils.UploadQiniu(data)
	res := utils.UploadQiniu(data)
	v.PlayUrl = "http://rbtdate4z.hn-bkt.clouddn.com/" + res
	v.CoverUrl = "http://rbtdate4z.hn-bkt.clouddn.com/cover/" + res
	v.Title = req.Title
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
		//log.Info(v.CreateDate.Unix(), LatestTime)
		if v.CreateDate.Unix() < LatestTime {
			LatestTime = v.CreateDate.Unix()
		}
		var tmpV videoService.Video
		u := model.QueryUserById(v.AuthorId)
		tmpV.Author = &videoService.User{
			Id:            u.UserId,
			Name:          u.Username,
			Fol1OwCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      false,
		}
		tmpV.Id = v.Id
		tmpV.CoverUrl = v.CoverUrl
		tmpV.PlayUrl = v.PlayUrl
		tmpV.IsFavorite = model.IsFavorite(u.UserId, v.Id)
		tmpV.CommentCount = v.CommentCount
		tmpV.FavoriteCount = v.FavoriteCount
		tmpV.Title = v.Title

		Videos = append(Videos, &tmpV)
	}
	return LatestTime * 1000, Videos
}

func (video *VideoService) Feed(c context.Context, req *videoService.DouyinFeedRequest, rep *videoService.DouyinFeedResponse) error {
	LatestTime := req.LatestTime
	format := "2006-01-02 15:04:05"
	t := time.Unix(LatestTime/1000, 0)
	date := t.Format(format)
	log.Info(date)

	rep.StatusCode = 0
	rep.StatusMsg = "success"
	rep.NextTime, rep.VideoList = PaserModel(date, nil)
	log.Info(len(rep.VideoList), rep.VideoList)
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
	userId, err := model.QueryUserIdByToken(c, token)
	if err != nil {
		log.Errorf("redis query error:%s, %s", err.Error(), token)
		rep.StatusCode = 1
		rep.StatusMsg = err.Error()
		return nil
	}
	videoList := model.QueryVideoByUserId(userId)
	_, rep.VideoList = PaserModel("", videoList)
	log.Info("rep.VideoList:", rep.VideoList)
	rep.StatusCode = 0
	rep.StatusMsg = "success"
	return nil
}

func (f *VideoService) FavoriteAction(ctx context.Context, req *videoService.DouyinFavoriteActionRequest, rsp *videoService.DouyinFavoriteActionResponse) error {
	token := req.Token
	log.Info("FavoriteAction in videoserive")
	if token == "" {
		log.Error("token is empty")
		rsp.StatusCode = 1
		rsp.StatusMsg = "token error"
		return nil
	}

	userId, err := model.QueryUserIdByToken(ctx, token)
	if err != nil {
		log.Errorf("redis query error:%v, %v", err, token)
		rsp.StatusCode = 1
		rsp.StatusMsg = err.Error()
		return nil
	}
	fmt.Printf("FavroiteAction's userId:%d\n", userId)

	videoId := req.VideoId
	actionType := req.ActionType
	queryFavor, row := model.QueryFavor(userId, videoId)

	//点赞
	if actionType == utils.Like {
		//首次点赞插入记录，非首次点赞则修改status
		if row == 0 {
			var favor model.Favorite
			favor.UserId = userId
			favor.VideoId = videoId
			favor.CreateDate = time.Now()
			favor.Status = "LIKE"
			result := model.InsertFavorite(&favor)
			if result == false {
				rsp.StatusCode = 1
				rsp.StatusMsg = "LIKE Failed"
				return nil
			}
			result = model.UpdateVideoByVideoId(videoId, actionType)
			if result == false {
				rsp.StatusCode = 1
				rsp.StatusMsg = "LIKE Failed"
				return nil
			}
		} else {
			result := model.UpdateFavor(&queryFavor, "LIKE")
			if result == false {
				rsp.StatusCode = 1
				rsp.StatusMsg = "LIKE Failed"
				return nil
			}
			result = model.UpdateVideoByVideoId(videoId, actionType)
			if result == false {
				rsp.StatusCode = 1
				rsp.StatusMsg = "LIKE Failed"
				return nil
			}
		}
	}

	//取消点赞
	if actionType == utils.Cancel {
		result := model.UpdateFavor(&queryFavor, "Cancel")
		if result == false {
			rsp.StatusCode = 1
			rsp.StatusMsg = "Cancel failed"
			return nil
		}
		result = model.UpdateVideoByVideoId(videoId, actionType)
		if result == false {
			rsp.StatusCode = 1
			rsp.StatusMsg = "Cancel Failed"
			return nil
		}
	}

	rsp.StatusCode = 0
	if actionType == utils.Like {
		rsp.StatusMsg = "Like Success"
	} else {
		rsp.StatusMsg = "Cancel Success"
	}
	return nil
}

func (f *VideoService) FavoriteList(ctx context.Context, req *videoService.DouyinFavoriteListRequest, rsp *videoService.DouyinFavoriteListResponse) error {
	log.Info("FavoriteList started!")
	token := req.Token
	if token == "" {
		log.Error("token is empty")
		rsp.StatusCode = 1
		rsp.StatusMsg = "token error"
		return nil
	}

	userId, err := model.QueryUserIdByToken(ctx, token)
	if err != nil {
		log.Errorf("redis query error:%s, %s", err.Error(), token)
		rsp.StatusCode = 1
		rsp.StatusMsg = err.Error()
		return nil
	}

	fmt.Printf("FavorList's userId:%d\n", userId)
	favorList, row1 := model.QueryFavorByUserId(userId)
	var Videos []*videoService.Video
	if row1 > 0 {
		for _, v := range favorList {
			var tmpV videoService.Video
			tmpV.Author = &videoService.User{
				Id:            0,
				Name:          "",
				Fol1OwCount:   0,
				FollowerCount: 0,
				IsFollow:      false,
			}
			tmpV.Id = v.VideoId
			u, _ := model.QueryVideoByVideoId(v.VideoId)
			tmpV.Author.Id = u.AuthorId
			author, row2 := model.QueryUserByUserId(u.AuthorId)
			if row2 > 0 {
				tmpV.Author.Name = author.Username
				tmpV.Author.Fol1OwCount = author.FollowCount
				tmpV.Author.FollowerCount = author.FollowerCount
				//此用户是否关注该作者待后面关注列表完善后再修改
				tmpV.Author.IsFollow = false
			}
			tmpV.PlayUrl = u.PlayUrl
			tmpV.CoverUrl = u.CoverUrl
			tmpV.FavoriteCount = u.FavoriteCount
			tmpV.CommentCount = u.CommentCount
			tmpV.IsFavorite = true

			Videos = append(Videos, &tmpV)
		}
	}

	rsp.VideoList = Videos
	log.Info("rsp.VideoList:", rsp.VideoList)
	rsp.StatusCode = 0
	rsp.StatusMsg = "Get FavoriteList Success"

	return nil
}

func (video *VideoService) CommentAction(c context.Context, req *videoService.DouyinCommentActionRequest, rep *videoService.DouyinCommentActionResponse) error {
	userId, err := model.QueryUserIdByToken(c, req.Token)
	if err != nil {
		rep.StatusCode = -1
		rep.StatusMsg = "token error"
		return nil
	}
	videoId := req.VideoId
	actionType := req.ActionType

	log.Infof("userid: %v, videoId: %v, actionType: %v", userId, videoId, actionType)
	if actionType == 1 { //发布评论
		commentText := req.CommentText
		user := model.QueryUserById(userId)
		log.Info(user, userId)
		if user.UserId == -1 {
			rep.StatusCode = -1
			rep.StatusMsg = "userid error"
			return nil
		}
		if cnt := model.CountVideoByVideoId(videoId); cnt == 0 {
			rep.StatusCode = -1
			rep.StatusMsg = "video_id error"
			return nil
		}
		if check := model.InsertComment(userId, videoId, commentText); !check {
			rep.StatusCode = -1
			rep.StatusMsg = "insert error"
			return nil
		}
		model.UpdateCommentCount(videoId, 1)
		var comment videoService.Comment
		comment.User = &videoService.User{
			Id:            user.UserId,
			Name:          user.Username,
			Fol1OwCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      false,
		}
		comment.Content = commentText
		comment.CreateDate = time.Now().Format("01-02")
		rep.StatusCode = 0
		rep.Comment = &comment
		rep.StatusMsg = "add success"
	} else if actionType == 2 {
		commentId := req.CommentId

		if check := model.DeleteCommentByCommentId(commentId); !check {
			rep.StatusCode = -1
			rep.StatusMsg = "delete error"
			return nil
		}
		model.UpdateCommentCount(videoId, -1)
		rep.StatusCode = 0
		rep.StatusMsg = "delete success"
	} else {
		rep.StatusCode = -1
		rep.StatusMsg = "actionType error"
	}
	return nil
}

func (video *VideoService) CommentList(c context.Context, req *videoService.DouyinCommentListRequest, rep *videoService.DouyinCommentListResponse) error {
	videoId := req.VideoId

	//token := req.Token
	//userId, err := model.QueryUserIdByToken(c, token)
	//if err != nil {
	//	rep.StatusCode = -1
	//	rep.StatusMsg = "token err"
	//	return nil
	//}

	commentList := model.QueryCommentListByVideoId(videoId)
	var tmpCommentList []*videoService.Comment

	for _, comment := range commentList {
		user := model.QueryUserById(comment.UserId)
		if user.UserId == -1 {
			rep.StatusCode = -1
			rep.StatusMsg = "userid err"
			return nil
		}
		tmpComment := &videoService.Comment{
			Id: comment.Id,
			User: &videoService.User{
				Id:            user.UserId,
				Name:          user.Username,
				Fol1OwCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      false,
			},
			Content:    comment.Content,
			CreateDate: comment.CreateDate.Format("01-02"),
		}
		tmpCommentList = append(tmpCommentList, tmpComment)
	}
	rep.StatusCode = 0
	rep.StatusMsg = "success"
	rep.CommentList = tmpCommentList
	return nil
}
