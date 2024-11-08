package dnsservermock

import (
	"fmt"

	"github.com/sternisaea/dnsservermock/src/dnsconst"
	"github.com/sternisaea/dnsservermock/src/dnsstorage"
)

type TypeMX struct {
}

func (t *TypeMX) Process(req *DnsRequest, resp *DNSResponse, qst DnsQuestion, store dnsstorage.Storage) {
	result, err := store.Get(qst.Name, dnsconst.Type_MX)
	if err != nil {
		(*resp).Flags.RCODE = dnsconst.RcodeNXDomain
		return
	}

	// Create Answer for qst.Name with result
	fmt.Printf("RESULT: %s\n", result)

}
