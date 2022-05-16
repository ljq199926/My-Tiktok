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
	data, errr := c.FormFile("data")
	log.Infof("data_size：%d", data.Size)
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
		latestTime = time.Now().Unix()
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

	if (*rsp).VideoList == nil {
		c.JSON(200, gin.H{
			"status_code": (*rsp).StatusCode,
			"status_msg":  (*rsp).StatusMsg,
			"video_list":  []*pb.Video{},
		})
	} else {
		c.JSON(200, &rsp)
	}
}
