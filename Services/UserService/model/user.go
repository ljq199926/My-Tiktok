package model

import (
	log "github.com/micro/go-micro/v2/logger"
	"time"
)

type User struct {
	UserId        int64
	Username      string
	Password      string
	FollowCount   int64
	FollowerCount int64
	CreateDate    time.Time
}
type Follower struct {
	Id         int64
	UserId     int64
	ToUserId   int64
	CreateDate time.Time
}

func CheckUser(username string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.UserId > 0 {
		return 1 //1代表用户已存在
	}
	return 0 //0代表用户不存在
}
func AddUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return 1
	}
	return 0
}
func LoginCheck(data *User, username string) int {
	db.Where("username = ?", username).First(&data)

	if data.UserId <= 0 {
		return 1 //1代表用户不存在
	}
	return 0
}
func SelecUser(username string) int64 {
	var user User
	db.Where("username = ?", username).First(&user)
	return user.UserId
}
func GetUserById(userid int64) User {
	var user User
	db.Where("user_id = ?", userid).First(&user)
	return user
}
func AddFollower(UserId int64, ToUserId int64) int {
	var follower Follower
	follower.CreateDate = time.Now()
	follower.ToUserId = ToUserId
	follower.UserId = UserId
	err := db.Create(&follower).Error
	//var result *gorm.DB
	var user1 User
	db.Where("user_id = ?", UserId).First(&user1)
	log.Info(user1.UserId)
	log.Info(user1.FollowerCount)
	err = db.Model(&user1).Where("user_id=?", user1.UserId).Update("follow_count", user1.FollowCount+1).Error
	log.Info(err)
	//db.Model(&user1).Update("follow_count", user1.FollowCount+1)
	var user2 User
	db.Where("user_id = ?", ToUserId).First(&user2)
	log.Info(user2.Username)
	db.Model(&user2).Where("user_id=?", user2.UserId).Update("follower_count", user2.FollowerCount+1)

	if err != nil {
		return 1
	}
	return 0
}
func DeleteFollower(UserId int64, ToUserId int64) bool {
	db.Where("user_id=? and to_user_id=?", UserId, ToUserId).Delete(Follower{})
	var user1 User
	db.Where("user_id = ?", UserId).First(&user1)
	db.Model(&user1).Where("user_id=?", user1.UserId).Update("follow_count", user1.FollowCount-1)
	var user2 User
	db.Where("user_id = ?", ToUserId).First(&user2)
	db.Model(&user2).Where("user_id=?", user2.UserId).Update("follower_count", user2.FollowerCount-1)
	return true
}
func FollowList(UserId int64) []User {
	var userList []User
	err := db.Table("user").Select("user.user_id as UserId, username as Username, password as Password, follow_count as FollowCount, follower_count as FollowerCount, user.create_date as CreateDate").Joins("left join follower on user.user_id=follower.to_user_id").Where("follower.user_id=?", UserId).Scan(&userList).Error
	log.Info(err)
	return userList
}
func FollowerList(UserId int64) []User {
	var userList []User
	err := db.Table("user").Select("user.user_id as UserId, username as Username, password as Password, follow_count as FollowCount, follower_count as FollowerCount, user.create_date as CreateDate").Joins("left join follower on user.user_id=follower.user_id").Where("follower.to_user_id=?", UserId).Scan(&userList).Error
	log.Info(err)
	return userList
}
