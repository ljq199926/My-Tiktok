package handler

import (
	pb "apiGateway/proto/VideoService"
	"apiGateway/utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client/grpc"
	log "github.com/micro/go-micro/v2/logger"
	"mime/multipart"
	"strconv"
	"time"
)

/*
 @File Name          :videoHandler.go
 @Author             :cc
 @Version            :1.0.0
 @Date               :2022/5/14 14:18
 @Description        :
 @Function List      :
 @History            :
*/

func UploadVideo(c *gin.Context) {
	log.Info("UploadVideo method")
	token := c.PostForm("token")
	title := c.PostForm("title")
	data, errr := c.FormFile("data")
	log.Infof("data_size：%d %s %s", data.Size, token, title)
	if errr != nil {
		log.Errorf("failed in read file %s", errr.Error())
		c.JSON(200, pb.DouyinPublishActionResponse{StatusMsg: errr.Error(), StatusCode: 1})
		return
	}
	openFile, e := data.Open()
	if e != nil {
		log.Error(e)
		return
	}
	defer func(openFile multipart.File) {
		err := openFile.Close()
		if err != nil {
			log.Error(err)
		}
	}(openFile)
	Bdata := make([]byte, data.Size)
	count, err := openFile.Read(Bdata) //读取传入文件的内容
	if err != nil {
		log.Infof("failed in open file %s, %d", err.Error(), count)
		c.JSON(200, pb.DouyinPublishActionResponse{StatusMsg: err.Error(), StatusCode: 1})
		return
	}

	microService := utils.InitService()
	microService.Client().Init(grpc.MaxSendMsgSize(1024 * 1024 * 1024))
	videoService := pb.NewVideoService("go.micro.service.VideoService", microService.Client())
	rsp, err := videoService.PublishAction(context.Background(), &pb.DouyinPublishActionRequest{
		Token: token,
		Data:  Bdata,
		Title: title,
	})
	if err != nil {
		log.Error("upload failed")
	}
	c.JSON(200, &rsp)
}

func Feed(c *gin.Context) {
	timeStamp := c.Query("latest_time")
	var latestTime int64
	if timeStamp == "" {
		latestTime = time.Now().Unix() * 1000
	} else {
		latestTime, _ = strconv.ParseInt(c.Query("latest_time"), 10, 64)
	}

	log.Info("latest_time:", latestTime)
	microService := utils.InitService()
	//microService.Client().Init(grpc.MaxSendMsgSize(1024 * 1024 * 1024))
	videoService := pb.NewVideoService("go.micro.service.VideoService", microService.Client())
	rsp, err := videoService.Feed(context.Background(), &pb.DouyinFeedRequest{
		LatestTime: latestTime,
	})
	if err != nil {
		log.Error(err)
	}
	c.JSON(200, &rsp)
}

func PublishList(c *gin.Context) {
	token := c.Query("token")
	microService := utils.InitService()
	//microService.Client().Init(grpc.MaxSendMsgSize(1024 * 1024 * 1024))
	videoService := pb.NewVideoService("go.micro.service.VideoService", microService.Client())
	rsp, err := videoService.PublishList(context.Background(), &pb.DouyinPublishListRequest{
		Token: token,
	})

	if err != nil {
		log.Error(err)
	}
	//err = c.ShouldBindJSON(rsp)
	//if err != nil {
	//	log.Error(err)
	//}
	//c.JSON(200, &rsp)
	if rsp == nil {
		c.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  "error call",
			"video_list":  []*pb.Video{},
		})
	} else if (*rsp).VideoList == nil {
		c.JSON(200, gin.H{
			"status_code": (*rsp).StatusCode,
			"status_msg":  (*rsp).StatusMsg,
			"video_list":  []*pb.Video{},
		})
	} else {
		c.JSON(200, &rsp)
	}
}

func CommentAction(c *gin.Context) {
	//userId := c.Query("user_id")
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")
	commentId := c.Query("comment_id")
	//log.Info(userId, videoId, actionType, commentText, commentId, token)
	log.Infof(" videoid: %v, actiontype: %v, commenttext: %v, commentid: %v, token:%v",
		videoId, actionType, commentText, commentId, token)
	//_userId, _ := strconv.ParseInt(userId, 10, 64)
	_videoId, _ := strconv.ParseInt(videoId, 10, 64)
	_commentId, _ := strconv.ParseInt(commentId, 10, 64)
	_actionType, _ := strconv.Atoi(actionType)
	log.Infof(" _videoId: %v, _actionType: %v, _commentId: %v",
		_videoId, _actionType, _commentId)
	microService := utils.InitService()
	videoService := pb.NewVideoService("go.micro.service.VideoService", microService.Client())
	rsp, err := videoService.CommentAction(context.Background(), &pb.DouyinCommentActionRequest{
		Token: token,
		//UserId:      _userId,
		VideoId:     _videoId,
		ActionType:  int32(_actionType),
		CommentId:   _commentId,
		CommentText: commentText,
	})

	if err != nil {
		log.Error(err)
	}

	c.JSON(200, &rsp)
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")
	_videoId, _ := strconv.ParseInt(videoId, 10, 64)
	microService := utils.InitService()
	videoService := pb.NewVideoService("go.micro.service.VideoService", microService.Client())
	rsp, err := videoService.CommentList(context.Background(), &pb.DouyinCommentListRequest{
		Token:   token,
		VideoId: _videoId,
	})

	if err != nil {
		log.Error(err)
	}

	c.JSON(200, &rsp)
}
