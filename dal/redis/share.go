package redis

import (
	"fmt"
	"time"
)

// 分享码缓存在Redis中

const (
	set           = "object_storage|share"
	sharingFormat = "share_id:%s|pwd:%s" // share_id:password, password为空则不需要密码
)

// CheckShare 检查分享码, 返回分享码对应的文件uuid
func CheckShare(shareID, password string) (uint64, error) {
	key := fmt.Sprintf(set+"|"+sharingFormat, shareID, password)
	return redisClient.Get(key).Uint64()
}

func StoreShare(shareID, password string, linkID uint, expiration int64) error {
	key := fmt.Sprintf(set+"|"+sharingFormat, shareID, password)
	fmt.Println("key: ", key)
	return redisClient.Set(key, linkID, time.Second*time.Duration(expiration)).Err()
}
