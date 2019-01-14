package DistributedLock

import (
	"core/db"
	"core/util"
	"fmt"
	"time"
)

//分布式锁设置
//1-加锁 原子操作 set超时时间
//2-逻辑判断
//3-判断锁是否是这个进程是否自己的数据
//4-删除锁

const distributedLock = "DistributedLock:%s"

type RedisLock struct {
}

var DefaultRedisLock RedisLock

//锁时间为5s
func (RedisLock) AddLock() (string, error) {
	nonStr := util.RandomNonceStr(9)
	//分布式锁超时时间设定
	lockTime := time.Duration(5) * time.Second
	LockKey := fmt.Sprintf(distributedLock, nonStr)
	isLock, err := db.RedisClient.SetNX(LockKey, nonStr, lockTime).Result()
	if err != nil || !isLock {
		//日志记录 上锁失败
	}
	return nonStr, err
}

//解锁
func (RedisLock) DelLock(nonStr string) error {
	//Lua 脚本代码
	scriptStr := "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end"
	LockKey := fmt.Sprintf(distributedLock, nonStr)
	strList := []string{LockKey}
	return db.RedisClient.Eval(scriptStr, strList).Err()
}
