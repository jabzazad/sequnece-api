package sequence

import (
	"sequence-api/core/config"
	"sequence-api/core/handlers"

	"github.com/labstack/echo/v4"
)

// Endpoint endpoint interface
type Endpoint interface {
	GetResultSequence(c echo.Context) error
}

type endpoint struct {
	env     *config.Environment
	rr      *config.Results
	service Service
}

// NewEndpoint new endpoint
func NewEndpoint() Endpoint {
	return &endpoint{
		env:     config.ENV,
		rr:      config.RR,
		service: NewService(),
	}
}

// GetResultSequence
// @Tags Sequence
// @Summary GetResultSequence
// @Description Find x y z from [1, X, 8, 17, Y, Z, 78, 113]
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Success 200 {object} response.ResponseFirstQuestion
// @Failure 400 {object} response.Message
// @Failure 401 {object} response.Message
// @Failure 404 {object} response.Message
// @Failure 410 {object} response.Message
// @Router /sequences [get]
func (ep *endpoint) GetResultSequence(c echo.Context) error {
	return handlers.ResponseObjectWithoutRequest(c, ep.service.GetResultSequence)
}
