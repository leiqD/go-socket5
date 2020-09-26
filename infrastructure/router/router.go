package router

import (
	"errors"
	"github.com/leiqD/go-socket5/infrastructure/conf"
	"github.com/leiqD/go-socket5/infrastructure/logger"
	"github.com/leiqD/go-socket5/interface/controller"
	"math/rand"
	"net"
	"time"
)

type RouterConfig interface {
	Net() *conf.NetInfo
}

type router struct {
	netInfo *conf.NetInfo
	contron controller.AppController
	listen  net.Listener
}

const (
	Ipv4Tcp = "tcp4"
	Ipv4Udp = "udp4"
)

type Router interface {
	Start() error
	Run()
}

func NewRouter(cfg RouterConfig, c controller.AppController) Router {
	return &router{
		netInfo: cfg.Net(),
		contron: c,
	}
}

func (p *router) Start() error {
	switch p.netInfo.Protocol {
	default:
		return errors.New("unknow error")
	case "tcp":
		listen, err := net.Listen(Ipv4Tcp, p.netInfo.Addr)
		rand.Seed(time.Now().Unix())
		if err != nil {
			logger.Errorf("start tcp server error %s", err.Error())
			return err
		}
		logger.Info("start tcp server success ")
		p.listen = listen
	}
	return nil
}

func (p *router) Run() {
	p.accept()
}

func (p *router) accept() {
	conn, err := p.listen.Accept()
	if err != nil {
		return
	}
	p.newSession(conn)
}

func (p *router) newSession(conn net.Conn) {
	p.contron.Negotiate.NewSession(conn)
}
