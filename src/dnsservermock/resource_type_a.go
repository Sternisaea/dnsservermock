package dnsservermock

import (
	"bytes"
	"encoding/binary"
	"net"

	"github.com/sternisaea/dnsservermock/src/dnsconst"
	"github.com/sternisaea/dnsservermock/src/dnsstorage"
)

type ResourceTypeA struct {
	Base    ResourceBase
	Address net.IP
}

func NewResourceTypeA(base ResourceBase) *ResourceTypeA {
	return &ResourceTypeA{Base: base}
}

func (r *ResourceTypeA) Query(store dnsstorage.Storage) (dnsconst.Rcode, error) {
	result, err := store.Get((*r).Base.Name, (*r).Base.Type)
	if err != nil {
		return dnsconst.RcodeNXDomain, err
	}
	ip4 := net.ParseIP(result).To4()
	if ip4 == nil || len(ip4) != 4 {
		return dnsconst.RcodeServFail, err
	}
	(*r).Address = ip4
	return dnsconst.RcodeNoError, nil
}

func (r *ResourceTypeA) Write(buf *bytes.Buffer, dms *domains) {
	(*r).Base.Write(buf, dms)
	binary.Write(buf, binary.BigEndian, uint16(len((*r).Address)))
	buf.Write((*r).Address)
}
