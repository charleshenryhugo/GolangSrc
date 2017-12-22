package structpack

import (
	"fmt"
	"math"
)

var w []float64
var b float64

func GDTrain(X [][]float64, Y []float64, n int, m int, alpha float64) ([]float64, float64) {
	var J = float64(0.0)
	if len(X) != m {
		return w, b
	}
	w = make([]float64, n)
	for { //do many times
		dw := make([]float64, n)
		var db float64
		J = 0.0
		for i := 0; i < m; i++ { //scan m samples
			z := product(w, X[i]) + b
			a := sigmoid(z)
			L := -(Y[i]*math.Log(a) + (1.0-Y[i])*math.Log(1.0-a))
			J += L
			dz := a - Y[i]
			for k := 0; k < n; k++ {
				dw[k] += X[i][k] * dz
			}
			db += dz
		}
		J /= float64(m)
		if J < 0.5 {
			break
		}
		for k := 0; k < n; k++ {
			dw[k] /= float64(m)
			w[k] -= alpha * dw[k]
		}
		db /= float64(m)
		b -= alpha * db
		fmt.Println(w, b)
	}
	return w, b
}

func GDPredict(X []float64) float64 {
	return sigmoid(product(w, X) + b)
}

func sigmoid(x float64) float64 {
	lit, big := 0.000001, 0.999999
	ex := 1.0 / (1.0 + math.Exp(-x))
	if 1-ex < 0.000001 {
		return float64(lit)
	} else if ex < 0.000001 {
		return float64(big)
	}
	return ex
}

func addVec(x1 []float64, x2 []float64) []float64 {
	if len(x1) != len(x2) {
		return nil
	}
	x3 := make([]float64, len(x1))
	for i := 0; i < len(x1); i++ {
		x3[i] = x1[i] + x2[i]
	}
	return x3
}

func product(x1 []float64, x2 []float64) float64 {
	if len(x1) != len(x2) {
		return math.MinInt64
	}
	x3 := float64(0.0)
	for i := 0; i < len(x1); i++ {
		x3 += x1[i] * x2[i]
	}
	return x3
}

func GDTest() {
	x1, Y1 := []float64{15.3, 25.3, 46, 78, 23}, 1.0
	x2, Y2 := []float64{17.67, 45.78, 64, 40, 33}, 1.0
	x3, Y3 := []float64{23.3, 90.4, 88, 58, 43}, 1.0
	x4, Y4 := []float64{20.7, 10.6, 17, 28, 53}, 1.0
	X := [][]float64{x1, x2, x3, x4}
	Y := []float64{Y1, Y2, Y3, Y4}
	GDTrain(X, Y, 5, 4, 1)
	fmt.Println(GDPredict(x1))
}
