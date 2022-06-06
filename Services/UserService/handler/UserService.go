package handler

import (
	"UserService/model"
	userService "UserService/proto/UserService"
	"UserService/utils"
	"context"
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *UserService) Login(ctx context.Context, req *userService.DouyinUserLoginRequest, rsp *userService.DouyinUserLoginResponse) error {
	var user model.User
	username := req.Username
	password := req.Password
	//检查用户是否不存在
	if exist := model.LoginCheck(&user, username); exist == 1 {
		rsp.StatusCode = 1
		rsp.StatusMsg = "User don't exist"
		return nil
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = "Password wrong"
		return nil
	}
	//生成token
	token, err := utils.CreateToken(user.UserId)
	if err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = "Token Create failed"
		return nil
	}
	//token存redis token为key，用户信息为value
	//c := model.InitRedis()
	//if c != nil {
	//	_, err := c.Do("HMSET", redis.Args{}.Add(token).AddFlat(&user)...)
	//	if err != nil {
	//		fmt.Println("struct err: ", err)
	//		rsp.StatusCode = 1
	//		rsp.StatusMsg = "token save failed"
	//		return nil
	//	}
	//}

	err = model.AddToken(token, &user)
	if err != nil {
		fmt.Println("struct err: ", err)
		rsp.StatusCode = 1
		rsp.StatusMsg = "token save failed"
		return nil
	}

	rsp.StatusCode = 0 //0代表成功其他代表失败
	rsp.StatusMsg = "登录成功！"
	rsp.UserId = user.UserId
	rsp.Token = token
	return nil
}
func (e *UserService) Register(ctx context.Context, req *userService.DouyinUserRegisterRequest, rsp *userService.DouyinUserRegisterResponse) error {
	var user model.User
	username := req.Username
	password := req.Password
	//检查用户是否已经存在
	if exist := model.CheckUser(username); exist == 1 {
		rsp.StatusCode = 1
		rsp.StatusMsg = "User already exist"
		return nil
	}
	//密码非对称加密
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("加密失败:", err)
	}
	encodePWD := string(hash)
	//用户插入数据库
	user.Username = username
	user.Password = encodePWD
	user.FollowCount = 0
	user.FollowerCount = 0
	user.CreateDate = time.Now()
	if ok := model.AddUser(&user); ok == 1 {
		rsp.StatusCode = 1
		rsp.StatusMsg = "User create failed"
		return nil
	}
	id := model.SelecUser(username)
	user.UserId = id
	fmt.Println(id)
	//生成token
	token, err := utils.CreateToken(id)
	if err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = "Token Create failed"
		return nil
	}
	//token存redis token为key，用户信息为value
	//c := model.InitRedis()
	//defer c.Close()
	//if c != nil {
	//	_, err := c.Do("HMSET", redis.Args{}.Add(token).AddFlat(&user)...)
	//	if err != nil {
	//		fmt.Println("struct err: ", err)
	//		rsp.StatusCode = 1
	//		rsp.StatusMsg = "Token save failed"
	//		return nil
	//	}
	//}
	err = model.AddToken(token, &user)
	if err != nil {
		fmt.Println("struct err: ", err)
		rsp.StatusCode = 1
		rsp.StatusMsg = "token save failed"
		return nil
	}
	rsp.StatusCode = 0 //0代表成功其他代表失败
	rsp.StatusMsg = "注册成功！"
	rsp.UserId = id
	rsp.Token = token
	return nil
}
func (e *UserService) Info(ctx context.Context, req *userService.DouyinUserRequest, rsp *userService.DouyinUserResponse) error {
	id := req.UserId
	user := model.GetUserById(id)
	rsp_user := &userService.User{
		Name:          user.Username,
		Id:            user.UserId,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      true,
	}
	rsp.User = rsp_user
	rsp.StatusCode = 0
	rsp.StatusMsg = "Get user's information successfully"

	return nil
}

func (e *UserService) Action(ctx context.Context, req *userService.DouyinRelationActionRequest, rsp *userService.DouyinRelationActionResponse) error {
	req.UserId, _ = model.QueryUserIdByToken(ctx, req.Token)
	log.Info(req.Token)
	log.Info(req.UserId)
	if req.ActionType == 1 {
		model.AddFollower(req.UserId, req.ToUserId)
		rsp.StatusCode = 0
		rsp.StatusMsg = "Follow user successfully"
	} else {
		model.DeleteFollower(req.UserId, req.ToUserId)
		rsp.StatusCode = 0
		rsp.StatusMsg = "Cancell following users successfully"
	}
	return nil
}
func (e *UserService) FollowList(ctx context.Context, req *userService.DouyinRelationFollowListRequest, rsp *userService.DouyinRelationFollowListResponse) error {
	//req.UserId, _ = model.QueryUserIdByToken(ctx, req.Token)
	var anser []*userService.User
	log.Info("enter followlist")
	userList := model.FollowList(req.UserId)
	for _, user := range userList {
		user1 := &userService.User{
			Id:            user.UserId,
			Name:          user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      true,
		}
		anser = append(anser, user1)
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "Get followList successfully"
	rsp.UserList = anser
	return nil
}
func (e *UserService) FollowerList(ctx context.Context, req *userService.DouyinRelationFollowerListRequest, rsp *userService.DouyinRelationFollowerListResponse) error {
	var anser []*userService.User
	log.Info("enter followerlist")
	userList := model.FollowerList(req.UserId)
	myUserId, _ := model.QueryUserIdByToken(ctx, req.Token)
	followList := model.FollowList(myUserId)
	data := make(map[int64]bool)
	for _, follow := range followList {
		data[follow.UserId] = true
	}
	for _, user := range userList {
		user1 := &userService.User{
			Id:            user.UserId,
			Name:          user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      data[user.UserId],
		}
		anser = append(anser, user1)
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "Get followerList successfully"
	rsp.UserList = anser
	return nil
}
