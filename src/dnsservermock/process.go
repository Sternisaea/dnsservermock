package dnsservermock

import (
	"errors"
	"fmt"
)

var (
	ErrNotSupportedType = errors.New("type not supported")
	ErrUnknownType      = errors.New("unknown type ")
)

type TypeProcess interface {
	Process(req *DNSRequest, resp *DNSResponse, qst DNSQuestion)
}

func GetProcess(dnsType DnsType) (TypeProcess, error) {
	switch dnsType {
	case Type_A:
		return &TypeA{}, nil
	case Type_NS:
		return nil, fmt.Errorf("type %s (%d): %w", "NS", dnsType, ErrNotSupportedType)
	case Type_MD:
		return nil, fmt.Errorf("type %s (%d): %w", "MD", dnsType, ErrNotSupportedType)
	case Type_MF:
		return nil, fmt.Errorf("type %s (%d): %w", "MF", dnsType, ErrNotSupportedType)
	case Type_CNAME:
		return nil, fmt.Errorf("type %s (%d): %w", "CNAME", dnsType, ErrNotSupportedType)
	case Type_SOA:
		return nil, fmt.Errorf("type %s (%d): %w", "SOA", dnsType, ErrNotSupportedType)
	case Type_MB:
		return nil, fmt.Errorf("type %s (%d): %w", "MB", dnsType, ErrNotSupportedType)
	case Type_MG:
		return nil, fmt.Errorf("type %s (%d): %w", "MG", dnsType, ErrNotSupportedType)
	case Type_MR:
		return nil, fmt.Errorf("type %s (%d): %w", "MR", dnsType, ErrNotSupportedType)
	case Type_NULL:
		return nil, fmt.Errorf("type %s (%d): %w", "NULL", dnsType, ErrNotSupportedType)
	case Type_WKS:
		return nil, fmt.Errorf("type %s (%d): %w", "WKS", dnsType, ErrNotSupportedType)
	case Type_PTR:
		return nil, fmt.Errorf("type %s (%d): %w", "PTR", dnsType, ErrNotSupportedType)
	case Type_HINFO:
		return nil, fmt.Errorf("type %s (%d): %w", "HINFO", dnsType, ErrNotSupportedType)
	case Type_MINFO:
		return nil, fmt.Errorf("type %s (%d): %w", "MINFO", dnsType, ErrNotSupportedType)
	case Type_MX:
		return &TypeMX{}, nil
	case Type_TXT:
		return nil, fmt.Errorf("type %s (%d): %w", "TXT", dnsType, ErrNotSupportedType)
	case Type_AAAA:
		return nil, fmt.Errorf("type %s (%d): %w", "AAAA", dnsType, ErrNotSupportedType)
	case Type_SRV:
		return nil, fmt.Errorf("type %s (%d): %w", "SRV", dnsType, ErrNotSupportedType)
	case Type_NAPTR:
		return nil, fmt.Errorf("type %s (%d): %w", "NAPTR", dnsType, ErrNotSupportedType)
	case Type_OPT:
		return nil, fmt.Errorf("type %s (%d): %w", "OPT", dnsType, ErrNotSupportedType)
	case Type_DS:
		return nil, fmt.Errorf("type %s (%d): %w", "DS", dnsType, ErrNotSupportedType)
	case Type_RRSIG:
		return nil, fmt.Errorf("type %s (%d): %w", "RRSIG", dnsType, ErrNotSupportedType)
	case Type_NSEC:
		return nil, fmt.Errorf("type %s (%d): %w", "NSEC", dnsType, ErrNotSupportedType)
	case Type_DNSKEY:
		return nil, fmt.Errorf("type %s (%d): %w", "DNSKEY", dnsType, ErrNotSupportedType)
	case Type_TLSA:
		return nil, fmt.Errorf("type %s (%d): %w", "TLSA", dnsType, ErrNotSupportedType)
	case Type_SPF:
		return nil, fmt.Errorf("type %s (%d): %w", "SPF", dnsType, ErrNotSupportedType)
	case Type_AXFR:
		return nil, fmt.Errorf("type %s (%d): %w", "AXFR", dnsType, ErrNotSupportedType)
	case Type_MAILB:
		return nil, fmt.Errorf("type %s (%d): %w", "MAILB", dnsType, ErrNotSupportedType)
	case Type_MAILA:
		return nil, fmt.Errorf("type %s (%d): %w", "MAILA", dnsType, ErrNotSupportedType)
	case Type_ANY:
		return nil, fmt.Errorf("type %s (%d): %w", "*", dnsType, ErrNotSupportedType)
	}
	return nil, fmt.Errorf("type %d: %w", dnsType, ErrUnknownType)
}
