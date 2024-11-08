package dnsconst

type DnsType uint16

const (
	Type_A     DnsType = 1 // RFC 1035
	Type_NS    DnsType = 2
	Type_MD    DnsType = 3
	Type_MF    DnsType = 4
	Type_CNAME DnsType = 5
	Type_SOA   DnsType = 6
	Type_MB    DnsType = 7
	Type_MG    DnsType = 8
	Type_MR    DnsType = 9
	Type_NULL  DnsType = 10
	Type_WKS   DnsType = 11
	Type_PTR   DnsType = 12
	Type_HINFO DnsType = 13
	Type_MINFO DnsType = 14
	Type_MX    DnsType = 15
	Type_TXT   DnsType = 16

	Type_AAAA   DnsType = 28 // RFC 3596
	Type_SRV    DnsType = 33 // RFC 2782
	Type_NAPTR  DnsType = 35 // RFC 2915
	Type_OPT    DnsType = 41 // RFC 6891 (RFC 2671)
	Type_DS     DnsType = 43 // RFC 4034
	Type_RRSIG  DnsType = 46 // RFC 4034
	Type_NSEC   DnsType = 47 // RFC 4034
	Type_DNSKEY DnsType = 48 // RFC 4034
	Type_TLSA   DnsType = 52 // RFC 6698
	Type_SPF    DnsType = 99 // RFC 4408

	Type_AXFR  DnsType = 252
	Type_MAILB DnsType = 253
	Type_MAILA DnsType = 254
	Type_ANY   DnsType = 255
)

type DnsClass uint16

const (
	Class_IN DnsClass = 1
	Class_CS DnsClass = 2
	Class_CH DnsClass = 3
	Class_HS DnsClass = 4

	Class_ANY DnsClass = 255
)

type Rcode int

const (
	RcodeNoError  Rcode = 0  // No error condition
	RcodeFormErr  Rcode = 1  // Format error
	RcodeServFail Rcode = 2  // Server failure
	RcodeNXDomain Rcode = 3  // Non-Existent Domain
	RcodeNotImp   Rcode = 4  // Not Implemented
	RcodeRefused  Rcode = 5  // Refused
	RcodeYXDomain Rcode = 6  // Name Exists when it should not
	RcodeYXRRSet  Rcode = 7  // RR Set Exists when it should not
	RcodeNXRRSet  Rcode = 8  // RR Set that should exist does not
	RcodeNotAuth  Rcode = 9  // Not Authorized
	RcodeNotZone  Rcode = 10 // Name not contained in zone

	RcodeBADVERS  Rcode = 16 // Bad OPT Version or Bad Signature
	RcodeBADKEY   Rcode = 17 // Key not recognized
	RcodeBADTIME  Rcode = 18 // Signature out of time window
	RcodeBADMODE  Rcode = 19 // Bad TKEY Mode
	RcodeBADNAME  Rcode = 20 // Duplicate key name
	RcodeBADALG   Rcode = 21 // Algorithm not supported
	RcodeBADTRUNC Rcode = 22 // Bad truncation
)
