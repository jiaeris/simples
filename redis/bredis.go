package redis

import "github.com/garyburd/redigo/redis"

type BaseRedis struct {
	*redis.Pool
}

//一些说明：
//MaxIdle：最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
//MaxActive：最大的激活连接数，表示同时最多有N个连接
//IdleTimeout：最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
const (
	MaxIdle = 10
)

func NewBaseRedis() *BaseRedis {
	redisPool := &redis.Pool{
		MaxIdle: MaxIdle,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	br := &BaseRedis{
		redisPool,
	}
	return br
}

func (br *BaseRedis) SetString(key, val string) error {
	redisConn := br.Get()
	_, err := redisConn.Do("SET", key, val)
	return err
}

func (br *BaseRedis) GetString(key string) string {
	redisConn := br.Get()
	val, err := redis.String(redisConn.Do("GET", key))
	if err != nil {
		return ""
	}
	return val
}
