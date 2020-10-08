package trans

import (
	"github.com/leiqD/go-socket5/interface/controller"
)

type trans struct {
}

type Trans interface {
	NewAppController() controller.AppController
}

func NetTrans() Trans {
	return &trans{}
}

func (r *trans) NewAppController() controller.AppController {
	app := controller.AppController{}
	app.SetConn(r.NewTcpConnController())
	app.SetNeg(r.NewCtrlSessionController(app.GetConn().GetConnActor()))
	app.SetTran(r.NetTransferController(app.GetNeg().GetSessionNegInteractor()))
	return app
}
