package trans

import (
	"github.com/leiqD/go-socket5/interface/controller"
	"github.com/leiqD/go-socket5/usecase/interactor"
	up "github.com/leiqD/go-socket5/usecase/presenter"
	ur "github.com/leiqD/go-socket5/usecase/repository"
)

func (r *trans) NetTransferController(sessionNegActor interactor.SessionNegInteractor) controller.TransController {
	return controller.NewTransController(r.NewTransferInteractor(), sessionNegActor)
}

func (r *trans) NewTransferInteractor() interactor.TcpTransInteractor {
	return interactor.NewTcpTransInteractor(r.NewTcpTransRepository(), r.NewTransPresenter())
}

func (r *trans) NewTcpTransRepository() ur.TcpTransRepository {
	return ur.NewTcpTransRepository()
}

func (r *trans) NewTransPresenter() up.TcpTransPresenter {
	return up.NewtcpTransPresenter()
}
