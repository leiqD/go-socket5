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
	NegotiateSocket5() error
	Stop()
}

func NewTcpConnController(us interactor.TcpConnInterfactor) TcpConnController {
	return &tcpConnController{
		connInterfactor: us,
	}
}

func (p *tcpConnController) NewSession(conn net.Conn) {
	//p.tcpConn.Connect(conn)
	p.connInterfactor.Connect(conn)
}

func (p *tcpConnController) NegotiateSocket5() error {
	return p.connInterfactor.NegotiateSocket5()
}

func (p *tcpConnController) Stop() {

}
