package types

type Student struct {
	ID int 
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Age string `json:"age" validate:"required"`
}