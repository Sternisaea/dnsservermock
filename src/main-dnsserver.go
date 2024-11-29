package main

import (
	"fmt"
	"net"
	"time"

	"github.com/Sternisaea/dnsservermock/src/dnsconst"
	"github.com/Sternisaea/dnsservermock/src/dnsservermock"
	"github.com/Sternisaea/dnsservermock/src/dnsstorage/dnsstoragememory"
)

func main() {
	store := dnsstoragememory.NewMemoryStore()
	(*store).Set("test.com", dnsconst.Type_A, "127.0.0.1")
	(*store).Set("mail.test.com", dnsconst.Type_A, "127.0.0.1")
	(*store).Set("test.com", dnsconst.Type_MX, "mail.test.com")
	(*store).Set("test.com", dnsconst.Type_AAAA, "::1")

	ds := dnsservermock.NewDnsServer(net.ParseIP("127.0.0.1"), 5355, store)
	if err := (*ds).Start(); err != nil {
		panic(err)
	}

	time.Sleep(5 * time.Second)

	if err := (*ds).Stop(); err != nil {
		fmt.Println(err)
	}
}
