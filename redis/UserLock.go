package redis

import (
	"awesomeProject/utils/constant"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/garyburd/redigo/redis"
)

var profix = "lock:user"

func setLock(key string) error {
	redisClient, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return nil
	}

	_, err = redisClient.Do("SET", key, "lock")
	if err != nil {
		return err
	}

	//有效期两分钟
	_, err = redisClient.Do("EXPIRE", key, 2*60)
	if err != nil {
		return err
	}
	return nil
}

func getKeyName(message *openwechat.Message) string {

	sender, _ := message.Sender()
	return constant.FriendLock + sender.UserName

}

func IsLock(msg *openwechat.Message) (bool, error) {
	redisClient, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return false, err
	}

	b, err := redis.Bool(redisClient.Do("EXSITS", getKeyName(msg)))
	if err != nil {
		return false, err
	}

	return b, nil
}
