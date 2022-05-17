package model

import (
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	"time"
)

/*
 @File Name          :Video.go
 @Author             :cc
 @Version            :1.0.0
 @Date               :2022/5/13 15:35
 @Description        :
 @Function List      :
 @History            :
*/

type Video struct {
	Id            int64
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	CreateDate    time.Time
	Title         string
}

type User struct {
	UserId        int64
	Username      string
	Password      string
	FollowCount   int64
	FollowerCount int64
	CreateDate    time.Time
}

type Favorite struct {
	Id         int64
	VideoId    int64
	UserId     int64
	Status     string
	CreateDate time.Time
}

func InsertVideo(data *Video) bool {
	err = db.Create(&data).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func InitVideo(data *Video) {
	data.AuthorId = 1
	data.CoverUrl = ""
	data.PlayUrl = ""
	data.CommentCount = 0
	data.FavoriteCount = 0
	data.CreateDate = time.Now()
	data.Title = ""
}

func QueryVideo(date *string, limit *int) []*Video {
	var VideoList []*Video
	log.Info(*date)
	db.Where("create_date < ?", *date).Order("create_date desc").Find(&VideoList)
	//db.Where("create_date <= ?", date).Order("id desc").Limit(*limit).Find(&VideoList)
	log.Info(VideoList)
	if len(VideoList) <= 30 {
		return VideoList
	}
	return VideoList[0:*limit]
}

func QueryVideoByUserId(userId int64) []*Video {
	var VideoList []*Video
	db.Where("author_id =  ?", userId).Find(&VideoList)
	return VideoList
}

func QueryUserById(Id int64) User {
	var user User
	db.Find(&user, Id)
	return user
}

func IsFavorite(userId int64, VideoId int64) bool {
	var Count int64
	db.Model(&Favorite{}).Where("user_id = ? and video_id and status=?", userId, VideoId, "LIKE").Count(&Count)
	return Count == 1
}
