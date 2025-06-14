package main

type Recipe struct {
	ID          int
	Name        string
	Description string
	Amount
	Ingredients []IngredientAmount
}

type Ingredient struct {
	ID          int
	Name        string
	Description string
}

type Amount struct {
	Quantity float64
	Unit     Unit
}

type IngredientAmount struct {
	Amount
	Ingredient
}

type Unit int

const (
	GRAM Unit = iota
	LITER
	WHOLE
)

func (u Unit) String() string {
	switch u {
	case GRAM:
		return "gram"
	case LITER:
		return "liter"
	default:
		return "whole"
	}
}

func (u Unit) ConvertToBase() {

}
