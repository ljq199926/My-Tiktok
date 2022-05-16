package utils

import (
	"bytes"
	"context"
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
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
		Scope:               Bucket,
		PersistentOps:       "vframe/jpg/offset/1/w/1080/h/1920",
		PersistentNotifyURL: "http://fake.com/qiniu/notify",
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
	path := "%d"
	path = fmt.Sprintf(path, snowFlake.GetId())
	log.Infof("start upload：%s", path)
	go formUploader.Put(context.Background(), &ret, upToken, path, bytes.NewReader(data), dataLen, nil)
	log.Info("ret", ret)
	return path
}
