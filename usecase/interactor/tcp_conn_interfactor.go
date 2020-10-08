package interactor

import (
	"github.com/leiqD/go-socket5/domain/model"
	up "github.com/leiqD/go-socket5/usecase/presenter"
	ur "github.com/leiqD/go-socket5/usecase/repository"
	"net"
)

type tcpConnInterfactor struct {
	TcpConnReponsitory ur.TcpConnRepository
	TcpConnPresenter   up.TcpConnPresenter
}

type TcpConnInterfactor interface {
	NewSession(conn net.Conn)
	GetSessionWaitNeg() []model.CtrlSession
	Close(model.CtrlSession)
	CloseByConnectId(connectId model.ConnectId)
	RemoveSession(session *model.CtrlSession)
	SetSessionDoing(session *model.CtrlSession)
}

func NewTcpConnInterfactor(r ur.TcpConnRepository, p up.TcpConnPresenter) TcpConnInterfactor {
	return &tcpConnInterfactor{
		TcpConnReponsitory: r,
		TcpConnPresenter:   p,
	}
}

func (p *tcpConnInterfactor) NewSession(conn net.Conn) {
	p.TcpConnReponsitory.NewSession(conn)
}

func (p *tcpConnInterfactor) GetSessionWaitNeg() []model.CtrlSession {
	return p.TcpConnReponsitory.GetWithoutDoing()
}

func (p *tcpConnInterfactor) Close(session model.CtrlSession) {
	p.TcpConnReponsitory.CloseSession(session)
}

func (p *tcpConnInterfactor) CloseByConnectId(connectId model.ConnectId) {
	p.TcpConnReponsitory.CloseByConnectId(connectId)
}

func (p *tcpConnInterfactor) RemoveSession(session *model.CtrlSession) {
	p.TcpConnReponsitory.RemoveSession(session.GetId())
}

func (p *tcpConnInterfactor) SetSessionDoing(session *model.CtrlSession) {
	p.TcpConnReponsitory.SetDoing(session.GetId())
}
