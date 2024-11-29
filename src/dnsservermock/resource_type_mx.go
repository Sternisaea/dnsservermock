package dnsservermock

import (
	"bytes"
	"encoding/binary"

	"github.com/Sternisaea/dnsservermock/src/dnsconst"
	"github.com/Sternisaea/dnsservermock/src/dnsstorage"
)

type ResourceTypeMX struct {
	Base       ResourceBase
	Preference uint16
	Exchange   string
}

func NewResourceTypeMX(base ResourceBase) *ResourceTypeMX {
	return &ResourceTypeMX{Base: base}
}

func (r *ResourceTypeMX) Query(store dnsstorage.Storage) (dnsconst.Rcode, error) {
	result, err := store.Get((*r).Base.Name, (*r).Base.Type)
	if err != nil {
		return dnsconst.RcodeNXDomain, err
	}
	(*r).Preference = 10
	(*r).Exchange = result
	return dnsconst.RcodeNoError, nil
}

func (r *ResourceTypeMX) Write(buf *bytes.Buffer, dms *domains) {
	(*r).Base.Write(buf, dms)
	b := buf.Len()
	binary.Write(buf, binary.BigEndian, (*r).Preference)
	writeDomainName(buf, dms, (*r).Exchange)
	(*r).Base.RDLength = uint16(buf.Len() - b)
}
