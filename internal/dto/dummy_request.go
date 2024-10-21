package dto

type DummyRequest struct {
	Foo string `json:"foo" validate:"required"`
	Bar string `json:"bar" validate:"required"`
}
