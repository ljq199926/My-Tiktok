package utils

import (
	"errors"
	"sync"
	"time"
)

/*
 @File Name          :SnowFlake.go
 @Author             :cc
 @Version            :1.0.0
 @Date               :2022/5/16 10:35
 @Description        :
 @Function List      :
 @History            :
*/

const (
	workerBits uint8 = 10
	// 工作机器id位数
	numberBits uint8 = 12
	// 序列号位数
	workerMax int64 = -1 ^ (-1 << workerBits)
	// 工作机器id最大值
	numberMax int64 = -1 ^ (-1 << numberBits)
	// 序列号最大值
	timeShift uint8 = workerBits + numberBits
	// 时间左偏移量
	workerShift uint8 = numberBits
	// 工作机器id左偏移量
	startTime int64 = 1525705533000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
	// 开始时间戳
)

var snowFlake *Worker

func init() {
	snowFlake, _ = NewWorker(2)
}

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID excess of quantity")
	}
	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *Worker) GetId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				// 这里当段时间请求的id过多时可能会发生短暂的空循环
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	ID := int64((now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number))
	return ID
}
