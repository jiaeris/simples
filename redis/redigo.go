package redis

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func conn() {

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	//do函数使用
	v, err := c.Do("SET", "name", "yunga")
	if err != nil {
		panic(err)
	}
	fmt.Println(v)

	//获取string
	s, err := redis.String(c.Do("GET", "name"))
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	r, err := c.Do("DEL", "name")
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	//redis.Map
	c.Do("SET", "HiByte", []byte("helloCCC"))

	buf, err := redis.Bytes(c.Do("GET", "HiByte"))
	if err != nil {
		panic(err)
	}
	fmt.Println(buf)
	fmt.Println(string(buf))

}

//发布
func publish() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
		return
	}
	defer c.Close()
	for {
		var s string
		fmt.Scanln(&s)
		_, err := c.Do("PUBLISH", "chatRoom", s)
		if err != nil {
			panic(err)
			return
		}
	}
}

//订阅
func subcribe() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	psc := redis.PubSubConn{c}
	err = psc.Subscribe("chatRoom")
	if err != nil {
		panic(err)
	}
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}

}
