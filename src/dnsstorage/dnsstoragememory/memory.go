package dnsstoragememory

import (
	"github.com/sternisaea/dnsservermock/src/dnsstorage"
	"github.com/sternisaea/dnsservermock/src/dnstypes"
)

type MemoryStore struct {
	domains map[string]resolutions
}

func (ms *MemoryStore) initialiseDomains() {
	(*ms).domains = make(map[string]resolutions)
}

type resolutions map[dnstypes.DnsType]string

func NewMemoryStore() *MemoryStore {
	ms := &MemoryStore{}
	(*ms).initialiseDomains()
	return ms
}

func (ms *MemoryStore) Set(domain string, reqType dnstypes.DnsType, result string) error {
	if (*ms).domains == nil {
		(*ms).initialiseDomains()
	}
	res, ok := (*ms).domains[domain]
	if !ok {
		res = make(map[dnstypes.DnsType]string)
	}
	res[reqType] = result
	(*ms).domains[domain] = res
	return nil
}

func (ms *MemoryStore) Get(domain string, reqType dnstypes.DnsType) (string, error) {
	if (*ms).domains == nil {
		return "", dnsstorage.ErrDomainNotFound
	}
	rs, ok := (*ms).domains[domain]
	if !ok {
		return "", dnsstorage.ErrDomainNotFound
	}
	result, ok := rs[reqType]
	if !ok {
		return "", dnsstorage.ErrRequestTypeNotFound
	}
	return result, nil
}
