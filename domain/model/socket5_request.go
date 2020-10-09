package model

import (
	"encoding/binary"
	"fmt"
	"github.com/leiqD/go-socket5/domain/entity"
	"github.com/leiqD/go-socket5/infrastructure/logger"
	"net"
)

const (
	DestAddrIpaddrV4 = 1
	DestAddrHostName = 3
	DestAddrIpaddrV6 = 4
)

const (
	CONNECT       = 1
	BIND          = 2
	UDP_ASSOCIATE = 3
)

type request struct {
	req entity.Request
}

type Request interface {
	Parse(fromClient []byte) error
	Pack() []byte
	PackUdp(ip []byte) []byte
	ConnectTcpDst() (conn net.Conn, err error)
	IsTcp() bool
	IsUdp() bool
	Cmd() byte
	Port() int64
}

func NewRequest() Request {
	d := &request{}
	return d
}

func (p *request) Parse(fromClient []byte) error {
	p.req.C2s.Ver = fromClient[0]
	p.req.C2s.Cmd = fromClient[1]
	p.req.C2s.Rsv = fromClient[2]
	p.req.C2s.Atyp = fromClient[3]
	switch p.req.C2s.Atyp {
	case DestAddrIpaddrV4:
		addrLen := byte(4)
		p.req.C2s.Address = make([]byte, addrLen)
		for i := byte(0); i < addrLen; i++ {
			p.req.C2s.Address[i] = fromClient[i+4]
		}
		p.req.C2s.Port = make([]byte, 2)
		p.req.C2s.Port[0] = fromClient[4+4]
		p.req.C2s.Port[1] = fromClient[4+4+1]
		port := int64(binary.BigEndian.Uint16(p.req.C2s.Port))
		p.req.C2s.Host = fmt.Sprintf("%d.%d.%d.%d:%d", p.req.C2s.Address[0], p.req.C2s.Address[1], p.req.C2s.Address[2], p.req.C2s.Address[3], port)
		logger.Infof("%v:%v,host=%s", p.req.C2s.Address, p.req.C2s.Port, p.req.C2s.Host)
	case DestAddrHostName:
		addrLen := fromClient[4]
		p.req.C2s.Address = make([]byte, addrLen)
		for i := byte(0); i < addrLen; i++ {
			p.req.C2s.Address[i] = fromClient[i+5]
		}
		p.req.C2s.Port = make([]byte, 2)
		p.req.C2s.Port[0] = fromClient[5+addrLen]
		p.req.C2s.Port[1] = fromClient[5+addrLen+1]
		addRess := string(p.req.C2s.Address)
		port := int64(binary.BigEndian.Uint16(p.req.C2s.Port))
		//ip, _ := net.LookupHost()
		p.req.C2s.Host = fmt.Sprintf("%s:%d", addRess, port)
		logger.Infof("%v:%v,host=%s", p.req.C2s.Address, p.req.C2s.Port, p.req.C2s.Host)
	case DestAddrIpaddrV6:

	}

	return nil
}

func (p *request) IsTcp() bool {
	return p.req.C2s.Cmd == CONNECT
}

func (p *request) IsUdp() bool {
	return p.req.C2s.Cmd == UDP_ASSOCIATE
}

func (p *request) ConnectTcpDst() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", p.req.C2s.Host)

	logger.Infof("ConnectTcpDst Host=%s err=%v", p.req.C2s.Host, err)
	return
}

func (p *request) Pack() []byte {
	buff := []byte{}
	buff = append(buff, 5)
	buff = append(buff, 0)
	buff = append(buff, 0)
	buff = append(buff, 1)
	buff = append(buff, 0xc0)
	buff = append(buff, 0xa8)
	buff = append(buff, 0x01)
	buff = append(buff, 0x04)
	buff = append(buff, 0)
	buff = append(buff, 0)
	return buff
}

func (p *request) Cmd() byte {
	return p.req.C2s.Cmd
}

func (p *request) Port() int64 {
	port := int64(binary.BigEndian.Uint16(p.req.C2s.Port))
	return port
}

func (p *request) PackUdp(ip []byte) []byte {
	buff := []byte{}
	buff = append(buff, 5)
	buff = append(buff, 0)
	buff = append(buff, 0)
	buff = append(buff, 1)
	buff = append(buff, ip[0])
	buff = append(buff, ip[1])
	buff = append(buff, ip[2])
	buff = append(buff, ip[3])
	buff = append(buff, p.req.C2s.Port[0])
	buff = append(buff, p.req.C2s.Port[1])
	return buff
}
