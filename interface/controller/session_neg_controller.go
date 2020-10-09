package controller

import (
	"github.com/leiqD/go-socket5/usecase/interactor"
)

type sessionNegController struct {
	tcpSessionIntractor  interactor.TcpConnInterfactor
	sessionNegInteractor interactor.SessionNegInteractor
}

type SessionNegController interface {
	Negotiate()
	GetSessionNegInteractor() interactor.SessionNegInteractor
}

func NewCtrlSessionController(session interactor.SessionNegInteractor, connActor interactor.TcpConnInterfactor) SessionNegController {
	return &sessionNegController{
		sessionNegInteractor: session,
		tcpSessionIntractor:  connActor,
	}
}

func (p *sessionNegController) GetSessionNegInteractor() interactor.SessionNegInteractor {
	return p.sessionNegInteractor
}

func (p *sessionNegController) Negotiate() {
	sessions := p.tcpSessionIntractor.GetSessionWaitNeg()
	if len(sessions) == 0 {
		return
	}
	for _, session := range sessions {
		p.tcpSessionIntractor.SetSessionDoing(&session)
		go func() {
			ok := p.sessionNegInteractor.Negotiate(&session)
			if !ok {
				p.tcpSessionIntractor.CloseByConnectId(session.GetId())
				return
			}
			p.sessionNegInteractor.NewSession(&session)
			p.tcpSessionIntractor.RemoveSession(&session)
		}()
	}
}
