package utils

import (
	"bytes"
	"context"
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"math/rand"
	"time"
)

/*
 @File Name          :UploadToOss.go
 @Author             :cc
 @Version            :1.0.0
 @Date               :2022/5/13 17:16
 @Description        :
 @Function List      :
 @History            :
*/

func UploadQiniu(data []byte) string {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AK, SK)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	//formUploader := storage.NewFormUploader(&cfg)
	formUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}

	//data := []byte("hello, this is qiniu cloud")
	dataLen := int64(len(data))
	rand.Seed(time.Now().UnixNano())
	path := "%d%05d"
	path = fmt.Sprintf(path, time.Now().UnixNano(), rand.Intn(100000))
	log.Infof("start upload：%s", path)
	go formUploader.Put(context.Background(), &ret, upToken, path, bytes.NewReader(data), dataLen, nil)
	return path
}
