package presenter

import (
	"github.com/leiqD/go-socket5/domain/model"
	"github.com/leiqD/go-socket5/infrastructure/logger"
	"sync"
)

type ctrlSessionPresenter struct {
}

type CtrlSessionPresenter interface {
	Negotiate(session []model.CtrlSession) (succ, fail []*model.CtrlSession)
}

func NewCtrlSessionPresenter() CtrlSessionPresenter {
	return &ctrlSessionPresenter{}
}

func (p *ctrlSessionPresenter) Negotiate(sessions []model.CtrlSession) (succ, fail []*model.CtrlSession) {
	logger.Infof("Negotiate has session %d to negotiate", len(sessions))
	wg := sync.WaitGroup{}
	for _, session := range sessions {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := p.negotiate(&session); err != nil {
				fail = append(fail, &session)
				return
			}
			if err := p.auth(&session); err != nil {
				fail = append(fail, &session)
				return
			}
			if err := p.request(&session); err != nil {
				fail = append(fail, &session)
				return
			}
			succ = append(succ, &session)
		}()
	}
	wg.Wait()
	return
}

func (p *ctrlSessionPresenter) negotiate(session *model.CtrlSession) error {
	buff := make([]byte, 1024)
	n, err := session.GetConn().Read(buff)
	logger.Infof("read from %s buff len=%d", session.GetRemoteAddrContent(), n)
	if err != nil {
		logger.Errorf("Negotiate read error=%v", err)
		return err
	}
	neg := model.NewNegotiation()
	if err := neg.Parse(buff); err != nil {
		return err
	}
	n, err = session.GetConn().Write(neg.Pack())
	if n != 2 || err != nil {
		return err
	}
	return nil
}

func (p *ctrlSessionPresenter) auth(session *model.CtrlSession) error {
	buff := make([]byte, 1024)
	n, err := session.GetConn().Read(buff)
	logger.Infof("read from %s buff len=%d", session.GetRemoteAddrContent(), n)
	if err != nil {
		logger.Errorf("request read error=%v", err)
		return err
	}
	auth := model.NewAuthentication()
	if err := auth.Parse(buff); err != nil {
		logger.Errorf("request read error=%v", err)
		return err
	}
	n, err = session.GetConn().Write(auth.Pack())
	if n != 2 || err != nil {
		return err
	}
	return nil
}

func (p *ctrlSessionPresenter) request(session *model.CtrlSession) error {
	buff := make([]byte, 1024)
	n, err := session.GetConn().Read(buff)
	logger.Infof("read from %s buff len=%d", session.GetRemoteAddrContent(), n)
	if err != nil {
		logger.Errorf("request read error=%v", err)
		return err
	}
	req := model.NewRequest()
	if err := req.Parse(buff); err != nil {
		logger.Errorf("request read error=%v", err)
		return err
	}
	conn, err := req.ConnectTcpDst()
	if conn == nil || err != nil {
		logger.Infof("connect remote dst addr fail %s", err.Error())
		return err
	}
	logger.Infof("connect remote dst addr success")
	if req.IsTcp() {
		session.SetProtocol(model.ProtocolTcp)
	}
	if req.IsUdp() {
		session.SetProtocol(model.ProtocolUdp)
	}
	session.SetDestConn(conn)
	session.SetSetup(model.SessionSetupWaitTrans)
	n, err = session.GetConn().Write(req.Pack())
	if n != 2 || err != nil {
		return err
	}
	return nil
}

func (p *ctrlSessionPresenter) connect(session *model.CtrlSession) error {

	return nil
}

func (p *ctrlSessionPresenter) bind(session *model.CtrlSession) error {
	return nil
}

func (p *ctrlSessionPresenter) udpAssociate(session *model.CtrlSession) error {
	return nil
}
