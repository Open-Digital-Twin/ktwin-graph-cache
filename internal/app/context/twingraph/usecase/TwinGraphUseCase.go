package usecase

import (
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/context/twingraph/domain"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/context/twingraph/domain/repository"
)

func NewTwinGraphUseCase(
	repository repository.TwinGraphRepository,
) TwinGraphUseCase {
	return &twinGraphUseCase{
		repository: repository,
	}
}

type TwinGraphUseCase interface {
	GetTwinGraph() (domain.TwinGraph, error)
	GetTwinGraphByInterface(twinInterface string) (domain.TwinGraph, error)
	UpdateTwinGraph(twinInstanceGraph domain.TwinInstanceGraph) error
}

type twinGraphUseCase struct {
	repository repository.TwinGraphRepository
}

func (t *twinGraphUseCase) GetTwinGraphByInterface(twinInterface string) (domain.TwinGraph, error) {
	return t.repository.GetTwinSubGraph(twinInterface)
}

func (t *twinGraphUseCase) GetTwinGraph() (domain.TwinGraph, error) {
	return t.repository.GetTwinGraph()
}

func (t *twinGraphUseCase) UpdateTwinGraph(twinInstanceGraph domain.TwinInstanceGraph) error {
	err := t.repository.SetTwinGraph(twinInstanceGraph.GetTwinGraph())

	for _, twinInterface := range twinInstanceGraph.GetTwinInterfaces() {
		err = t.repository.SetTwinSubGraph(twinInstanceGraph.GetTwinGraphByTwinInterfaces(twinInterface), twinInterface)
		if err != nil {
			return err
		}
	}

	return err
}
