package main

import (
	"kang-by-xoverse/api/rest"
	"kang-by-xoverse/api/tcp"
	"kang-by-xoverse/api/utils"
)

func main() {
	// fmt.Println(generateWallet())

	utils.LoadDotEnv()
	rdb, close := utils.GetRedisClient()

	go tcp.CreateEventsServer(rdb)
	rest.RunRestServer(rdb)

	close()
}
