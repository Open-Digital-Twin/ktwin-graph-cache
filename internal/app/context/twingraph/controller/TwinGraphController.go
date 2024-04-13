package controller

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/context/twingraph/domain"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/context/twingraph/domain/repository"
	dtdv0 "github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/context/twingraph/domain/repository/dtd"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/context/twingraph/usecase"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/app/infra/validator"
	"github.com/Open-Digital-Twin/ktwin-graph-store/internal/pkg/log"
	"k8s.io/client-go/rest"

	"github.com/gin-gonic/gin"
)

type TwinGraphController interface {
	GetTwinGraph(g *gin.Context)
	GetTwinGraphByTwinInstance(g *gin.Context)
	UpdateTwinGraph(result rest.Result)
}

func NewTwinGraphController(
	twinGraphUseCase usecase.TwinGraphUseCase,
	twinGraphMapper repository.TwinGraphMapper,
	validator validator.Validator,
	logger log.Logger,
) TwinGraphController {
	return &twinGraphController{
		twinGraphUseCase: twinGraphUseCase,
		twinGraphMapper:  twinGraphMapper,
		validator:        validator,
		logger:           logger,
	}
}

type twinGraphController struct {
	twinGraphUseCase usecase.TwinGraphUseCase
	twinGraphMapper  repository.TwinGraphMapper
	validator        validator.Validator
	logger           log.Logger
}

// Get Twin Graph godoc
// @Summary Get Twin Graph
// @Schemes
// @Description do ping
// @Tags TwinGraphs
// @Accept json
// @Produce json
// @Success 200 {object} []domain.TwinGraph
// @Router /twin-graph [get]
func (t *twinGraphController) GetTwinGraph(g *gin.Context) {
	twinGraph, err := t.twinGraphUseCase.GetTwinGraph()

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
		t.logger.Error("Error: " + err.Error())
		return
	} else {
		g.JSON(http.StatusOK, twinGraph)
	}
}

func (t *twinGraphController) GetTwinGraphByTwinInstance(g *gin.Context) {
	interfaceId, _ := g.Params.Get("interfaceId")

	if interfaceId == "" {
		g.JSON(http.StatusBadRequest, "Missing parameters")
		return
	}

	twinGraph, err := t.twinGraphUseCase.GetTwinGraphByInterface(interfaceId)

	if err != nil {
		g.JSON(http.StatusInternalServerError, "Error: "+err.Error())
		t.logger.Error("Error: " + err.Error())
		return
	} else if reflect.DeepEqual(twinGraph, domain.TwinGraph{}) {
		g.JSON(http.StatusNotFound, "Not Found")
	} else {
		g.JSON(http.StatusOK, twinGraph)
	}
}

// Update Twin Graph godoc
// @Summary Update Twin Graph
// @Schemes
// @Description This endpoint updates the latest TwinGraph.
// @Tags TwinGraphs
// @Accept json
// @Produce json
// @Success 200 {object} domain.TwinGraph
// @Router /twin-graph [post]
func (t *twinGraphController) UpdateTwinGraph(result rest.Result) {

	var twinInstanceList dtdv0.TwinInstanceList
	rawData, _ := result.Raw()
	err := json.Unmarshal(rawData, &twinInstanceList)

	if err != nil {
		t.logger.Error("Error while unmarshalling twin instances: " + err.Error())
		return
	}

	twinGraph := t.twinGraphMapper.TwinInstanceToTwinGraph(twinInstanceList.Items)
	err = t.twinGraphUseCase.UpdateTwinGraph(twinGraph)

	if err != nil {
		t.logger.Error("Error while updating twin graph: " + err.Error())
	}
}
