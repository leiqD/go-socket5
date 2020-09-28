package interactor

import (
	"github.com/leiqD/go-socket5/domain/model"
	ip "github.com/leiqD/go-socket5/interface/presenter"
	ir "github.com/leiqD/go-socket5/interface/repository"
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
	GetSession() []model.CtrlSession
	Close(model.CtrlSession)
	CloseByConnectId(connectId model.ConnectId)
}

func NewTcpConnInterfactor(r ir.TcpConnRepository, p ip.TcpConnPresenter) TcpConnInterfactor {
	return &tcpConnInterfactor{
		TcpConnReponsitory: r,
		TcpConnPresenter:   p,
	}
}

func (p *tcpConnInterfactor) Connect(conn net.Conn) {
	p.TcpConnReponsitory.NewSession(conn)
}

func (p *tcpConnInterfactor) GetSession() []model.CtrlSession {
	return p.TcpConnReponsitory.GetSession()
}

func (p *tcpConnInterfactor) Close(session model.CtrlSession) {
	p.TcpConnReponsitory.CloseSession(session)
}

func (p *tcpConnInterfactor) CloseByConnectId(connectId model.ConnectId) {
	p.TcpConnReponsitory.CloseByConnectId(connectId)
}
