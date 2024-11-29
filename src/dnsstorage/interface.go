package dnsstorage

import (
	"errors"

	"github.com/Sternisaea/dnsservermock/src/dnsconst"
)

type Storage interface {
	Set(domain string, reqType dnsconst.DnsType, result string) error
	Get(domain string, reqType dnsconst.DnsType) (string, error)
}

var (
	ErrDomainNotFound      = errors.New("domain not found")
	ErrRequestTypeNotFound = errors.New("request type not found")
)
