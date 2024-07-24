package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"math/rand"
	"runtime"
	"time"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		//Addr:     "localhost:16379",
		Addr:     "192.168.8.100:16379",
		Password: "easy-chat",
		DB:       0,
		PoolSize: 20,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	fmt.Println("链接redis成功！")
}

// string 操作  简单动态字符串类型
// 底层是一个struct
//
//	struct sdshdr{
//		//记录buf数组中已使用字节的数量
//		//等于 SDS 保存字符串的长度
//		int len;
//		//记录 buf 数组中未使用字节的数量
//		int free;
//		//字节数组，用于保存字符串
//		char buf[];
//	}
func test1() {
	ctx := context.Background()
	// 1. set aa bb             aa永不过期
	// set key value [EX seconds] [PX milliseconds]
	err1 := client.Set(ctx, "aa", "bb", 0).Err()
	if err1 != nil {
		fmt.Println("1: ", err1.Error())
	}

	// 2. setex cc 20 dd        cc30秒后过期
	err2 := client.SetEX(ctx, "cc", "dd", 30*time.Second).Err()
	if err2 != nil {
		fmt.Println("2: ", err2.Error())
	}

	// 3. get aa
	res1, err3 := client.Get(ctx, "aa").Result()
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	fmt.Println("3aa：", res1)

	// 3. get cc
	res2, err4 := client.Get(ctx, "cc").Result()
	if err4 != nil {
		fmt.Println(err4.Error())
	}
	fmt.Println("3cc：", res2)

	// 4. del aa
	err5 := client.Del(ctx, "aa").Err()
	if err4 != nil {
		fmt.Println(err5.Error())
	}
	res3, err6 := client.Get(ctx, "aa").Result()
	fmt.Println("4: ", res3, err6.Error())

	// 5. setnx cc eee  不存在时创建成功，存在时创建失败！
	flag1, err7 := client.SetNX(ctx, "cc", "ee", 30*time.Second).Result()
	fmt.Println("5: ", flag1, err7)

	err8 := client.Del(ctx, "cc").Err()
	if err4 != nil {
		fmt.Println(err8.Error())
	}

	flag2, err9 := client.SetNX(ctx, "cc", "ee", 30*time.Second).Result()
	fmt.Println("5-1: ", flag2, err9)

	// 6. 设置对象   set dct:company  {name:dct, age:20}
	v := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "dct",
		Age:  20,
	}
	bytes, _ := json.Marshal(v)
	err10 := client.Set(ctx, "dct:company", bytes, 30*time.Second).Err()
	fmt.Println("6-1: ", err10)

	bytes1, err15 := client.Get(ctx, "dct:company").Bytes()
	fmt.Println("6-2: ", string(bytes1), err15)

	// 7. 设置多个值 获取多个值
	// 这里注意 ...interface{}参数叫做可变参数，可以接收任意数量，任意类型的参数，太🐮了
	err11 := client.MSet(ctx, "ff", "gg", "hh", "ii").Err()
	err12 := client.MSet(ctx, []string{"jj", "kk", "ll", "mm"}).Err()
	err13 := client.MSet(ctx, map[string]string{"nn": "oo", "pp": "qq"}).Err()
	fmt.Println("7: ", err11, err12, err13)

	res4, err14 := client.MGet(ctx, "ff", "jj", "pp").Result()
	if err14 != nil {
		fmt.Println(err14)
	}

	for _, item := range res4 {
		if v, ok := item.(string); ok {
			fmt.Println("7-1", v)
		}
	}
}

const lockPrefix = "redis-lock"
const maxBackoff = 16
const unlockScript = `
id redis.call("get", KEYS[1]) == ARG[1] then
	return redis.call("del", KEYS[1])
else
	return 0
end`

type Lock struct {
	name    string
	traceId string
	time.Duration
}

func NewLock(name string, duration time.Duration) *Lock {
	return &Lock{
		name:     name,
		traceId:  uuid.NewString(),
		Duration: duration,
	}
}

func (l *Lock) GetLock(ctx context.Context) error {
	backoff := 1
	for {
		// 基于 setnx 实现分布式锁
		ok, err := client.SetNX(ctx, lockPrefix+l.name, l.traceId, l.Duration).Result() // 这个地方设置key成功还是失败err都是nil, 要根据result判断是否加锁成功
		if err != nil {
			return err
		}

		if ok {
			break
		}

		for i := 0; i < backoff; i++ {
			runtime.Gosched()
		}

		if backoff < maxBackoff {
			backoff <<= 1 // 左移乘以2
		}

		// 增加随机退避时间
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	}
	return nil
}

func (l *Lock) Unlock(ctx context.Context) {
	script := redis.NewScript(unlockScript)
	script.Run(ctx, client, []string{lockPrefix + l.name}, l.traceId)
}

// 列表类型 list  列表只能存储字符串
func test2() {

	ctx := context.Background()
	// 1.从左或者从右添加元素  lpush aa bb
	err1 := client.LPush(ctx, "aa", "bb").Err()
	if err1 != nil {
		fmt.Println(err1)
	}

	//  2.rpush aa cc
	err2 := client.RPush(ctx, "aa", "cc").Err()
	if err2 != nil {
		fmt.Println(err2)
	}

	// 3.lrange aa 0 -1
	res1, err3 := client.LRange(ctx, "aa", 0, -1).Result()
	if err3 != nil {
		fmt.Println(err3)
	}
	for _, item := range res1 {
		fmt.Println("3:", item)
	}

	// 4. linsert aa before cc dd
	err4 := client.LInsertBefore(ctx, "aa", "cc", "dd").Err()
	if err4 != nil {
		fmt.Println(err4)
	}

	res2, err5 := client.LRange(ctx, "aa", 0, -1).Result()
	if err5 != nil {
		fmt.Println(err5)
	}
	for _, item := range res2 {
		fmt.Println("4:", item)
	}

	// 5. lindex aa 0
	res3, err6 := client.LIndex(ctx, "aa", 1).Result()
	if err5 != nil {
		fmt.Println(err6)
	}
	fmt.Println(res3)

	// 6. 获取列表的长度 llen  aa
	res4, err7 := client.LLen(ctx, "aa").Result()
	if err7 != nil {
		fmt.Println(err7)
	}
	fmt.Println(res4)

	// 7.删除列表元素  lpop aa  rpop aa
	res5, err8 := client.LPop(ctx, "aa").Result()
	if err8 != nil {
		fmt.Println(err8)
	}
	fmt.Println(res5)

}

func main() {
	defer client.Close()
	//test1()
	test2()
}
