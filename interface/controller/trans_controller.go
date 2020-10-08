package controller

import (
	"github.com/leiqD/go-socket5/usecase/interactor"
)

type transController struct {
	tcpTransIneracotr   interactor.TcpTransInteractor
	sessionNegIntractor interactor.SessionNegInteractor
}

type TransController interface {
	Trans()
}

func NewTransController(us interactor.TcpTransInteractor, sessionNegIntractor interactor.SessionNegInteractor) TransController {
	return &transController{
		tcpTransIneracotr:   us,
		sessionNegIntractor: sessionNegIntractor,
	}
}

func (p *transController) Trans() {
	waitTrans := p.sessionNegIntractor.GetSessionWaitTrans()
	if len(waitTrans) == 0 {
		return
	}
	for _, session := range waitTrans {
		p.sessionNegIntractor.RemoveSession(&session)
		p.tcpTransIneracotr.NewSession(&session)
		go func() {
			err := p.tcpTransIneracotr.Handle(&session)
			if err != nil {
				p.tcpTransIneracotr.CloseByConnectId(session.GetId())
				return
			}
		}()
	}
}
