package main

import (
	"fmt"
	"net"
	"time"

	"github.com/sternisaea/dnsservermock/src/dnsservermock"
)

func main() {

	ds := dnsservermock.NewDnsServer(net.ParseIP("127.0.0.1"), 5355)
	if err := (*ds).Start(); err != nil {
		panic(err)
	}

	time.Sleep(900 * time.Second)

	if err := (*ds).Stop(); err != nil {
		fmt.Println(err)
	}
}
