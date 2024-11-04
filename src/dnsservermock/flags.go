package dnsservermock

type DNSFlags struct {
	QR     bool  // false = query, true = response
	Opcode uint8 // Kind of query (4-bits)
	AA     bool  // Authoritative Answer
	TC     bool  // Truncation
	RD     bool  // Recursion Desired
	RA     bool  // Recursion Available
	Z      bool  // Not used
	AD     bool  // Authentic Data
	CD     bool  // Checking Disabled
	RCODE  Rcode // Reponse Code (4-bits)
}

func (df *DNSFlags) Set(flags uint16) {
	df.QR = (flags & 0x8000) != 0
	df.Opcode = uint8((flags & 0x7800) >> 11)
	df.AA = (flags & 0x0400) != 0
	df.TC = (flags & 0x0200) != 0
	df.RD = (flags & 0x0100) != 0
	df.RA = (flags & 0x0080) != 0
	df.Z = (flags & 0x0040) != 0
	df.AD = (flags & 0x0020) != 0
	df.CD = (flags & 0x0010) != 0
	df.RCODE = Rcode(flags & 0x000F)
}

func (df *DNSFlags) Get() uint16 {
	var flags uint16
	if df.QR {
		flags |= 0x8000
	}
	flags |= uint16(df.Opcode) << 11
	if df.AA {
		flags |= 0x0400
	}
	if df.TC {
		flags |= 0x0200
	}
	if df.RD {
		flags |= 0x0100
	}
	if df.RA {
		flags |= 0x0080
	}
	if df.Z {
		flags |= 0x0040
	}
	if df.AD {
		flags |= 0x0020
	}
	if df.CD {
		flags |= 0x0010
	}
	flags |= uint16(df.RCODE)
	return flags
}
