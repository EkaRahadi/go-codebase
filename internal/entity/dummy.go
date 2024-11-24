package entity

import "github.com/EkaRahadi/go-codebase/internal/dto"

type Dummy struct {
	Message string
}

func (d Dummy) GenerateCourierDTO() *dto.DummyResponse {
	return &dto.DummyResponse{
		Message: d.Message,
	}
}
