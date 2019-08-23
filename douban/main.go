package main

import "crawler/douban/service"

func main() {
	service.InitDB()
	service.Start()
}