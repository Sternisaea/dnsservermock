package main

import (
	"fmt"
	"net"
	"time"

	"github.com/sternisaea/dnsservermock/src/dnsservermock"
	"github.com/sternisaea/dnsservermock/src/dnsstorage/dnsstoragememory"
	"github.com/sternisaea/dnsservermock/src/dnstypes"
)

func main() {
	store := dnsstoragememory.NewMemoryStore()
	(*store).Set("test.com", dnstypes.Type_A, "127.0.0.1")
	(*store).Set("mail.test.com", dnstypes.Type_A, "127.0.0.1")
	(*store).Set("test.com", dnstypes.Type_MX, "mail.test.com")

	ds := dnsservermock.NewDnsServer(net.ParseIP("127.0.0.1"), 5355, store)
	if err := (*ds).Start(); err != nil {
		panic(err)
	}

	time.Sleep(900 * time.Second)

	if err := (*ds).Stop(); err != nil {
		fmt.Println(err)
	}
}
