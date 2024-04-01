//go:build wireinject
// +build wireinject

package TwinGraph

import (
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/context/twingraph/controller"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/context/twingraph/domain/repository"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/context/twingraph/usecase"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/infra/cache"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/infra/validator"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/pkg/log"

	"github.com/google/wire"
)

func NewTwinGraphContainer(
	controller controller.TwinGraphController,
	mapper repository.TwinGraphMapper,
	repository repository.TwinGraphRepository,
	usecase usecase.TwinGraphUseCase,
) TwinGraphContainer {
	return TwinGraphContainer{
		controller: controller,
		repository: repository,
		mapper:     mapper,
		usecase:    usecase,
	}
}

type TwinGraphContainer struct {
	controller controller.TwinGraphController
	repository repository.TwinGraphRepository
	mapper     repository.TwinGraphMapper
	usecase    usecase.TwinGraphUseCase
}

func InitializeTwinGraphContainer(cacheConnection cache.CacheConnection) TwinGraphContainer {
	wire.Build(
		NewTwinGraphContainer,
		controller.NewTwinGraphController,
		usecase.NewTwinGraphUseCase,
		repository.NewTwinGraphRepository,
		repository.NewTwinGraphMapper,

		validator.NewValidator,
		log.NewLogger,
	)

	return TwinGraphContainer{}
}
