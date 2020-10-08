package presenter

import (
	"fmt"
	"github.com/leiqD/go-socket5/domain/model"
	"github.com/leiqD/go-socket5/infrastructure/logger"
	"sync"
)

type tcpTransPresenter struct {
}

type TcpTransPresenter interface {
	Handle(session *model.CtrlSession) error
	client2remote(session *model.CtrlSession) error
	remote2client(session *model.CtrlSession) error
}

func NewtcpTransPresenter() TcpTransPresenter {
	return &tcpTransPresenter{}
}

func (p *tcpTransPresenter) Handle(session *model.CtrlSession) error {

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() error {
		for {
			if err := p.client2remote(session); err != nil {
				logger.Errorf("err=%s", err.Error())
				wg.Done()
				return err
			}
		}
	}()
	wg.Add(1)
	go func() {
		for {
			if err := p.remote2client(session); err != nil {
				logger.Errorf("err=%s", err.Error())
				wg.Done()
				return
			}
		}
	}()
	wg.Wait()
	return fmt.Errorf("session commplete")
}

func (p *tcpTransPresenter) client2remote(session *model.CtrlSession) error {
	recvBuff := make([]byte, 1024*1024)
	client := session.GetConn()
	n, err := client.Read(recvBuff)
	//logger.Infof("client2remote n=%d", n)
	if err != nil {
		return err
	}
	remote := session.GetDestConn()
	sendBuff := recvBuff[0:n]
	wn, err := remote.Write(sendBuff)
	if err != nil {
		return err
	}
	if n != wn {
		//return fmt.Errorf("write buff is full n=%d wn=%d", n, wn)
	}
	logger.Debugf("client2remote buff n=%d,wn=%d client=%s remote=%s", n, wn, client.RemoteAddr().String(), remote.RemoteAddr().String())
	return nil
}

func (p *tcpTransPresenter) remote2client(session *model.CtrlSession) error {
	recvBuff := make([]byte, 1024*1024)
	remote := session.GetDestConn()
	n, err := remote.Read(recvBuff)
	//logger.Infof("remote2client n=%d", n)
	if err != nil {
		return err
	}
	client := session.GetConn()
	sendBuff := recvBuff[0:n]
	wn, err := client.Write(sendBuff)
	if err != nil {
		return err
	}
	if n != wn {
		//return fmt.Errorf("write buff is full n=%d wn=%d", n, wn)
	}
	logger.Debugf("remote2client buff n=%d,wn=%d remote=%s client=%s", n, wn, remote.RemoteAddr().String(), client.RemoteAddr().String())

	return nil
}
