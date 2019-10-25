package models

import (
	"github.com/go-redis/redis"
	"go_email/src/pkg/setting"
	"log"
	"time"
)

var (
	Rdb *redis.Client
)

func init() {
	// 获取redis 配置
	r, err := setting.Cfg.GetSection("redis")
	if err != nil {
		log.Fatal(2, "Fail to get section 'redis': %v", err)
	}

	ins, _ := r.Key("DB").Int()
	ConnectRedis(r.Key("ADDRESS").String(),
		r.Key("PASSWD").String(),
		ins)
}

func ConnectRedis(addr string, passwd string, db int){
	Rdb = redis.NewClient(&redis.Options{
		Addr 		: addr,
		Password	: passwd,
		DB       	: db,
	});
}

func GetValByKey(key string) interface{}{
	return  Rdb.Get(key).Val()
}

func SetValByKey(key string, val interface{}, expiration time.Duration) error{
	_, err :=Rdb.Set(key, val, expiration).Result()

	return  err
}

func PutinQueue(queue_url string, url string){
	Rdb.LPush(queue_url, []byte(url))
}

func PopfromQueue(queue_url string) string{
	res, err := Rdb.RPop(queue_url).Result()
	if err != nil{
		panic(err)
	}

	return res
}

func GetQueueLength(queue_url string) int64{
	length,err := Rdb.LLen(queue_url).Result()
	if err != nil{
		return 0
	}

	return length
}