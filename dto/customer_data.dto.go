package dto

type CustomerData struct {
	ID   string `Json:"id"`
	Code string `Json:"code"`
	Name string `Json:"name"`
}

type CreateCustomerRequest struct {
	Code string `Json:"code" validate:"required"`
	Name string `Json:"name" validate:"required"`
}

type UpdateCustomerRequest struct {
	ID   string `Json:"-"`
	Name string `Json:"name" validate:"required"`
	Code string `Json:"code" validate:"required"`
}
