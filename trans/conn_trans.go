package trans

import (
	"github.com/leiqD/go-socket5/interface/controller"
	ip "github.com/leiqD/go-socket5/interface/presenter"
	ir "github.com/leiqD/go-socket5/interface/repository"
	"github.com/leiqD/go-socket5/usecase/interactor"
	up "github.com/leiqD/go-socket5/usecase/presenter"
	ur "github.com/leiqD/go-socket5/usecase/repository"
)

func (r *trans) NewTcpConnController() controller.TcpConnController {
	return controller.NewTcpConnController(r.NewTcpConnInteractor())
}

func (r *trans) NewTcpConnInteractor() interactor.TcpConnInterfactor {
	return interactor.NewTcpConnInterfactor(r.NewTcpConnPresenter(), r.NewTcpConnRepository())
}

func (r *trans) NewTcpConnRepository() ur.TcpConnReponsitory {
	return ir.NewTcpConnRepository()
}

func (r *trans) NewTcpConnPresenter() up.TcpConnPresenter {
	return ip.NewTcpConnPresenter()
}
