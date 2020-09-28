package presenter

import "github.com/leiqD/go-socket5/domain/model"

type ctrlSessionPresenter struct {
}

type CtrlSessionPresenter interface {
	Handle(session model.CtrlSession) error
}

func NewCtrlSessionPresenter() CtrlSessionPresenter {
	return &ctrlSessionPresenter{}
}

func (p *ctrlSessionPresenter) Handle(session model.CtrlSession) error {
	return nil
}
