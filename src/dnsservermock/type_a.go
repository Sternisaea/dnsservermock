package dnsservermock

import (
	"net"

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

	ip4 := net.ParseIP(result).To4()
	if ip4 == nil || len(ip4) != 4 {
		(*resp).Flags.RCODE = dnstypes.RcodeServFail
		return
	}

	answer := DNSAnswer{
		Name:     qst.Name,
		Type:     qst.Type,
		Class:    qst.Class,
		TTL:      3600,
		RDLength: uint16(len(ip4)),
		RData:    ip4,
	}

	(*resp).Answers = append((*resp).Answers, answer)
}
