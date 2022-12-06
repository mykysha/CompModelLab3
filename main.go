package main

import (
	"fmt"
	"math"
)

func p(x float64) float64 {
	return 4.0 * x / (x*x + 1.0)
}

func q(x float64) float64 {
	return -1.0 / (x*x + 1.0)
}

func f(x, h float64) float64 {
	return 3.0 / math.Pow(x*x+1.0, 2) / 2.0 * h * h
}

func a(x, h float64) float64 {
	return 1.0 / 2.0 * (1.0 + h/2.0*p(x))
}

func b(x, h float64) float64 {
	return 1.0 / 2.0 * (1.0 - h/2.0*p(x))
}

func c(x, h float64) float64 {
	return 1.0 + h*h*q(x)/2.0
}

func main() {
	from := 0.0
	to := 1.0
	h := 0.1
	N := int((to - from) / h)

	X := make([]float64, N+1)

	for i := 1; i < N+1; i++ {
		X[i] = float64(i) * h
	}

	Y := make([]float64, N+1)

	Y[0] = 0
	Y[N] = 0.5

	alpha := make([]float64, N+1)
	beta := make([]float64, N+1)

	for i := 1; i < N; i++ {
		alpha[i] = (1.0 / (c(X[i], h) - a(X[i], h)*alpha[i-1])) * b(X[i], h)
		beta[i] = (1.0 / (c(X[i], h) - a(X[i], h)*alpha[i-1])) * (f(X[i], h) + a(X[i], h)*beta[i-1])
	}

	for i := N - 1; i > 0; i-- {
		Y[i] = alpha[i]*Y[i+1] + beta[i]
	}

	for i := 0; i < N+1; i++ {
		fmt.Printf("x = %.1f ; y = %g\n", X[i], Y[i])
	}

	p := NewPlotter("Thomas Algorithm")

	if err := p.Plot(X, Y, "graph"); err != nil {
		panic(err)
	}

	if err := p.Save("graph.png"); err != nil {
		panic(err)
	}
}
