package dnsstorage

import (
	"errors"

	"github.com/sternisaea/dnsservermock/src/dnstypes"
)

type Storage interface {
	Set(domain string, reqType dnstypes.DnsType, result string) error
	Get(domain string, reqType dnstypes.DnsType) (string, error)
}

var (
	ErrDomainNotFound      = errors.New("domain not found")
	ErrRequestTypeNotFound = errors.New("request type not found")
)
