package dnsservermock

import (
	"errors"
	"fmt"

	"github.com/sternisaea/dnsservermock/src/dnsconst"
	"github.com/sternisaea/dnsservermock/src/dnsstorage"
)

var (
	ErrNotSupportedType = errors.New("type not supported")
	ErrUnknownType      = errors.New("unknown type ")
)

type TypeProcess interface {
	Process(req *DNSRequest, resp *DNSResponse, qst DNSQuestion, store dnsstorage.Storage)
}

func GetProcess(dnsType dnsconst.DnsType) (TypeProcess, error) {
	switch dnsType {
	case dnsconst.Type_A:
		return &TypeA{}, nil
	case dnsconst.Type_NS:
		return nil, fmt.Errorf("type %s (%d): %w", "NS", dnsType, ErrNotSupportedType)
	case dnsconst.Type_MD:
		return nil, fmt.Errorf("type %s (%d): %w", "MD", dnsType, ErrNotSupportedType)
	case dnsconst.Type_MF:
		return nil, fmt.Errorf("type %s (%d): %w", "MF", dnsType, ErrNotSupportedType)
	case dnsconst.Type_CNAME:
		return nil, fmt.Errorf("type %s (%d): %w", "CNAME", dnsType, ErrNotSupportedType)
	case dnsconst.Type_SOA:
		return nil, fmt.Errorf("type %s (%d): %w", "SOA", dnsType, ErrNotSupportedType)
	case dnsconst.Type_MB:
		return nil, fmt.Errorf("type %s (%d): %w", "MB", dnsType, ErrNotSupportedType)
	case dnsconst.Type_MG:
		return nil, fmt.Errorf("type %s (%d): %w", "MG", dnsType, ErrNotSupportedType)
	case dnsconst.Type_MR:
		return nil, fmt.Errorf("type %s (%d): %w", "MR", dnsType, ErrNotSupportedType)
	case dnsconst.Type_NULL:
		return nil, fmt.Errorf("type %s (%d): %w", "NULL", dnsType, ErrNotSupportedType)
	case dnsconst.Type_WKS:
		return nil, fmt.Errorf("type %s (%d): %w", "WKS", dnsType, ErrNotSupportedType)
	case dnsconst.Type_PTR:
		return nil, fmt.Errorf("type %s (%d): %w", "PTR", dnsType, ErrNotSupportedType)
	case dnsconst.Type_HINFO:
		return nil, fmt.Errorf("type %s (%d): %w", "HINFO", dnsType, ErrNotSupportedType)
	case dnsconst.Type_MINFO:
		return nil, fmt.Errorf("type %s (%d): %w", "MINFO", dnsType, ErrNotSupportedType)
	case dnsconst.Type_MX:
		return &TypeMX{}, nil
	case dnsconst.Type_TXT:
		return nil, fmt.Errorf("type %s (%d): %w", "TXT", dnsType, ErrNotSupportedType)
	case dnsconst.Type_AAAA:
		return nil, fmt.Errorf("type %s (%d): %w", "AAAA", dnsType, ErrNotSupportedType)
	case dnsconst.Type_SRV:
		return nil, fmt.Errorf("type %s (%d): %w", "SRV", dnsType, ErrNotSupportedType)
	case dnsconst.Type_NAPTR:
		return nil, fmt.Errorf("type %s (%d): %w", "NAPTR", dnsType, ErrNotSupportedType)
	case dnsconst.Type_OPT:
		return nil, fmt.Errorf("type %s (%d): %w", "OPT", dnsType, ErrNotSupportedType)
	case dnsconst.Type_DS:
		return nil, fmt.Errorf("type %s (%d): %w", "DS", dnsType, ErrNotSupportedType)
	case dnsconst.Type_RRSIG:
		return nil, fmt.Errorf("type %s (%d): %w", "RRSIG", dnsType, ErrNotSupportedType)
	case dnsconst.Type_NSEC:
		return nil, fmt.Errorf("type %s (%d): %w", "NSEC", dnsType, ErrNotSupportedType)
	case dnsconst.Type_DNSKEY:
		return nil, fmt.Errorf("type %s (%d): %w", "DNSKEY", dnsType, ErrNotSupportedType)
	case dnsconst.Type_TLSA:
		return nil, fmt.Errorf("type %s (%d): %w", "TLSA", dnsType, ErrNotSupportedType)
	case dnsconst.Type_SPF:
		return nil, fmt.Errorf("type %s (%d): %w", "SPF", dnsType, ErrNotSupportedType)
	case dnsconst.Type_AXFR:
		return nil, fmt.Errorf("type %s (%d): %w", "AXFR", dnsType, ErrNotSupportedType)
	case dnsconst.Type_MAILB:
		return nil, fmt.Errorf("type %s (%d): %w", "MAILB", dnsType, ErrNotSupportedType)
	case dnsconst.Type_MAILA:
		return nil, fmt.Errorf("type %s (%d): %w", "MAILA", dnsType, ErrNotSupportedType)
	case dnsconst.Type_ANY:
		return nil, fmt.Errorf("type %s (%d): %w", "*", dnsType, ErrNotSupportedType)
	}
	return nil, fmt.Errorf("type %d: %w", dnsType, ErrUnknownType)
}
