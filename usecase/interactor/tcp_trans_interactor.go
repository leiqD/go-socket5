package interactor

import (
	"github.com/leiqD/go-socket5/domain/model"
	"github.com/leiqD/go-socket5/infrastructure/logger"
	up "github.com/leiqD/go-socket5/usecase/presenter"
	ur "github.com/leiqD/go-socket5/usecase/repository"
)

type tcpTransInteractor struct {
	reponsitory ur.TcpTransRepository
	presenter   up.TcpTransPresenter
}

type TcpTransInteractor interface {
	Handle([]model.CtrlSession) error
}

func NewTcpTransInteractor(r ur.TcpTransRepository, p up.TcpTransPresenter) TcpTransInteractor {
	return &tcpTransInteractor{
		reponsitory: r,
		presenter:   p,
	}
}

func (p *tcpTransInteractor) Handle(sessions []model.CtrlSession) error {
	for _, session := range sessions {
		p.reponsitory.NewSession(&session)
	}
	for _, session := range sessions {
		go func() {
			logger.Infof("%s start tansfer", session.GetRemoteAddrContent())
			for {
				if err := p.presenter.Handle(&session); err != nil {
					logger.Errorf("err=%s", err.Error())
					p.reponsitory.CloseByConnectId(session.GetId())
					return
				}
			}
		}()
	}
	return nil
}
