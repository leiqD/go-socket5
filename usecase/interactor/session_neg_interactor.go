package interactor

import (
	"github.com/leiqD/go-socket5/domain/model"
	up "github.com/leiqD/go-socket5/usecase/presenter"
	ur "github.com/leiqD/go-socket5/usecase/repository"
)

type sessionNegInteractor struct {
	reponsitory ur.CtrlSessionRepository
	presenter   up.CtrlSessionPresenter
}

type SessionNegInteractor interface {
	Negotiate(session *model.CtrlSession) (succ bool)
	NewSession(session *model.CtrlSession)
	GetSessionWaitTrans() []model.CtrlSession
	Close(model.CtrlSession)
	CloseByConnectId(connectId model.ConnectId)
	RemoveSession(session *model.CtrlSession)
	SetSessionDoing(session *model.CtrlSession)
}

func NewCtrlSessionInteractor(r ur.CtrlSessionRepository, p up.CtrlSessionPresenter) SessionNegInteractor {
	return &sessionNegInteractor{
		reponsitory: r,
		presenter:   p,
	}
}

func (p *sessionNegInteractor) Negotiate(session *model.CtrlSession) (succ bool) {
	return p.presenter.Negotiate(session)
}

func (p *sessionNegInteractor) NewSession(session *model.CtrlSession) {
	p.reponsitory.NewSession(session)
}
func (p *sessionNegInteractor) GetSessionWaitTrans() []model.CtrlSession {
	return p.reponsitory.GetWithoutDoing()
}
func (p *sessionNegInteractor) Close(session model.CtrlSession) {
	p.reponsitory.CloseSession(session)
}

func (p *sessionNegInteractor) CloseByConnectId(connectId model.ConnectId) {
	p.reponsitory.CloseByConnectId(connectId)
}
func (p *sessionNegInteractor) RemoveSession(session *model.CtrlSession) {
	p.reponsitory.RemoveSession(session.GetId())
}
func (p *sessionNegInteractor) SetSessionDoing(session *model.CtrlSession) {
	p.reponsitory.SetDoing(session.GetId())
}
