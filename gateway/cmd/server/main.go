package main

import (
	"github.com/PechatnovVladimir/gateway_gprs_test/gateway/internal/gateway"
	"log"
)

func main() {
	if err := gateway.Run(); err != nil {
		log.Fatal(err)
	}

}
