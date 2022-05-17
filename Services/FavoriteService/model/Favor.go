package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	Id            int64
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	CreateDate    time.Time
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

const Like = 1
const Cancel = 2

func UpdateVideoByVideoId(videoId int64, actionType int32) bool {
	var video Video
	var result *gorm.DB
	db.Where("id = ?", videoId).First(&video)
	if actionType == Like {
		result = db.Model(&video).Update("favorite_count", video.FavoriteCount+1)
	} else if actionType == Cancel {
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

func QueryUserByUserId(userId int64) (User, int64) {
	var user User
	find := db.Where("user_id = ?", userId).First(&user)
	return user, find.RowsAffected
}

func InsertFavorite(data *Favorite) bool {
	err = db.Create(&data).Error
	if err != nil {
		fmt.Println(err)
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
