package model

import (
	"github.com/leiqD/go-socket5/infrastructure/idgenerator"
	"net"
	"time"
)

type CtrlSession struct {
	conn       net.Conn
	id         ConnectId
	updateTime int64
}

func NewCtrlSession(conn net.Conn) *CtrlSession {
	now := time.Now().Unix()
	session := &CtrlSession{
		conn:       conn,
		updateTime: now,
		id:         ConnectId(idgenerator.GetId()),
	}
	session.conn.SetReadDeadline(time.Now().Add(100 * time.Microsecond))
	session.conn.SetWriteDeadline(time.Now().Add(100 * time.Microsecond))
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

func (session *CtrlSession) Update() {
	session.updateTime = time.Now().Unix()
}

func (session CtrlSession) GetRemoteAddrContent() string {
	return session.conn.RemoteAddr().String()
}

func (session *CtrlSession) CloseConn() {
	session.conn.Close()
}
