package controller

import (
	"github.com/leiqD/go-socket5/usecase/interactor"
	"net"
)

type transController struct {
	connInterfactor interactor.TcpConnInterfactor
}

type TransController interface {
	Run(conn net.Conn)
	Stop()
}

func NewTransController(us interactor.TcpConnInterfactor) TransController {
	return &transController{
		connInterfactor: us,
	}
}

func (p *transController) Run(conn net.Conn) {
	//p.tcpConn.Connect(conn)
	p.connInterfactor.Connect(conn)
}

func (p *transController) Stop() {

}
