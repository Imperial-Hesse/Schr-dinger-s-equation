package main

import "fmt"

func conj(z complex128) complex128 {
	return complex(real(z), -imag(z))
}

func Schrödinger(ψ []complex128, t float64, H [][]complex128) []complex128{
	n := len(ψ)
	ψNew := make([]complex128, n)
	for i := 0; i < n; i++ {
		ψNew[i] = 0
		for j := 0; j < n; j++ {
			ψNew[i] += H[i][j] * ψ[j]
		}
		ψNew[i] /= complex(-imag(ψNew[i])*t, 0)
	}
	return ψNew
}
func proabilitycurrent(ψ []complex128, m float64) []float64{
	n:= len(ψ)
	j := make([]float64, n)
	for i := 0; i < n; i++ {
		j[i] = real(ψ[i] * conj(ψ[i])) / m
	}
	return j
}
func main() {
	ψ := []complex128{1,0} //Wave function
	t := 1.0 //time
	H := [][]complex128{{0,1},{1,0}}//Hamilton operator
	m := 1.0
	
	ψNew := Schrödinger(ψ, t, H)
	j := proabilitycurrent(ψNew, m)
	hold := " "
	fmt.Println(ψNew)
	fmt.Println(j)
	fmt.Scanln(&hold)
}