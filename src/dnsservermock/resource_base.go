package dnsservermock

import (
	"bytes"
	"encoding/binary"

	"github.com/sternisaea/dnsservermock/src/dnsconst"
)

type ResourceBase struct {
	Name     string
	Type     dnsconst.DnsType
	Class    dnsconst.DnsClass
	TTL      uint32
	RDLength uint16
}

func NewResourceBase(name string, dtype dnsconst.DnsType, dclass dnsconst.DnsClass) *ResourceBase {
	return &ResourceBase{
		Name:  name,
		Type:  dtype,
		Class: dclass,
		TTL:   3600,
	}
}

func (r *ResourceBase) Write(buf *bytes.Buffer, dms *domains) {
	writeDomainName(buf, dms, (*r).Name)
	binary.Write(buf, binary.BigEndian, (*r).Type)
	binary.Write(buf, binary.BigEndian, (*r).Class)
	binary.Write(buf, binary.BigEndian, (*r).TTL)
	binary.Write(buf, binary.BigEndian, (*r).RDLength)
}
