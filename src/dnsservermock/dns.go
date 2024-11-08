package dnsservermock

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/sternisaea/dnsservermock/src/dnsconst"
	"github.com/sternisaea/dnsservermock/src/dnsstorage"
)

type DNSServer struct {
	conn *net.UDPConn
	ip   net.IP
	port int

	storage dnsstorage.Storage
}

func NewDnsServer(ip net.IP, port int, storage dnsstorage.Storage) *DNSServer {
	return &DNSServer{
		conn: nil,
		ip:   ip,
		port: port,

		storage: storage,
	}
}

func (ds *DNSServer) Start() error {
	addr := net.UDPAddr{
		Port: (*ds).port,
		IP:   (*ds).ip,
	}

	var err error
	(*ds).conn, err = net.ListenUDP("udp", &addr)
	if err != nil {
		return err
	}

	log.Printf("DNS server listening on %s UDP port %d", (*ds).ip, (*ds).port)
	go func() {
		defer (*ds).Stop()
		for {
			buf := make([]byte, 4096)
			n, clientAddr, err := (*ds).conn.ReadFromUDP(buf)
			if err != nil {
				if errors.Is(err, net.ErrClosed) {
					return
				}
				log.Printf("Error reading from UDP: %s", err)
				continue
			}
			go (*ds).handleRequest(buf[:n], n, clientAddr)
		}
	}()

	return nil
}

func (ds *DNSServer) Stop() error {
	if (*ds).conn != nil {
		if err := (*ds).conn.Close(); err != nil {
			log.Printf("Error while stopping DNS server: %s", err)
			return err
		}
	}
	log.Printf("DNS server closed")
	return nil
}

func (ds *DNSServer) handleRequest(buf []byte, n int, clientAddr *net.UDPAddr) {
	req := &DNSRequest{}
	if err := req.ProcessRequestBuffer(buf, n); err != nil {
		ds.sendErrorResponse(req, clientAddr, dnsconst.RcodeFormErr, err)
		return
	}

	resp := NewDnsResponse()
	resp.CopyHeaderAndQuestions(req)
	for _, q := range req.Questions {
		proc, err := GetProcess(dnsconst.DnsType(q.Type))
		if err != nil {
			switch {
			case errors.Is(err, ErrNotSupportedType):
				ds.sendErrorResponse(req, clientAddr, dnsconst.RcodeNotImp, err)
			case errors.Is(err, ErrUnknownType):
				ds.sendErrorResponse(req, clientAddr, dnsconst.RcodeFormErr, err)
			default:
				ds.sendErrorResponse(req, clientAddr, dnsconst.RcodeServFail, err)
			}
			return
		}
		proc.Process(req, resp, q, (*ds).storage)
	}

	if err := ds.sendResponse(resp, clientAddr); err != nil {
		log.Printf("Error sending DNS Response (ID: %04X): %s", resp.ID, err)
		return
	}
	log.Printf("DNS Response have been sent (ID: %04X)", resp.ID)
}

func (ds *DNSServer) sendErrorResponse(req *DNSRequest, clientAddr *net.UDPAddr, rcode dnsconst.Rcode, err error) {
	log.Printf("DNS error (ID: %04X): %s", req.ID, err)

	resp := &DNSResponse{}
	resp.CopyHeaderAndQuestions(req)
	resp.Flags.RCODE = rcode
	if err := (*ds).sendResponse(resp, clientAddr); err != nil {
		log.Printf("Error sending DNS response error (ID: %04X): %s", resp.ID, err)
		return
	}
}

func (ds *DNSServer) sendResponse(resp *DNSResponse, clientAddr *net.UDPAddr) error {
	respBuf := resp.SerializeResponse()

	printbuffer(respBuf)

	_, err := ds.conn.WriteToUDP(respBuf, clientAddr)
	if err != nil {
		return err
	}
	return nil
}

func printbuffer(buf []byte) {
	fmt.Println("Buffer:")
	chrs := ""
	for i, b := range buf {
		if i != 0 && i%16 == 0 {
			if len(chrs) != 0 {
				fmt.Printf("%s", chrs)
			}
			fmt.Println()
			chrs = ""
		}

		// if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') || b == '-' {
		if b >= ' ' && b <= '~' {
			chrs += string(b)
		} else {
			chrs += "."
		}

		fmt.Printf("%02x ", b)
	}
	fmt.Println()
}
