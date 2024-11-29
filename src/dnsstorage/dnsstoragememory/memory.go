package dnsstoragememory

import (
	"github.com/Sternisaea/dnsservermock/src/dnsconst"
	"github.com/Sternisaea/dnsservermock/src/dnsstorage"
)

type MemoryStore struct {
	domains map[string]resolutions
}

func (ms *MemoryStore) initialiseDomains() {
	(*ms).domains = make(map[string]resolutions)
}

type resolutions map[dnsconst.DnsType]string

func NewMemoryStore() *MemoryStore {
	ms := &MemoryStore{}
	(*ms).initialiseDomains()
	return ms
}

func (ms *MemoryStore) Set(domain string, reqType dnsconst.DnsType, result string) error {
	if (*ms).domains == nil {
		(*ms).initialiseDomains()
	}
	res, ok := (*ms).domains[domain]
	if !ok {
		res = make(map[dnsconst.DnsType]string)
	}
	res[reqType] = result
	(*ms).domains[domain] = res
	return nil
}

func (ms *MemoryStore) Get(domain string, reqType dnsconst.DnsType) (string, error) {
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
