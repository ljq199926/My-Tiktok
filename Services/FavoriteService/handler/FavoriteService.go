package handler

import (
	"FavoriteService/model"
	favoriteService "FavoriteService/proto/FavoriteService"
	"context"
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	"time"
)

const Like = 1
const Cancel = 2

type FavoriteService struct{}

func (f FavoriteService) FavoriteAction(ctx context.Context, req *favoriteService.DouyinFavoriteActionRequest, rsp *favoriteService.DouyinFavoriteActionResponse) error {
	token := req.Token
	if token == "" {
		log.Error("token is empty")
		rsp.StatusCode = 1
		rsp.StatusMsg = "token error"
		return nil
	}

	userId, err := model.QueryUserIdByToken(ctx, token)
	if err != nil {
		log.Errorf("redis query error:%s, %d", err.Error(), token)
		rsp.StatusCode = 1
		rsp.StatusMsg = err.Error()
		return nil
	}
	fmt.Printf("FavroiteAction's userId:%d\n", userId)

	videoId := req.VideoId
	actionType := req.ActionType
	queryFavor, row := model.QueryFavor(userId, videoId)

	//点赞
	if actionType == Like {
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
	if actionType == Cancel {
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
	if actionType == Like {
		rsp.StatusMsg = "Like Success"
	} else {
		rsp.StatusMsg = "Cancel Success"
	}
	return nil
}

func (f FavoriteService) FavoriteList(ctx context.Context, req *favoriteService.DouyinFavoriteListRequest, rsp *favoriteService.DouyinFavoriteListResponse) error {
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
	var Videos []*favoriteService.Video
	if row1 > 0 {
		for _, v := range favorList {
			var tmpV favoriteService.Video
			tmpV.Author = &favoriteService.User{
				Id:            0,
				Name:          "",
				FollowCount:   0,
				FollowerCount: 0,
				IsFollow:      false,
			}
			tmpV.Id = v.VideoId
			u, _ := model.QueryVideoByVideoId(v.VideoId)
			tmpV.Author.Id = u.AuthorId
			author, row2 := model.QueryUserByUserId(u.AuthorId)
			if row2 > 0 {
				tmpV.Author.Name = author.Username
				tmpV.Author.FollowCount = author.FollowCount
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
