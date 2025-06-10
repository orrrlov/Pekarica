package main

type recipe struct {
	id          int
	name        string
	amount      amount
	ingredients []ingredient
}

type ingredient struct {
	id     int
	name   string
	amount amount
}

type amount struct {
	quantity float64
	unit     unit
}

type unit int

const (
	GRAM unit = iota
	LITER
	WHOLE
)

func (u unit) String() string {
	switch u {
	case GRAM:
		return "gram"
	case LITER:
		return "liter"
	default:
		return "whole"
	}
}

func (u unit) ConvertToBase() {

}
