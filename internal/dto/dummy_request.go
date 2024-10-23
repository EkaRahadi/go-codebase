package dto

type DummyRequest struct {
	Foo string `json:"foo" validate:"required"`
	Bar string `json:"bar" validate:"required"`
}

/*
all query params set to string.
use validator to check if string is convertable to number (valid number)
*/
type DummyRequestQuery struct {
	Foo string `form:"foo" validate:"required"`
	Bar string `form:"bar" validate:"required"`
}

/*
all uri is set to string.
example use validator to check if string is convertable to number (valid number)
*/
type DummyRequestUri struct {
	ExampleId string `uri:"example_id" validate:"required,number"`
}
