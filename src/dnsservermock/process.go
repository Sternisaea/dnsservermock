package dnsservermock

import (
	"errors"
	"fmt"

	"github.com/sternisaea/dnsservermock/src/dnsstorage"
	"github.com/sternisaea/dnsservermock/src/dnstypes"
)

var (
	ErrNotSupportedType = errors.New("type not supported")
	ErrUnknownType      = errors.New("unknown type ")
)

type TypeProcess interface {
	Process(req *DNSRequest, resp *DNSResponse, qst DNSQuestion, store dnsstorage.Storage)
}

func GetProcess(dnsType dnstypes.DnsType) (TypeProcess, error) {
	switch dnsType {
	case dnstypes.Type_A:
		return &TypeA{}, nil
	case dnstypes.Type_NS:
		return nil, fmt.Errorf("type %s (%d): %w", "NS", dnsType, ErrNotSupportedType)
	case dnstypes.Type_MD:
		return nil, fmt.Errorf("type %s (%d): %w", "MD", dnsType, ErrNotSupportedType)
	case dnstypes.Type_MF:
		return nil, fmt.Errorf("type %s (%d): %w", "MF", dnsType, ErrNotSupportedType)
	case dnstypes.Type_CNAME:
		return nil, fmt.Errorf("type %s (%d): %w", "CNAME", dnsType, ErrNotSupportedType)
	case dnstypes.Type_SOA:
		return nil, fmt.Errorf("type %s (%d): %w", "SOA", dnsType, ErrNotSupportedType)
	case dnstypes.Type_MB:
		return nil, fmt.Errorf("type %s (%d): %w", "MB", dnsType, ErrNotSupportedType)
	case dnstypes.Type_MG:
		return nil, fmt.Errorf("type %s (%d): %w", "MG", dnsType, ErrNotSupportedType)
	case dnstypes.Type_MR:
		return nil, fmt.Errorf("type %s (%d): %w", "MR", dnsType, ErrNotSupportedType)
	case dnstypes.Type_NULL:
		return nil, fmt.Errorf("type %s (%d): %w", "NULL", dnsType, ErrNotSupportedType)
	case dnstypes.Type_WKS:
		return nil, fmt.Errorf("type %s (%d): %w", "WKS", dnsType, ErrNotSupportedType)
	case dnstypes.Type_PTR:
		return nil, fmt.Errorf("type %s (%d): %w", "PTR", dnsType, ErrNotSupportedType)
	case dnstypes.Type_HINFO:
		return nil, fmt.Errorf("type %s (%d): %w", "HINFO", dnsType, ErrNotSupportedType)
	case dnstypes.Type_MINFO:
		return nil, fmt.Errorf("type %s (%d): %w", "MINFO", dnsType, ErrNotSupportedType)
	case dnstypes.Type_MX:
		return &TypeMX{}, nil
	case dnstypes.Type_TXT:
		return nil, fmt.Errorf("type %s (%d): %w", "TXT", dnsType, ErrNotSupportedType)
	case dnstypes.Type_AAAA:
		return nil, fmt.Errorf("type %s (%d): %w", "AAAA", dnsType, ErrNotSupportedType)
	case dnstypes.Type_SRV:
		return nil, fmt.Errorf("type %s (%d): %w", "SRV", dnsType, ErrNotSupportedType)
	case dnstypes.Type_NAPTR:
		return nil, fmt.Errorf("type %s (%d): %w", "NAPTR", dnsType, ErrNotSupportedType)
	case dnstypes.Type_OPT:
		return nil, fmt.Errorf("type %s (%d): %w", "OPT", dnsType, ErrNotSupportedType)
	case dnstypes.Type_DS:
		return nil, fmt.Errorf("type %s (%d): %w", "DS", dnsType, ErrNotSupportedType)
	case dnstypes.Type_RRSIG:
		return nil, fmt.Errorf("type %s (%d): %w", "RRSIG", dnsType, ErrNotSupportedType)
	case dnstypes.Type_NSEC:
		return nil, fmt.Errorf("type %s (%d): %w", "NSEC", dnsType, ErrNotSupportedType)
	case dnstypes.Type_DNSKEY:
		return nil, fmt.Errorf("type %s (%d): %w", "DNSKEY", dnsType, ErrNotSupportedType)
	case dnstypes.Type_TLSA:
		return nil, fmt.Errorf("type %s (%d): %w", "TLSA", dnsType, ErrNotSupportedType)
	case dnstypes.Type_SPF:
		return nil, fmt.Errorf("type %s (%d): %w", "SPF", dnsType, ErrNotSupportedType)
	case dnstypes.Type_AXFR:
		return nil, fmt.Errorf("type %s (%d): %w", "AXFR", dnsType, ErrNotSupportedType)
	case dnstypes.Type_MAILB:
		return nil, fmt.Errorf("type %s (%d): %w", "MAILB", dnsType, ErrNotSupportedType)
	case dnstypes.Type_MAILA:
		return nil, fmt.Errorf("type %s (%d): %w", "MAILA", dnsType, ErrNotSupportedType)
	case dnstypes.Type_ANY:
		return nil, fmt.Errorf("type %s (%d): %w", "*", dnsType, ErrNotSupportedType)
	}
	return nil, fmt.Errorf("type %d: %w", dnsType, ErrUnknownType)
}
