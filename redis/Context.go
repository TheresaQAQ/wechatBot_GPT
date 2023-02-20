package redis

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var redisClient redis.Conn

/*func init() {
	redisClient, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer redisClient.Close()
}*/

// 拼接Prompt
func builderPrompt(old, add string, role string) string {
	newPrompt := fmt.Sprintf("%s%s:%s", old, role, add)
	return newPrompt
}

func WriteRedis(key, str, role string) error {
	redisClient, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return nil
	}
	old, err := ReadRedis(key)
	if err != nil {
		//不存在value
		_, err := redisClient.Do("SET", key, builderPrompt("", str, role))

		if err != nil {
			fmt.Println("redis set failed:", err)
			return err
		}
	} else {
		//存在value
		_, err := redisClient.Do("SET", key, builderPrompt(old, str, role))
		if err != nil {
			fmt.Println("redis set failed:", err)
			return err
		}
	}

	// 设置过期时间为1小时
	redisClient.Do("EXPIRE", key, 60*60)

	return nil
}

func ReadRedis(key string) (string, error) {
	redisClient, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return "", nil
	}
	defer redisClient.Close()
	reply, err := redisClient.Do("EXISTS", key)
	if err != nil {
		return "", err
	}

	if reply == false {
		return "", errors.New("key不存在")
	}

	//存在
	do, _ := redis.String(redisClient.Do("GET", key))
	return do, nil
}

func DelRedis(key string) error {
	redisClient, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return nil
	}
	defer redisClient.Close()
	redisClient.Do("DEL", key)
	return nil
}
