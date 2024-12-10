package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int).Exp(big.NewInt(2), big.NewInt(22), nil) // 2^21
	b := new(big.Int).Exp(big.NewInt(2), big.NewInt(21), nil) // 2^20

	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n", b.String())

	sum := new(big.Int).Add(a, b)
	fmt.Printf("Сумма (a + b) = %s\n", sum.String())

	diff := new(big.Int).Sub(a, b)
	fmt.Printf("Разность (a - b) = %s\n", diff.String())

	product := new(big.Int).Mul(a, b)
	fmt.Printf("Произведение (a * b) = %s\n", product.String())

	quotient := new(big.Int).Div(a, b)
	fmt.Printf("Частное (a / b) = %s\n", quotient.String())
}
