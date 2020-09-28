package presenter

import (
	"github.com/leiqD/go-socket5/domain/model"
	"github.com/leiqD/go-socket5/infrastructure/logger"
	"sync"
)

type ctrlSessionPresenter struct {
}

type CtrlSessionPresenter interface {
	Negotiate(session []model.CtrlSession) (succSession []model.ConnectId, failSession []model.ConnectId)
}

func NewCtrlSessionPresenter() CtrlSessionPresenter {
	return &ctrlSessionPresenter{}
}

func (p *ctrlSessionPresenter) Negotiate(sessions []model.CtrlSession) (succSession []model.ConnectId, failSession []model.ConnectId) {
	logger.Infof("Negotiate has session %d to negotiate", len(sessions))
	wg := sync.WaitGroup{}
	for _, session := range sessions {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := p.negotiate(session); err != nil {
				failSession = append(failSession, session.GetId())
				return
			}
			succSession = append(succSession, session.GetId())
		}()
	}
	wg.Wait()
	return
}

func (p *ctrlSessionPresenter) negotiate(session model.CtrlSession) error {
	buff := make([]byte, 1024)
	_, err := session.GetConn().Read(buff)
	if err != nil {
		logger.Errorf("Negotiate read error=%v", err)
		return err
	}
	return nil
}
