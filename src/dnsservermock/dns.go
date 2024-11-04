package dnsservermock

import (
	"errors"
	"fmt"
	"log"
	"net"
)

type DNSServer struct {
	conn *net.UDPConn
	ip   net.IP
	port int
}

func NewDnsServer(ip net.IP, port int) *DNSServer {
	return &DNSServer{
		conn: nil,
		ip:   ip,
		port: port,
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
		ds.sendErrorResponse(req, clientAddr, RcodeFormErr, err)
		return
	}

	fmt.Printf("%#v\n", req) // TEST

	resp := &DNSResponse{}
	resp.CopyHeaderAndQuestions(req)
	for _, q := range req.Questions {
		proc, err := GetProcess(DnsType(q.Type))
		if err != nil {
			switch {
			case errors.Is(err, ErrNotSupportedType):
				ds.sendErrorResponse(req, clientAddr, RcodeNotImp, err)
			case errors.Is(err, ErrUnknownType):
				ds.sendErrorResponse(req, clientAddr, RcodeFormErr, err)
			default:
				ds.sendErrorResponse(req, clientAddr, RcodeServFail, err)
			}
			return
		}
		proc.Process(req, resp, q)
	}

	if err := ds.sendResponse(resp, clientAddr); err != nil {
		log.Printf("Error sending DNS Response (ID: %04X): %s", resp.ID, err)
		return
	}
	log.Printf("DNS Response have been sent (ID: %04X)", resp.ID)
}

func (ds *DNSServer) sendErrorResponse(req *DNSRequest, clientAddr *net.UDPAddr, rcode Rcode, err error) {
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
	_, err := ds.conn.WriteToUDP(respBuf, clientAddr)
	if err != nil {
		return err
	}
	return nil
}
