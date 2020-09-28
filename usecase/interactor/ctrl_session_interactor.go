package interactor

import (
	"github.com/leiqD/go-socket5/domain/model"
	up "github.com/leiqD/go-socket5/usecase/presenter"
	ur "github.com/leiqD/go-socket5/usecase/repository"
)

type ctrlSessionInteractor struct {
	reponsitory ur.CtrlSessionRepository
	presenter   up.CtrlSessionPresenter
}

type CtrlSessionInteractor interface {
	Negotiate(session []model.CtrlSession) (succSession []model.ConnectId, failSession []model.ConnectId)
}

func NewCtrlSessionInteractor(r ur.CtrlSessionRepository, p up.CtrlSessionPresenter) CtrlSessionInteractor {
	return &ctrlSessionInteractor{
		reponsitory: r,
		presenter:   p,
	}
}

func (p *ctrlSessionInteractor) Negotiate(session []model.CtrlSession) (succSession []model.ConnectId, failSession []model.ConnectId) {
	return p.presenter.Negotiate(session)
}
