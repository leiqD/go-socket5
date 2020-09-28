package repository

import (
	"github.com/leiqD/go-socket5/domain/model"
	"net"
)

type tcpConnRepository struct {
}

type TcpConnRepository interface {
	NewSession(conn net.Conn)
	GetSession() []model.CtrlSession
	CloseSession(session model.CtrlSession)
	CloseByConnectId(connectId model.ConnectId)
}

func NewTcpConnRepository() TcpConnRepository {
	return &tcpConnRepository{}
}

func (p *tcpConnRepository) NewSession(conn net.Conn) {

}

func (p *tcpConnRepository) GetSession() (res []model.CtrlSession) {
	return
}

func (p *tcpConnRepository) CloseSession(session model.CtrlSession) {

}
func (p *tcpConnRepository) CloseByConnectId(connectId model.ConnectId) {

}
