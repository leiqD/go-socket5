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
	Connect(conn net.Conn)
	GetSessionWaitNeg() []model.CtrlSession
	Close(model.CtrlSession)
	CloseByConnectId(connectId model.ConnectId)
	UpdateSession(session *model.CtrlSession)
	GetSessionTcpWaitTrans() []model.CtrlSession
	SetSessionTrans(session *model.CtrlSession)
}

func NewTcpConnInterfactor(r ur.TcpConnRepository, p up.TcpConnPresenter) TcpConnInterfactor {
	return &tcpConnInterfactor{
		TcpConnReponsitory: r,
		TcpConnPresenter:   p,
	}
}

func (p *tcpConnInterfactor) Connect(conn net.Conn) {
	p.TcpConnReponsitory.NewSession(conn)
}

func (p *tcpConnInterfactor) GetSessionWaitNeg() []model.CtrlSession {
	return p.TcpConnReponsitory.GetSessionWaitNeg()
}

func (p *tcpConnInterfactor) GetSessionTcpWaitTrans() []model.CtrlSession {
	return p.TcpConnReponsitory.GetSessionTcpWaitTrans()
}

func (p *tcpConnInterfactor) Close(session model.CtrlSession) {
	p.TcpConnReponsitory.CloseSession(session)
}

func (p *tcpConnInterfactor) CloseByConnectId(connectId model.ConnectId) {
	p.TcpConnReponsitory.CloseByConnectId(connectId)
}

func (p *tcpConnInterfactor) UpdateSession(session *model.CtrlSession) {
	p.TcpConnReponsitory.Update(session)
}
func (p *tcpConnInterfactor) SetSessionTrans(session *model.CtrlSession) {
	p.TcpConnReponsitory.SetSessionSetupTrnas(session)
}
