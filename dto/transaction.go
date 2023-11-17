package dto

type Transaction struct {
	Description string  `json:"description" validate:"required"`
	Amount      float32 `json:"amount" validte:"required"`
}
