package dnsservermock

import (
	"bytes"
	"encoding/binary"
	"strings"
)

func writeDomainName(buf *bytes.Buffer, dms *domains, name string) {
	parts := strings.Split(name, ".")
	for i, part := range parts {
		dm := strings.Join(parts[i:], ".")
		if pos, ok := (*dms)[dm]; ok {
			// Write pointer offset
			binary.Write(buf, binary.BigEndian, uint16(pos)|0xC000)
			return
		} else {
			// Write name
			(*dms)[dm] = buf.Len()
			buf.WriteByte(byte(len(part)))
			buf.WriteString(part)
		}
	}
	buf.WriteByte(0)
}
