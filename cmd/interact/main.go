package main

import (
	"log"
	interact "tiktokrpc/kitex_gen/interact/interactservice"
)

func main() {
	svr := interact.NewServer(new(InteractServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
