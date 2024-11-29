package dnsservermock

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/Sternisaea/dnsservermock/src/dnsstorage"
)

type DnsServer struct {
	conn    *net.UDPConn
	ip      net.IP
	port    int
	storage dnsstorage.Storage
}

func NewDnsServer(ip net.IP, port int, storage dnsstorage.Storage) *DnsServer {
	return &DnsServer{
		ip:      ip,
		port:    port,
		conn:    nil,
		storage: storage,
	}
}

func (ds *DnsServer) Start() error {
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

func (ds *DnsServer) Stop() error {
	if (*ds).conn != nil {
		if err := (*ds).conn.Close(); err != nil {
			log.Printf("Error while stopping DNS server: %s", err)
			return err
		}
	}
	log.Printf("DNS server closed")
	return nil
}

func (ds *DnsServer) handleRequest(buf []byte, n int, clientAddr *net.UDPAddr) {
	dh := NewDnsHandling()
	if err := (*ds).processRequest(dh, buf, n); err != nil {
		log.Println(err)
	}

	if err := ds.writeResponse(dh, clientAddr); err != nil {
		log.Printf("Error sending DNS response (ID: %04x): %s", (*dh).ID, err)
	}
	log.Printf("DNS response was succesfully sent (ID: %04x)", (*dh).ID)
}

func (ds *DnsServer) processRequest(dh *DnsHandling, buf []byte, n int) error {
	if err := dh.ReadingQuery(buf, n); err != nil {
		return fmt.Errorf("error reading DNS response (ID: %04x): %s", (*dh).ID, err)
	}
	if err := dh.CreateResponse(); err != nil {
		return fmt.Errorf("error creating DNS response (ID: %04x): %s", (*dh).ID, err)
	}
	if err := dh.ExecuteQueries((*ds).storage); err != nil {
		return fmt.Errorf("error querying DNS response (ID: %04x): %s", (*dh).ID, err)
	}
	return nil
}

func (ds *DnsServer) writeResponse(dh *DnsHandling, clientAddr *net.UDPAddr) error {
	if err := dh.WriteResponse(); err != nil {
		return err
	}

	_, err := ds.conn.WriteToUDP((*dh).GetOutput(), clientAddr)
	if err != nil {
		return err
	}
	return nil
}
