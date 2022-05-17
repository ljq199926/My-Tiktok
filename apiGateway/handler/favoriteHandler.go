package handler

import (
	pb "apiGateway/proto/FavoriteService"
	"apiGateway/utils"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FavorAction(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	token := ctx.Query("token")
	videoId := ctx.Query("video_id")
	actionType := ctx.Query("action_type")
	_userid, _ := strconv.Atoi(userId)
	userid := int64(_userid)
	_videoid, _ := strconv.Atoi(videoId)
	videoid := int64(_videoid)
	_actiontype, _ := strconv.Atoi(actionType)
	actiontype := int32(_actiontype)

	microService := utils.InitService()
	favoriteService := pb.NewFavoriteService("go.micro.service.FavoriteService", microService.Client())
	rsp, err := favoriteService.FavoriteAction(context.Background(), &pb.DouyinFavoriteActionRequest{
		UserId:     userid,
		Token:      token,
		VideoId:    videoid,
		ActionType: actiontype,
	})

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, &rsp)
}

func FavorList(ctx *gin.Context) {
	token := ctx.Query("token")

	microService := utils.InitService()
	favoriteService := pb.NewFavoriteService("go.micro.service.FavoriteService", microService.Client())
	rsp, err := favoriteService.FavoriteList(context.Background(), &pb.DouyinFavoriteListRequest{
		Token: token,
	})

	if err != nil {
		ctx.JSON(500, err)
		return
	}
	if (*rsp).VideoList == nil {
		ctx.JSON(200, gin.H{
			"status_code": (*rsp).StatusCode,
			"status_msg":  (*rsp).StatusMsg,
			"video_list":  []*pb.Video{},
		})
	} else {
		ctx.JSON(200, &rsp)
	}
}
