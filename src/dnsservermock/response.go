package dnsservermock

import (
	"bytes"
	"encoding/binary"
)

type DnsResponse struct {
	ID    uint16
	Flags DnsFlags
	// QDCount   uint16
	// ANCount   uint16
	// NSCount   uint16
	// ARCount   uint16
	Questions   []DnsQuestion
	Answers     []DnsResource
	Authorities []DnsResourceRecord
	Additionals []DnsResourceRecord
}

type domains map[string]int

func NewDnsResponse(id uint16) *DnsResponse {
	return &DnsResponse{
		ID: id,
		Flags: DnsFlags{
			QR: true,
		},
	}
}

func (resp *DnsResponse) Write() []byte {
	(*resp).writePass()        // First pass to fill length fields
	return (*resp).writePass() // Second pass with full data
}

func (resp *DnsResponse) writePass() []byte {
	dms := make(domains, len(resp.Answers))
	var buf bytes.Buffer

	binary.Write(&buf, binary.BigEndian, resp.ID)
	binary.Write(&buf, binary.BigEndian, (*resp).Flags.Get())
	binary.Write(&buf, binary.BigEndian, uint16(len(resp.Questions)))
	binary.Write(&buf, binary.BigEndian, uint16(len(resp.Answers)))
	binary.Write(&buf, binary.BigEndian, uint16(len(resp.Authorities)))
	binary.Write(&buf, binary.BigEndian, uint16(len(resp.Additionals)))

	// Write questions
	for _, q := range resp.Questions {
		writeDomainName(&buf, &dms, q.Name)
		binary.Write(&buf, binary.BigEndian, q.Type)
		binary.Write(&buf, binary.BigEndian, q.Class)
	}

	// Write answers
	for _, a := range resp.Answers {
		a.Write(&buf, &dms)
	}

	// Write authority records
	// for _, rr := range resp.Authorities {
	// 	writeResourceRecord(&buf, rr)
	// }

	// Write additional records
	// for _, rr := range resp.Additionals {
	// 	writeResourceRecord(&buf, rr)
	// }
	return buf.Bytes()
}
