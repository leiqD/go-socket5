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
	sessions := p.tcpSessionIntractor.GetSessionWaitNeg()
	if len(sessions) == 0 {
		return
	}
	succ, fail := p.ctrlSessionInteractor.Negotiate(sessions)
	for _, session := range fail {
		p.tcpSessionIntractor.CloseByConnectId(session.GetId())
	}
	for _, session := range succ {
		p.tcpSessionIntractor.UpdateSession(session)
	}
}
