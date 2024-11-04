package dnsservermock

import (
	"fmt"

	"github.com/sternisaea/dnsservermock/src/dnsstorage"
	"github.com/sternisaea/dnsservermock/src/dnstypes"
)

type TypeA struct {
}

func (t *TypeA) Process(req *DNSRequest, resp *DNSResponse, qst DNSQuestion, store dnsstorage.Storage) {
	result, err := store.Get(qst.Name, dnstypes.Type_A)
	if err != nil {
		(*resp).Flags.RCODE = dnstypes.RcodeNXDomain
		return
	}

	// Create Answer for qst.Name with result
	fmt.Printf("RESULT: %s\n", result)
}
