package main

import (
	g "github.com/PechatnovVladimir/gateway_gprs_test/gateway/internal/gateway"
	"log"
)

func main() {
	log.Println("starting gateway...")
	go g.RunRest()

	for {
		_ = 1
	}
}
