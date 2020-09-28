package controller

import (
	"github.com/leiqD/go-socket5/usecase/interactor"
)

type ctrlSessionController struct {
	tcpSessionIntractor   interactor.TcpConnInterfactor
	ctrlSessionInteractor interactor.CtrlSessionInteractor
}

type CtrlSessionController interface {
	Negotiate()
}

func NewCtrlSessionController(session interactor.CtrlSessionInteractor, connActor interactor.TcpConnInterfactor) CtrlSessionController {
	return &ctrlSessionController{
		ctrlSessionInteractor: session,
		tcpSessionIntractor:   connActor,
	}
}

func (p *ctrlSessionController) Negotiate() {
	sessions := p.tcpSessionIntractor.GetSession()
	if len(sessions) == 0 {
		return
	}
	_, fail := p.ctrlSessionInteractor.Negotiate(sessions)
	for _, connectId := range fail {
		p.tcpSessionIntractor.CloseByConnectId(connectId)
	}
}
