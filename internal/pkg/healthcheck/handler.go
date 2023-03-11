package healthcheck

import (
	"sequence-api/core/config"
	"sequence-api/core/handlers"

	"github.com/labstack/echo/v4"
)

// Endpoint endpoint interface
type Endpoint interface {
	HealthCheck(c echo.Context) error
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

// HealthCheck
// @Tags HealthCheck
// @Summary HealthCheck
// @Description HealthCheck server
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Success 200 {object} response.Message
// @Failure 400 {object} response.Message
// @Failure 401 {object} response.Message
// @Failure 404 {object} response.Message
// @Failure 410 {object} response.Message
// @Router /healthz [get]
func (ep *endpoint) HealthCheck(c echo.Context) error {
	return handlers.ResponseSuccessWithoutRequest(c, ep.service.HealthCheck)
}
