package presenter

import (
	"fmt"
	"github.com/leiqD/go-socket5/domain/model"
	"github.com/leiqD/go-socket5/infrastructure/logger"
)

type ctrlSessionPresenter struct {
}

type CtrlSessionPresenter interface {
	Negotiate(session *model.CtrlSession) (succ bool)
}

func NewCtrlSessionPresenter() CtrlSessionPresenter {
	return &ctrlSessionPresenter{}
}

func (p *ctrlSessionPresenter) Negotiate(session *model.CtrlSession) (succ bool) {
	//logger.Infof("Negotiate has session %d to negotiate", len(sessions))
	if err := p.negotiate(session); err != nil {
		return false
	}
	if err := p.auth(session); err != nil {
		return false
	}
	if err := p.request(session); err != nil {
		return false
	}
	return true
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
	cmd := req.Cmd()
	switch cmd {
	default:
		return fmt.Errorf("undefine cmd")
	case model.CONNECT:
		logger.Infof("cmd=%d:connect", cmd)
		return p.connect(session, req)
	case model.BIND:
		logger.Infof("cmd=%d:bind", cmd)
		return p.bind(session, req)
	case model.UDP_ASSOCIATE:
		logger.Infof("cmd=%d:udp associate", cmd)
		return p.udpAssociate(session, req)
	}
	return nil
}

func (p *ctrlSessionPresenter) connect(session *model.CtrlSession, req model.Request) error {
	conn, err := req.ConnectTcpDst()
	if conn == nil || err != nil {
		logger.Infof("connect remote dst addr fail %s", err.Error())
		return err
	}
	logger.Infof("connect remote dst addr success")
	session.SetProtocol(model.ProtocolTcp)
	session.SetDestConn(conn)
	n, err := session.GetConn().Write(req.Pack())
	if n != 2 || err != nil {
		return err
	}
	return nil
}

func (p *ctrlSessionPresenter) bind(session *model.CtrlSession, req model.Request) error {
	return fmt.Errorf("not support bind")
}

func (p *ctrlSessionPresenter) udpAssociate(session *model.CtrlSession, req model.Request) error {
	session.SetProtocol(model.ProtocolUdp)
	return nil
}
