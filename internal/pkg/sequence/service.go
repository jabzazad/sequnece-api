package sequence

import (
	"sequence-api/core/config"
	"sequence-api/core/context"
	"sequence-api/internal/response.go"
)

// Service service interface
type Service interface {
	GetResultSequence(c *context.Context) (*response.ResponseFirstQuestion, error)
}

type service struct {
	env *config.Environment
	rr  *config.Results
}

// NewService new service
func NewService() Service {
	return &service{
		env: config.ENV,
		rr:  config.RR,
	}
}

// GetResultSequence get x y z from question 1
func (s *service) GetResultSequence(c *context.Context) (*response.ResponseFirstQuestion, error) {
	seq := []int{1, 0, 8, 17, 0, 0, 78, 113}
	maxLenght := len(seq)
	for i := len(seq) - 1; i > 0; i-- {
		if i <= len(seq)-3 {
			if seq[i] == 0 {
				seq[i] = seq[i+1] - ((seq[i+2] - seq[i+1]) - maxLenght)
			}

			maxLenght--
		}
	}

	return &response.ResponseFirstQuestion{
		X: seq[1],
		Y: seq[4],
		Z: seq[5],
	}, nil
}
