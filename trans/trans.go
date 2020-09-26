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
	return controller.AppController{
		Negotiate: r.NewTcpConnController(),
	}
}
