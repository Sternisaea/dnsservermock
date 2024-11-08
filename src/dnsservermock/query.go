package dnsservermock

import (
	"encoding/binary"
	"fmt"
	"strings"
)

var (
	ErrBufferTooShort = fmt.Errorf("buffer too short")
)

type DnsQuery struct {
	ID          uint16
	Flags       DnsFlags
	QDCount     uint16
	ANCount     uint16
	NSCount     uint16
	ARCount     uint16
	Questions   []DnsQuestion
	Authorities []DnsResourceRecord
	Additionals []DnsResourceRecord
}

type DnsQuestion struct {
	Name  string
	Type  uint16
	Class uint16
}

type DnsResourceRecord struct {
	Name     string
	Type     uint16
	Class    uint16
	TTL      uint32
	RDLength uint16
	RData    []byte
}

func (d *DnsQuery) ProcessRequestBuffer(buf []byte, length int) error {
	if length < 12 {
		return fmt.Errorf("dns request: %w", ErrBufferTooShort)
	}

	(*d).ID = binary.BigEndian.Uint16(buf[0:2])
	(*d).Flags.Set(binary.BigEndian.Uint16(buf[2:4]))
	(*d).QDCount = binary.BigEndian.Uint16(buf[4:6])
	(*d).ANCount = binary.BigEndian.Uint16(buf[6:8])
	(*d).NSCount = binary.BigEndian.Uint16(buf[8:10])
	(*d).ARCount = binary.BigEndian.Uint16(buf[10:12])

	offset := 12
	for i := 0; i < int(d.QDCount); i++ {
		question, newOffset, err := parseDNSQuestion(buf, offset)
		if err != nil {
			return fmt.Errorf("dns request: %w", err)
		}
		(*d).Questions = append((*d).Questions, question)
		offset = newOffset
	}
	return nil
}

func parseDNSQuestion(buf []byte, offset int) (DnsQuestion, int, error) {
	var question DnsQuestion
	name, newOffset, err := parseDNSName(buf, offset)
	if err != nil {
		return question, offset, fmt.Errorf("question: %w", err)
	}
	question.Name = name
	if newOffset+4 > len(buf) {
		return question, offset, fmt.Errorf("question: %w", ErrBufferTooShort)
	}
	question.Type = binary.BigEndian.Uint16(buf[newOffset : newOffset+2])
	question.Class = binary.BigEndian.Uint16(buf[newOffset+2 : newOffset+4])
	return question, newOffset + 4, nil
}

func parseDNSName(buf []byte, offset int) (string, int, error) {
	var name []string
	for {
		if offset >= len(buf) {
			return "", offset, fmt.Errorf("name pointer: %w", ErrBufferTooShort)
		}
		length := int(buf[offset])
		if length == 0 {
			offset++
			break
		}
		// RFC 1035 - 4.1.4. Message compression
		if length&0xC0 == 0xC0 { // 2 most significant bits (11) indicate pointer offset
			if offset+1 >= len(buf) {
				return "", offset, fmt.Errorf("name: %w", ErrBufferTooShort)
			}
			ptr := int(binary.BigEndian.Uint16(buf[offset:offset+2]) & 0x3FFF) // Pointer offset (filter out 2 most significant bits)
			part, _, err := parseDNSName(buf, ptr)
			if err != nil {
				return "", offset, err
			}
			name = append(name, part)
			offset += 2
			break
		}
		if offset+length+1 > len(buf) {
			return "", offset, fmt.Errorf("name segment: %w", ErrBufferTooShort)
		}
		name = append(name, string(buf[offset+1:offset+1+length]))
		offset += length + 1
	}
	return strings.Join(name, "."), offset, nil
}
