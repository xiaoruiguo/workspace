//create: 2015/09/11 13:26:08 change: 2016/05/06 13:44:33 author:lijiao
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

func echoKV(c redis.Conn, key, vtype string, value ...interface{}) {
	fmt.Printf("Key: %s\n", string(key))
	args := make([]interface{}, 0, 0)
	args = append(args, key)
	args = append(args, value...)
	switch vtype {
	case "hash":
		c.Do("HMSET", args...)
		//c.Do("HMSET", key, "g","goo","y","yaho")
		c.Flush()
		ret, err := redis.Strings(c.Do("HGETALL", key))
		if err == nil {
			fmt.Printf("Value: %s\n", ret)
		} else {
			fmt.Printf("%s\n", err)
		}
	case "string":
		c.Do("SET", key, value)
		c.Flush()
		ret, err := redis.String(c.Do("GET", key))
		if err == nil {
			fmt.Printf("Value: %s\n", ret)
		} else {
			fmt.Printf("%s\n", err)
		}
	case "list":
		c.Do("DEL", key)
		c.Do("RPUSH", args...)
		c.Flush()
		ret, err := redis.Strings(c.Do("LRANGE", key, 0, -1))
		if err == nil {
			fmt.Printf("Value: %s\n", ret)
		} else {
			fmt.Printf("%s\n", err)
		}
	}
}

func main() {
	log.SetFlags(log.Lshortfile)
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal(err)
	}
	/*
		_, err = c.Do("auth", "dddd:1234")
		if err != nil {
			log.Fatal(err)
		}
	*/

	var key string

	key = "game_coin_plus_score_hash"
	echoKV(c, key, "hash", "google", "Newwww.google.com", "yahoo", "www.yahoo.com", "abcd", "1234")
	key = "game_coin_plus_score_hash_XXX"
	echoKV(c, key, "hash", "google2", "Newwww.google.com", "yahoo2", "www.yahoo.com", "abcd2", "1234")

	key = "game_coin_game_sign_in_counter"
	echoKV(c, key, "string", 3)
	key = "game_coin_game_sign_in_counter_XXX"
	echoKV(c, key, "string", 4)

	key = "game_coin_risk_user_"
	echoKV(c, key, "string", "Newadmin")
	key = "game_coin_risk_user_XXX"
	echoKV(c, key, "string", "admin2")

	key = "game_coin_risk_ip_"
	echoKV(c, key, "string", "192.168.1.1")
	key = "game_coin_risk_ip_XXX"
	echoKV(c, key, "string", "192.168.1.2")

	key = "game_coin_game_order_list_prefix_"
	echoKV(c, key, "list", "Newab", "cd", "ef")
	key = "game_coin_game_order_list_prefix_XXX"
	echoKV(c, key, "list", "Newab2", "cd2", "ef2")

	ret, err := redis.String(c.Do("info"))
	if err == nil {
		fmt.Printf("info: %s\n", ret)
	} else {
		fmt.Printf("%s\n", err)
	}

	rets, err := redis.Strings(c.Do("config", "get", "maxmemory"))
	if err == nil {
		fmt.Printf("config get maxmemory: %s, %s\n", rets[0], rets[1])
	} else {
		fmt.Printf("%s\n", err)
	}

	c.Close()
}
