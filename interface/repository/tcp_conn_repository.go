package repository

import (
	"github.com/leiqD/go-socket5/domain/model"
	"net"
)

type tcpConnRepository struct {
}

type TcpConnRepository interface {
	NewSession(conn net.Conn)
	GetSessionWaitNeg() []model.CtrlSession
	CloseSession(model.CtrlSession)
	CloseByConnectId(connectId model.ConnectId)
	UpdateSession(session *model.CtrlSession)
	GetSessionWaitTrans() []model.CtrlSession
	SetSessionTrans(session *model.CtrlSession)
}

func NewTcpConnRepository() TcpConnRepository {
	return &tcpConnRepository{}
}

func (p *tcpConnRepository) NewSession(conn net.Conn) {

}

func (p *tcpConnRepository) GetSessionWaitNeg() (res []model.CtrlSession) {
	return
}

func (p *tcpConnRepository) CloseSession(session model.CtrlSession) {

}
func (p *tcpConnRepository) CloseByConnectId(connectId model.ConnectId) {

}

func (p *tcpConnRepository) UpdateSession(session *model.CtrlSession) {

}

func (p *tcpConnRepository) GetSessionWaitTrans() (res []model.CtrlSession) {
	return
}

func (p *tcpConnRepository) SetSessionTrans(session *model.CtrlSession) {

}
