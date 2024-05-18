package main

import (
	"log"
	social "tiktokrpc/kitex_gen/social/socialservice"
)

func main() {
	svr := social.NewServer(new(SocialServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
