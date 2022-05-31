package model

import (
	"VideoService/utils"
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	"gorm.io/gorm"
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

type Comment struct {
	Id         int64
	UserId     int64
	VideoId    int64
	Content    string
	CreateDate time.Time
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
	user.UserId = -1
	err := db.Find(&user, Id).Error
	if err != nil {
		log.Error(err)
		return user
	}
	return user
}

func IsFavorite(userId int64, VideoId int64) bool {
	var Count int64
	db.Model(&Favorite{}).Where("user_id = ? and video_id and status=?", userId, VideoId, "LIKE").Count(&Count)
	return Count == 1
}

func UpdateVideoByVideoId(videoId int64, actionType int32) bool {
	var video Video
	var result *gorm.DB
	db.Where("id = ?", videoId).First(&video)
	if actionType == utils.Like {
		result = db.Model(&video).Update("favorite_count", video.FavoriteCount+1)
	} else if actionType == utils.Cancel {
		result = db.Model(&video).Update("favorite_count", video.FavoriteCount-1)
	}
	return result.RowsAffected == 1
}

func QueryFavorByUserId(userId int64) ([]*Favorite, int64) {
	var favor []*Favorite
	find := db.Where("user_id = ? and status = ?", userId, "LIKE").Find(&favor)
	return favor, find.RowsAffected
}

func QueryVideoByVideoId(videoId int64) (Video, int64) {
	var video Video
	find := db.Where("id = ?", videoId).First(&video)
	return video, find.RowsAffected
}

func CountVideoByVideoId(videoId int64) int64 {
	var count int64
	err := db.Model(&Video{}).Where("id=?", videoId).Count(&count).Error
	if err != nil {
		log.Error(err)
		return 0
	}
	return count
}

func QueryUserByUserId(userId int64) (User, int64) {
	var user User
	find := db.Where("user_id = ?", userId).First(&user)
	return user, find.RowsAffected
}

func InsertFavorite(data *Favorite) bool {
	err = db.Create(&data).Error
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func QueryFavor(userId int64, VideoId int64) (Favorite, int64) {
	var favor Favorite
	result := db.Where("user_id = ? and video_id = ?", userId, VideoId).First(&favor)
	return favor, result.RowsAffected
}

func UpdateFavor(favor *Favorite, status string) bool {
	result := db.Model(&favor).Update("status", status)
	return result.RowsAffected == 1
}

func InsertComment(userId int64, videoId int64, commentText string) bool {
	var comment Comment
	comment.CreateDate = time.Now()
	comment.Content = commentText
	comment.UserId = userId
	comment.VideoId = videoId
	err := db.Create(&comment).Error
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func DeleteCommentByCommentId(commentId int64) bool {
	err := db.Where("id=?", commentId).Delete(Comment{}).Error
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func QueryCommentListByVideoId(videoId int64) []Comment {
	var commentList []Comment
	err := db.Where("video_id=?", videoId).Order("create_date desc").Find(&commentList).Error
	if err != nil {
		return nil
	}
	return commentList
}
