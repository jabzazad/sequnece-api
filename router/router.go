package router

import (
	"sequence-api/core/config"
	"sequence-api/core/errors"
	"sequence-api/core/validator"
	"sequence-api/internal/pkg/healthcheck"
	"sequence-api/internal/pkg/sequence"

	customMiddleware "sequence-api/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	swagger "github.com/swaggo/echo-swagger"
)

// Options option for new router
type Options struct {
	LogMiddleware echo.MiddlewareFunc
	Environment   *config.Environment
	Results       *config.Results
}

// NewWithOptions new router with options
func NewWithOptions(options *Options) *echo.Echo {
	if options.Environment == nil {
		panic("Not found Environment")
	}

	router := echo.New()
	router.Validator = validator.New()
	router.HTTPErrorHandler = errors.HTTPErrorHandler

	router.Logger.SetPrefix("Qchang")
	router.Use(middleware.Recover())
	router.Use(customMiddleware.CustomContext())

	api := router.Group("api/:version")
	if options != nil {
		api.Use(options.LogMiddleware)
	}

	api.Use(
		middleware.Secure(),
		middleware.Gzip(),
		customMiddleware.WrapResponse(options.Results),
		customMiddleware.AcceptLanguage(),
	)

	if config.ENV.Swagger.Enable {
		api.GET("/swagger/*", swagger.WrapHandler)
	}

	healthCheckEndpoint := healthcheck.NewEndpoint()
	api.GET("/healthz", healthCheckEndpoint.HealthCheck)

	sequenceEndpoint := sequence.NewEndpoint()
	sequenceGroup := api.Group("/sequences")
	{
		sequenceGroup.GET("", sequenceEndpoint.GetResultSequence)
	}

	return router
}
