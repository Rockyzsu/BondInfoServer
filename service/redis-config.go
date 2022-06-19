package service

import (
	"encoding/json"
	"log"
	"os"
)

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func ReadRedisConfig() *RedisConfig {

	file, err := os.Open("service/redis.json")
	if err != nil {
		log.Fatalln(err)
	}

	decoder := json.NewDecoder(file)
	conf := &RedisConfig{}
	err = decoder.Decode(conf)
	if err != nil {
		log.Fatalln(err)
	}

	return conf
}
