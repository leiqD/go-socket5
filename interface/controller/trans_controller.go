package controller

import (
	"github.com/leiqD/go-socket5/usecase/interactor"
)

type transController struct {
	tcpTransIneracotr   interactor.TcpTransInteractor
	tcpSessionIntractor interactor.TcpConnInterfactor
}

type TransController interface {
	Run()
	Stop()
}

func NewTransController(us interactor.TcpTransInteractor, tcpSessionIntractor interactor.TcpConnInterfactor) TransController {
	return &transController{
		tcpTransIneracotr:   us,
		tcpSessionIntractor: tcpSessionIntractor,
	}
}

func (p *transController) Run() {
	waitTrans := p.tcpSessionIntractor.GetSessionTcpWaitTrans()
	if len(waitTrans) == 0 {
		return
	}
	p.tcpTransIneracotr.Handle(waitTrans)
	for _, session := range waitTrans {
		p.tcpSessionIntractor.SetSessionTrans(&session)
	}
}

func (p *transController) Stop() {

}
