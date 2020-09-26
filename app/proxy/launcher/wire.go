//+build wireinject

package launcher

import (
	"github.com/google/wire"
	"github.com/leiqD/go-socket5/infrastructure/conf"
	"github.com/leiqD/go-socket5/infrastructure/datastore"
	"github.com/leiqD/go-socket5/infrastructure/logger"
	"github.com/leiqD/go-socket5/infrastructure/router"
	"github.com/leiqD/go-socket5/interface/controller"
	"github.com/leiqD/go-socket5/trans"
	"gorm.io/gorm"
)

func InitializeConfig(cfgPath string) *conf.Configs {
	wire.Build(conf.NewConfig, conf.NewViper)
	return nil
}

func InitializeLog(cfg logger.LoggerConfig) *logger.Zap {
	wire.Build(logger.NewLogger)
	return nil
}

func InitialDataStore(cfg datastore.MYSQLconfig) (*gorm.DB, error) {
	wire.Build(datastore.NewMYSQLDB)
	return nil, nil
}

func InitialRouter(cfg router.RouterConfig, control controller.AppController) router.Router {
	wire.Build(router.NewRouter)
	return nil
}

func InitialTrans() trans.Trans {
	wire.Build(trans.NetTrans)
	return nil
}
