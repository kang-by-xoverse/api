package main

import (
	"kang-by-xoverse/api/rest"
	"kang-by-xoverse/api/utils"
)

func main() {
	// fmt.Println(generateWallet())

	utils.LoadDotEnv()
	rdb, close := utils.GetRedisClient()

	rest.RunRestServer(rdb)

	close()
}
