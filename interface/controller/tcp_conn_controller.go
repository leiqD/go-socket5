package controller

import (
	"github.com/leiqD/go-socket5/usecase/interactor"
	"net"
)

type tcpConnController struct {
	connInterfactor interactor.TcpConnInterfactor
}

type TcpConnController interface {
	NewSession(conn net.Conn)
	GetConnActor() interactor.TcpConnInterfactor
}

func NewTcpConnController(us interactor.TcpConnInterfactor) TcpConnController {
	return &tcpConnController{
		connInterfactor: us,
	}
}

func (p *tcpConnController) NewSession(conn net.Conn) {
	p.connInterfactor.Connect(conn)
}

func (p *tcpConnController) GetConnActor() interactor.TcpConnInterfactor {
	return p.connInterfactor
}
