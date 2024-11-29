package dnsservermock

import (
	"bytes"
	"net"

	"github.com/Sternisaea/dnsservermock/src/dnsconst"
	"github.com/Sternisaea/dnsservermock/src/dnsstorage"
)

type ResourceTypeAAAA struct {
	Base    ResourceBase
	Address net.IP
}

func NewResourceTypeAAAA(base ResourceBase) *ResourceTypeAAAA {
	return &ResourceTypeAAAA{Base: base}
}

func (r *ResourceTypeAAAA) Query(store dnsstorage.Storage) (dnsconst.Rcode, error) {
	result, err := store.Get((*r).Base.Name, (*r).Base.Type)
	if err != nil {
		return dnsconst.RcodeNXDomain, err
	}
	ip6 := net.ParseIP(result).To16()
	if ip6 == nil || len(ip6) != 16 {
		return dnsconst.RcodeServFail, err
	}
	(*r).Address = ip6
	return dnsconst.RcodeNoError, nil
}

func (r *ResourceTypeAAAA) Write(buf *bytes.Buffer, dms *domains) {
	(*r).Base.Write(buf, dms)
	b := buf.Len()
	buf.Write((*r).Address)
	(*r).Base.RDLength = uint16(buf.Len() - b)
}
