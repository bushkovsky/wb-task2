package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных

	a и b, значение которых > 2^20.
*/

func sum(x, y *big.Int) *big.Int {
	result := &big.Int{}
	result.Add(x, y)
	return result
}

func sub(x, y *big.Int) *big.Int {
	result := &big.Int{}
	result.Sub(x, y)
	return result
}

func mult(x, y *big.Int) *big.Int {
	result := &big.Int{}
	result.Mul(x, y)
	return result
}

func div(x, y *big.Int) *big.Int {
	result := &big.Int{}
	result.Div(x, y)
	return result
}

func main() {
	x := big.NewInt(1 << 21)
	y := big.NewInt(1 << 22)

	fmt.Printf("sum = %d\n", sum(x, y))
	fmt.Printf("sub = %d\n", sub(x, y))
	fmt.Printf("mult = %d\n", mult(x, y))
	fmt.Printf("div = %d\n", div(x, y))

}
