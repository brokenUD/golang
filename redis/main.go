package main

import (
	"fmt"
	"time"
	// "strconv"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func initRedisClient() error {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func RedisSet(key, value string, t time.Duration) error {
	err := rdb.Set(key, value, t).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func RedisGet(key string) (value string, err error) {
	value, err = rdb.Get(key).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func RedisMSet(kvs []string) error {
	err := rdb.MSet(kvs).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func RedisMGet(keys []string) (values []string, err error) {
	vs, err := rdb.MGet(keys...).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range vs {
		vs := fmt.Sprintf("%v", v)
		values = append(values, vs)
	}
	return
}

func RedisLPush(key string, values []interface{}) (err error) {
	err = rdb.LPush(key, values...).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func RedisLRange(key string) (values []string, err error) {
	values, err = rdb.LRange(key, 0, -1).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func RedisLPop(key string) (value string, err error) {
	value, err = rdb.LPop(key).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func RedisHSet(key, field string, value interface{}) error {
	err := rdb.HSet(key, field, value).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func RedisHGet(key, field string) (value string, err error) {
	value, err = rdb.HGet(key, field).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func RedisHMSet(key string, fields map[string]interface{}) error {
	err := rdb.HMSet(key, fields).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func RedisHMGet(key string, fields []string) (values []string, err error) {
	vs, err := rdb.HMGet(key, fields...).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range vs {
		vs := fmt.Sprintf("%v", v)
		values = append(values, vs)
	}
	return
}

func RedisSAdd(key string, members ...interface{}) (err error) {
	err = rdb.SAdd(key, members...).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func RedisSMembers(key string) (values []string, err error) {
	values, err = rdb.SMembers(key).Result()
	if err != nil {
		fmt.Printf("SMembers error key:%v,  err:%v", key, err)
		return
	}
	return
}

func RedisSIsMember(key string, member interface{}) (is bool, err error) {
	is, err = rdb.SIsMember(key, member).Result()
	if err != nil {
		fmt.Printf("SIsMember error key:%v,  member:%v, err:%v", key, member, err)
		return
	}
	return
}

func RedisSCard(key string) (count int64, err error) {
	count, err = rdb.SCard(key).Result()
	if err != nil {
		fmt.Printf("SCard error key:%v,  err:%v", key, err)
		return
	}
	return
}

func RedisZAdd(key string, scores []float64, members []string) (count int64, err error) {
	values := []redis.Z{}
	for i, v := range scores {
		values = append(values, redis.Z{
			Score:  v,
			Member: members[i],
		})
	}
	count, err = rdb.ZAdd(key, values...).Result()
	if err != nil {
		fmt.Printf("ZAdd error key:%v, values:%v, err:%v", key, values, err)
		return
	}
	return
}

func RedisZRem(key string, members ...interface{}) (count int64, err error) {
	count, err = rdb.ZRem(key, members...).Result()
	if err != nil {
		fmt.Printf("ZRem error key:%v, members:%v, err:%v", key, members, err)
		return
	}
	return
}

func RedisZRange(key string) (values []string, err error) {
	values, err = rdb.ZRange(key, 0, -1).Result()
	if err != nil {
		fmt.Printf("ZRange error key:%v, err:%v", key, err)
		return
	}
	return
}

func RedisZIncrBy(key, member string, inc float64) (score float64, err error) {
	score, err = rdb.ZIncrBy(key, inc, member).Result()
	if err != nil {
		fmt.Printf("ZIncrBy error, key:%v, inc:%v, member:%v, err:%v", key, inc, member, err)
		return
	}
	return
}

func RedisZRangeByScoreWithScores(key string, min, max string, offset, count int64) (values map[string]float64) {
	res, err := rdb.ZRangeByScoreWithScores(key, redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}).Result()
	if err != nil {
		fmt.Printf("ZRangeByScoreWithScores error, key:%v, min:%v, max:%v, offset:%v, count:%v, err:%v", key, min, max, offset, count, err)
		return
	}
	values = make(map[string]float64, len(res))
	for _, v := range res {
		values[v.Member.(string)] = v.Score
	}
	return
}

func RedisZRevRangeWithScores(key string, min, max int64) (values map[string]float64) {
	res, err := rdb.ZRevRangeWithScores(key, min, max).Result()
	if err != nil {
		fmt.Printf("ZRevRangeWithScores error, key:%v, min:%v, max:%v, err:%v", key, min, max, err)
		return
	}
	values = make(map[string]float64, len(res))
	for _, v := range res {
		values[v.Member.(string)] = v.Score
	}
	return
}

func main() {
	initRedisClient()
	RedisSet("a", "111", 0)
}
