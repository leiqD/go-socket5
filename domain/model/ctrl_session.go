package model

import (
	"github.com/leiqD/go-socket5/infrastructure/idgenerator"
	"net"
	"time"
)

type SessionSetup int8

const (
	SetupAccept = iota + 1
	SetupNegotiate
	SetupTrans
	SetupDoing
)

type ProtocolType int8

const (
	ProtocolTcp = iota
	ProtocolUdp
)

type CtrlSession struct {
	conn       net.Conn
	id         ConnectId
	updateTime int64
	dest       net.Conn
	setup      SessionSetup
	protocol   ProtocolType
}

func NewCtrlSession(conn net.Conn) *CtrlSession {
	now := time.Now().Unix()
	session := &CtrlSession{
		conn:       conn,
		updateTime: now,
		id:         ConnectId(idgenerator.GetId()),
		setup:      SetupAccept,
	}
	session.conn.SetReadDeadline(time.Now().Add(1000 * time.Second))
	session.conn.SetWriteDeadline(time.Now().Add(1000 * time.Second))
	return session
}

func (session CtrlSession) GetConn() net.Conn {
	return session.conn
}

func (session CtrlSession) GetId() ConnectId {
	return session.id
}

func (session CtrlSession) GetUpdateTime() int64 {
	return session.updateTime
}

func (session *CtrlSession) Flush() {
	session.updateTime = time.Now().Unix()
}

func (session CtrlSession) GetRemoteAddrContent() string {
	return session.conn.RemoteAddr().String()
}

func (session *CtrlSession) CloseConn() {
	session.conn.Close()
}

func (session *CtrlSession) SetDestConn(dst net.Conn) {
	session.dest = dst
}

func (session *CtrlSession) GetDestConn() net.Conn {
	return session.dest
}

func (session *CtrlSession) SetSetup(setup SessionSetup) {
	session.setup = setup
}

func (session *CtrlSession) GetSetup() SessionSetup {
	return session.setup
}

func (session *CtrlSession) SetProtocol(protocol ProtocolType) {
	session.protocol = protocol
}

func (session *CtrlSession) GetProtocol() ProtocolType {
	return session.protocol
}
