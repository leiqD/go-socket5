package app

import (
	"github.com/leiqD/go-socket5/app/proxy/launcher"
	"github.com/leiqD/go-socket5/infrastructure/conf"
	"github.com/leiqD/go-socket5/infrastructure/logger"
	"github.com/leiqD/go-socket5/infrastructure/router"
	"github.com/leiqD/go-socket5/interface/controller"
	"github.com/leiqD/go-socket5/trans"
	"gorm.io/gorm"
)

type Program struct {
	conf    *conf.Configs
	log     logger.LoggerInterface
	db      *gorm.DB
	tran    trans.Trans
	route   router.Router
	control controller.AppController
}

func (p *Program) Init() error {
	p.conf = launcher.InitializeConfig("")
	p.log = launcher.InitializeLog(p.conf)
	db, err := launcher.InitialDataStore(p.conf)
	if err != nil {
		return err
	}
	p.db = db
	p.tran = launcher.InitialTrans()
	p.control = p.tran.NewAppController()
	p.route = launcher.InitialRouter(p.conf, p.control)
	return nil
}

func (p *Program) Start() error {
	logger.Info("Service Start")
	p.route.Start()
	return nil
}

func (p *Program) Stop() error {
	return nil
}

func (p *Program) ReloadConfig() error {
	p.conf.ReloadViper()
	return nil
}

func (p *Program) OneLoop() error {
	p.control.Negotiation.Negotiate()
	p.control.Trans.Run()
	return nil
}
