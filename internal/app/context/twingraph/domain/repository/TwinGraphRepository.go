package repository

import (
	"context"
	"fmt"

	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/context/twingraph/domain"
	cache "github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/infra/cache"
)

const (
	TWIN_GRAPH_KEY    = "twin_graph"
	TWIN_SUBGRAPH_KEY = "twin_graph.%s"
)

type TwinGraphRepository interface {
	GetTwinGraph() (domain.TwinGraph, error)
	SetTwinGraph(twinGraph domain.TwinGraph) error
	SetTwinSubGraph(twinGraph domain.TwinGraph, twinInterface string) error
	GetTwinSubGraph(twinInterface string) (domain.TwinGraph, error)
}

func NewTwinGraphRepository(
	mapper TwinGraphMapper,
	dbConnection cache.CacheConnection,
) TwinGraphRepository {
	return &twinGraphRepository{
		cacheConnection: dbConnection,
		mapper:          mapper,
	}
}

type twinGraphRepository struct {
	cacheConnection cache.CacheConnection
	mapper          TwinGraphMapper
}

func (t *twinGraphRepository) GetTwinGraph() (domain.TwinGraph, error) {
	var twinGraph domain.TwinGraph
	err := t.cacheConnection.Get(context.Background(), TWIN_GRAPH_KEY, &twinGraph)

	if err != nil {
		return domain.TwinGraph{}, err
	}

	return twinGraph, nil
}

func (t *twinGraphRepository) SetTwinGraph(twinGraph domain.TwinGraph) error {
	return t.cacheConnection.Set(context.Background(), TWIN_GRAPH_KEY, twinGraph)
}

func (t *twinGraphRepository) SetTwinSubGraph(twinGraph domain.TwinGraph, twinInterface string) error {
	return t.cacheConnection.Set(context.Background(), fmt.Sprintf(TWIN_SUBGRAPH_KEY, twinInterface), twinGraph)
}

func (t *twinGraphRepository) GetTwinSubGraph(twinInterface string) (domain.TwinGraph, error) {
	var twinGraph domain.TwinGraph
	err := t.cacheConnection.Get(context.Background(), fmt.Sprintf(TWIN_SUBGRAPH_KEY, twinInterface), &twinGraph)

	if err != nil {
		return domain.TwinGraph{}, err
	}

	return twinGraph, nil
}
