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
	Handle(session *model.CtrlSession) error
	NewSession(session *model.CtrlSession)
	Close(session *model.CtrlSession)
	CloseByConnectId(connectId model.ConnectId)
}

func NewTcpTransInteractor(r ur.TcpTransRepository, p up.TcpTransPresenter) TcpTransInteractor {
	return &tcpTransInteractor{
		reponsitory: r,
		presenter:   p,
	}
}

func (p *tcpTransInteractor) Handle(session *model.CtrlSession) error {
	if err := p.presenter.Handle(session); err != nil {
		logger.Errorf("err=%s", err.Error())
		return err
	}
	return nil
}

func (p *tcpTransInteractor) NewSession(session *model.CtrlSession) {
	p.reponsitory.NewSession(session)
}
func (p *tcpTransInteractor) Close(session *model.CtrlSession) {
	p.reponsitory.CloseByConnectId(session.GetId())
}
func (p *tcpTransInteractor) CloseByConnectId(connectId model.ConnectId) {
	p.reponsitory.CloseByConnectId(connectId)
}
