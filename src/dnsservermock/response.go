package dnsservermock

import (
	"bytes"
	"encoding/binary"
	"strings"
)

type DNSResponse struct {
	ID          uint16
	Flags       DNSFlags
	QDCount     uint16
	ANCount     uint16
	NSCount     uint16
	ARCount     uint16
	Questions   []DNSQuestion
	Answers     []DNSResourceRecord
	Authorities []DNSResourceRecord
	Additionals []DNSResourceRecord
}

func (resp *DNSResponse) CopyHeaderAndQuestions(req *DNSRequest) {
	(*resp).ID = (*req).ID
	(*resp).Flags = (*req).Flags
	(*resp).Flags.QR = true
	(*resp).Flags.RA = false
	(*resp).QDCount = (*req).QDCount
	(*resp).Questions = make([]DNSQuestion, 0, len((*req).Questions))
	copy((*resp).Questions, (*req).Questions)
}

func (resp *DNSResponse) SerializeResponse() []byte {
	var buf bytes.Buffer

	// Write DNS header
	binary.Write(&buf, binary.BigEndian, resp.ID)
	binary.Write(&buf, binary.BigEndian, (*resp).Flags.Get())
	binary.Write(&buf, binary.BigEndian, uint16(len(resp.Questions)))
	binary.Write(&buf, binary.BigEndian, uint16(len(resp.Answers)))
	binary.Write(&buf, binary.BigEndian, uint16(len(resp.Authorities)))
	binary.Write(&buf, binary.BigEndian, uint16(len(resp.Additionals)))

	// Write questions
	for _, q := range resp.Questions {
		writeDomainName(&buf, q.Name)
		binary.Write(&buf, binary.BigEndian, q.Type)
		binary.Write(&buf, binary.BigEndian, q.Class)
	}

	// Write answers
	for _, rr := range resp.Answers {
		writeResourceRecord(&buf, rr)
	}

	// Write authority records
	for _, rr := range resp.Authorities {
		writeResourceRecord(&buf, rr)
	}

	// Write additional records
	for _, rr := range resp.Additionals {
		writeResourceRecord(&buf, rr)
	}

	return buf.Bytes()
}

func writeDomainName(buf *bytes.Buffer, name string) {
	parts := strings.Split(name, ".")
	for _, part := range parts {
		buf.WriteByte(byte(len(part)))
		buf.WriteString(part)
	}
	buf.WriteByte(0) // End of domain name
}

func writeResourceRecord(buf *bytes.Buffer, rr DNSResourceRecord) {
	writeDomainName(buf, rr.Name)
	binary.Write(buf, binary.BigEndian, rr.Type)
	binary.Write(buf, binary.BigEndian, rr.Class)
	binary.Write(buf, binary.BigEndian, rr.TTL)
	binary.Write(buf, binary.BigEndian, uint16(len(rr.RData)))
	buf.Write(rr.RData)
}
