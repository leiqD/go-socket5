package trans

import (
	"github.com/leiqD/go-socket5/interface/controller"
	ip "github.com/leiqD/go-socket5/interface/presenter"
	"github.com/leiqD/go-socket5/usecase/interactor"
	up "github.com/leiqD/go-socket5/usecase/presenter"
	ur "github.com/leiqD/go-socket5/usecase/repository"
)

func (r *trans) NewTcpConnController() controller.TcpConnController {
	return controller.NewTcpConnController(r.NewTcpConnInteractor())
}

func (r *trans) NewTcpConnInteractor() interactor.TcpConnInterfactor {
	return interactor.NewTcpConnInterfactor(r.NewTcpConnRepository(), r.NewTcpConnPresenter())
}

func (r *trans) NewTcpConnRepository() ur.TcpConnRepository {
	return ur.NewTcpConnRepository()
}

func (r *trans) NewTcpConnPresenter() up.TcpConnPresenter {
	return ip.NewTcpConnPresenter()
}
