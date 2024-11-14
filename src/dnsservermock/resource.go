package dnsservermock

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/sternisaea/dnsservermock/src/dnsconst"
	"github.com/sternisaea/dnsservermock/src/dnsstorage"
)

type DnsResource interface {
	Query(store dnsstorage.Storage) (dnsconst.Rcode, error)
	Write(buf *bytes.Buffer, dms *domains)
}

var (
	ErrNotSupportedType = errors.New("type not supported")
	ErrUnknownType      = errors.New("unknown type ")
)

func GetResource(qst DnsQuestion) (DnsResource, error) {
	base := ResourceBase{
		Name:  qst.Name,
		Type:  dnsconst.DnsType(qst.Type),
		Class: dnsconst.DnsClass(qst.Class),
	}
	switch dtype := dnsconst.DnsType(qst.Type); dtype {
	case dnsconst.Type_A:
		return NewResourceTypeA(base), nil
	case dnsconst.Type_NS:
		return nil, fmt.Errorf("type %s (%d): %w", "NS", dtype, ErrNotSupportedType)
	case dnsconst.Type_MD:
		return nil, fmt.Errorf("type %s (%d): %w", "MD", dtype, ErrNotSupportedType)
	case dnsconst.Type_MF:
		return nil, fmt.Errorf("type %s (%d): %w", "MF", dtype, ErrNotSupportedType)
	case dnsconst.Type_CNAME:
		return nil, fmt.Errorf("type %s (%d): %w", "CNAME", dtype, ErrNotSupportedType)
	case dnsconst.Type_SOA:
		return nil, fmt.Errorf("type %s (%d): %w", "SOA", dtype, ErrNotSupportedType)
	case dnsconst.Type_MB:
		return nil, fmt.Errorf("type %s (%d): %w", "MB", dtype, ErrNotSupportedType)
	case dnsconst.Type_MG:
		return nil, fmt.Errorf("type %s (%d): %w", "MG", dtype, ErrNotSupportedType)
	case dnsconst.Type_MR:
		return nil, fmt.Errorf("type %s (%d): %w", "MR", dtype, ErrNotSupportedType)
	case dnsconst.Type_NULL:
		return nil, fmt.Errorf("type %s (%d): %w", "NULL", dtype, ErrNotSupportedType)
	case dnsconst.Type_WKS:
		return nil, fmt.Errorf("type %s (%d): %w", "WKS", dtype, ErrNotSupportedType)
	case dnsconst.Type_PTR:
		return nil, fmt.Errorf("type %s (%d): %w", "PTR", dtype, ErrNotSupportedType)
	case dnsconst.Type_HINFO:
		return nil, fmt.Errorf("type %s (%d): %w", "HINFO", dtype, ErrNotSupportedType)
	case dnsconst.Type_MINFO:
		return nil, fmt.Errorf("type %s (%d): %w", "MINFO", dtype, ErrNotSupportedType)
	case dnsconst.Type_MX:
		return NewResourceTypeMX(base), nil
	case dnsconst.Type_TXT:
		return nil, fmt.Errorf("type %s (%d): %w", "TXT", dtype, ErrNotSupportedType)
	case dnsconst.Type_AAAA:
		return NewResourceTypeAAAA(base), nil
	case dnsconst.Type_SRV:
		return nil, fmt.Errorf("type %s (%d): %w", "SRV", dtype, ErrNotSupportedType)
	case dnsconst.Type_NAPTR:
		return nil, fmt.Errorf("type %s (%d): %w", "NAPTR", dtype, ErrNotSupportedType)
	case dnsconst.Type_OPT:
		return nil, fmt.Errorf("type %s (%d): %w", "OPT", dtype, ErrNotSupportedType)
	case dnsconst.Type_DS:
		return nil, fmt.Errorf("type %s (%d): %w", "DS", dtype, ErrNotSupportedType)
	case dnsconst.Type_RRSIG:
		return nil, fmt.Errorf("type %s (%d): %w", "RRSIG", dtype, ErrNotSupportedType)
	case dnsconst.Type_NSEC:
		return nil, fmt.Errorf("type %s (%d): %w", "NSEC", dtype, ErrNotSupportedType)
	case dnsconst.Type_DNSKEY:
		return nil, fmt.Errorf("type %s (%d): %w", "DNSKEY", dtype, ErrNotSupportedType)
	case dnsconst.Type_TLSA:
		return nil, fmt.Errorf("type %s (%d): %w", "TLSA", dtype, ErrNotSupportedType)
	case dnsconst.Type_SPF:
		return nil, fmt.Errorf("type %s (%d): %w", "SPF", dtype, ErrNotSupportedType)
	case dnsconst.Type_AXFR:
		return nil, fmt.Errorf("type %s (%d): %w", "AXFR", dtype, ErrNotSupportedType)
	case dnsconst.Type_MAILB:
		return nil, fmt.Errorf("type %s (%d): %w", "MAILB", dtype, ErrNotSupportedType)
	case dnsconst.Type_MAILA:
		return nil, fmt.Errorf("type %s (%d): %w", "MAILA", dtype, ErrNotSupportedType)
	case dnsconst.Type_ANY:
		return nil, fmt.Errorf("type %s (%d): %w", "*", dtype, ErrNotSupportedType)
	default:
		return nil, fmt.Errorf("type %d: %w", dtype, ErrUnknownType)
	}
}
