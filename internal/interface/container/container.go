package container

import (
	"go_service/internal/infrastructure/psql/database"
	"go_service/internal/shared/config"
	"go_service/internal/shared/constants"
)

type Container struct {
	Config         *config.Config
	ProductService product.Service
}

func NewContainer(conf *config.Config) *Container {
	var dbRepo = repositories.NewDBRepository(database.NewDBConnection(&conf.Database))
	if conf.Env == constants.EnvLocal {
		err := database.MigrateAndSeed(dbRepo.DB())
		if err != nil {
			panic(err)
		}
	}

	return &Container{
		Config:         conf,
		ProductService: product.NewService(dbRepo, dbRepo),
	}
}
