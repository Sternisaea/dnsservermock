package dnsservermock

import (
	"net"

	"github.com/sternisaea/dnsservermock/src/dnsconst"
	"github.com/sternisaea/dnsservermock/src/dnsstorage"
)

type TypeA struct {
}

func (t *TypeA) Process(req *DnsRequest, resp *DNSResponse, qst DnsQuestion, store dnsstorage.Storage) {
	result, err := store.Get(qst.Name, dnsconst.Type_A)
	if err != nil {
		(*resp).Flags.RCODE = dnsconst.RcodeNXDomain
		return
	}

	ip4 := net.ParseIP(result).To4()
	if ip4 == nil || len(ip4) != 4 {
		(*resp).Flags.RCODE = dnsconst.RcodeServFail
		return
	}

	answer := DnsAnswer{
		Name:     qst.Name,
		Type:     qst.Type,
		Class:    qst.Class,
		TTL:      3600,
		RDLength: uint16(len(ip4)),
		RData:    ip4,
	}

	(*resp).Answers = append((*resp).Answers, answer)
}
