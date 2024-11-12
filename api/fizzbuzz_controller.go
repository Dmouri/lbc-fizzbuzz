package api

import (
	"lbc/fizzbuzz/domain"
	"lbc/fizzbuzz/repository"
	"lbc/fizzbuzz/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mwm-io/gapi/errors"
	"go.uber.org/zap"
)

type fizzBuzzController struct {
	fizzBuzzService    service.FizzBuzzService
	fizzBuzzRepository repository.FizzBuzzRepository
	logger             *zap.Logger
}

type FizzBuzzResponse struct {
	Result string `json:"result"`
}

func SetupFizzBuzzController(
	logger *zap.Logger,
	router gin.IRouter,
	fizzBuzzService service.FizzBuzzService,
	fizzBuzzRepository repository.FizzBuzzRepository) {
	c := fizzBuzzController{
		logger:             logger,
		fizzBuzzService:    fizzBuzzService,
		fizzBuzzRepository: fizzBuzzRepository,
	}

	root := router.Group("/api/v1/fizzbuzz")
	GET(root, "/", c.generateFizzBuzzEndpoint)
	GET(root, "/stats", c.getFizzBuzzStatsEndpoint)
}

// generateFizzBuzzEndpoint handles the FizzBuzz generation request
func (c *fizzBuzzController) generateFizzBuzzEndpoint(ctx *gin.Context) {
	fbInput, err := GetQueryParams(ctx)
	if err != nil {
		c.logger.Error("Failed to parse query parameters", zap.Error(err))
		ctx.JSON(err.StatusCode(), gin.H{"error": err})
		return
	}

	result, err := c.fizzBuzzService.GenerateFizzBuzz(fbInput)
	if err != nil {
		c.logger.Error("Failed to generate FizzBuzz", zap.Error(err))
		ctx.JSON(err.StatusCode(), gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, FizzBuzzResponse{Result: result})
}

// GetQueryParams retrieves and parses query parameters with validation and defaults
func GetQueryParams(ctx *gin.Context) (domain.FizzBuzzInput, errors.Error) {
	int1, err := strconv.Atoi(ctx.Query("int1"))
	if err != nil {
		return domain.FizzBuzzInput{}, errors.BadRequest("failed_to_parse_int1", "failed to parse int1")
	}

	int2, err := strconv.Atoi(ctx.Query("int2"))
	if err != nil {
		return domain.FizzBuzzInput{}, errors.BadRequest("failed_to_parse_int2", "failed to parse int2")
	}

	limitStr := ctx.Query("limit")
	// Set 100 has default limit if not provided
	limit := 100
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return domain.FizzBuzzInput{}, errors.BadRequest("failed_to_parse_limit", "failed to parse limit")
		}
	}

	str1 := ctx.Query("str1")
	str2 := ctx.Query("str2")

	return domain.FizzBuzzInput{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  str1,
		Str2:  str2,
	}, nil
}

// getFizzBuzzStatsEndpoint handles the FizzBuzz stats request
func (c *fizzBuzzController) getFizzBuzzStatsEndpoint(ctx *gin.Context) {
	fbRequest, err := c.fizzBuzzRepository.GetMostHits(ctx)
	if err != nil {
		c.logger.Error("Failed to get most hits FizzBuzzRequest", zap.Error(err))
		ctx.JSON(err.StatusCode(), gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, fbRequest)
}
