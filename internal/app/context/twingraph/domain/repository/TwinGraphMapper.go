package repository

import (
	"fmt"

	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/context/twingraph/domain"
	dtdv0 "github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/context/twingraph/domain/repository/dtd"
)

type TwinGraphMapper interface {
	ToDomain(TwinGraph domain.TwinGraph) domain.TwinGraph
	TwinInstanceToTwinGraph(twinInstances []dtdv0.TwinInstance) domain.TwinInstanceGraph
}

type twinGraphMapper struct {
}

func NewTwinGraphMapper() TwinGraphMapper {
	return &twinGraphMapper{}
}

func (*twinGraphMapper) ToDomainList(TwinGraphs []domain.TwinGraph) []domain.TwinGraph {
	var TwinGraphsDomain []domain.TwinGraph

	for _, TwinGraph := range TwinGraphs {
		TwinGraphsDomain = append(TwinGraphsDomain, domain.TwinGraph(TwinGraph))
	}

	return TwinGraphsDomain
}

func (*twinGraphMapper) ToDomain(TwinGraph domain.TwinGraph) domain.TwinGraph {
	return domain.TwinGraph(TwinGraph)
}

func (t *twinGraphMapper) TwinInstanceToTwinGraph(twinInstances []dtdv0.TwinInstance) domain.TwinInstanceGraph {
	twinGraphInstance := domain.NewEmptyTwinInstanceGraph()

	for _, twinInstance := range twinInstances {
		twinGraphInstance.AddVertex(twinInstance)
	}

	for _, twinInstance := range twinInstances {
		for _, relationship := range twinInstance.Spec.TwinInstanceRelationships {
			twinInstanceVertex := twinGraphInstance.GetVertex(relationship.Instance)
			if twinInstanceVertex != nil {
				err := twinGraphInstance.AddEdge(twinInstance, *twinInstanceVertex)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}

	return twinGraphInstance
}
