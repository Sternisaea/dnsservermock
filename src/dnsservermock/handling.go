package dnsservermock

import (
	"errors"

	"github.com/Sternisaea/dnsservermock/src/dnsconst"
	"github.com/Sternisaea/dnsservermock/src/dnsstorage"
)

type DnsHandling struct {
	ID       uint16
	query    *DnsQuery
	response *DnsResponse
	output   []byte
}

func NewDnsHandling() *DnsHandling {
	return &DnsHandling{}
}

func (dh *DnsHandling) ReadingQuery(buf []byte, n int) error {
	(*dh).query = &DnsQuery{}
	(*dh).ID = (*dh).query.ID
	if err := (*dh).query.ProcessRequestBuffer(buf, n); err != nil {
		(*dh).response = NewDnsResponse((*dh).query.ID)
		(*dh).response.Flags.RCODE = dnsconst.RcodeFormErr
		return err
	}
	return nil
}

func (dh *DnsHandling) CreateResponse() error {
	(*dh).response = NewDnsResponse((*dh).query.ID)
	(*dh).response.Questions = append((*dh).response.Questions, (*dh).query.Questions...)

	for _, q := range (*dh).response.Questions {
		dres, err := GetResource(q)
		if err != nil {
			switch {
			case errors.Is(err, ErrNotSupportedType):
				(*dh).response.Flags.RCODE = dnsconst.RcodeNotImp
			case errors.Is(err, ErrUnknownType):
				(*dh).response.Flags.RCODE = dnsconst.RcodeFormErr
			default:
				(*dh).response.Flags.RCODE = dnsconst.RcodeServFail
			}
			return err
		}
		(*dh).response.Answers = append((*dh).response.Answers, dres)
	}

	// Authorities
	// Additionals

	return nil
}

func (dh *DnsHandling) ExecuteQueries(store dnsstorage.Storage) error {
	for _, dres := range (*dh).response.Answers {
		rcode, err := dres.Query(store)
		(*dh).response.Flags.RCODE = rcode
		if err != nil {
			return err
		}
	}
	return nil
}

func (dh *DnsHandling) WriteResponse() error {
	(*dh).output = (*dh).response.Write()
	if len((*dh).output) == 0 {
		return errors.New("empty reponse output")
	}
	return nil
}

func (dh *DnsHandling) GetOutput() []byte {
	return (*dh).output
}
