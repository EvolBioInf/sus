package sus

import (
	"log"
	"math"
)

func Prob(l int, g float64, x int) float64 {
	l *= 2
	g /= 2.0
	var p1, p2 float64
	for k := 0; k < x; k++ {
		p1 += pLen(l, g, x-1, k)
	}
	for k := 0; k <= x; k++ {
		p2 += pLen(l, g, x, k)
	}
	return p2 - p1
}
func pLen(l int, g float64, x, k int) float64 {
	l1 := ln(pow(2.0, x))
	l2 := ln(Bico(x, k))
	l3 := ln(pow(g, k))
	l4 := ln(pow(0.5-g, x-k))
	l5 := ln(pow(1.0-pow(g, k)*pow(0.5-g, x-k), l))
	return exp(l1 + l2 + l3 + l4 + l5)
}
func ln(x float64) float64 {
	return math.Log(x)
}
func pow(x float64, y int) float64 {
	return math.Pow(x, float64(y))
}
func exp(x float64) float64 {
	return math.Exp(x)
}
func Bico(n, k int) float64 {
	x := factln(n) - factln(k) - factln(n-k)
	return math.Floor(0.5 + exp(x))
}

var a [101]float64

func factln(n int) float64 {
	if n < 0 {
		m := "Negative factorial in function factln."
		log.Fatal(m)
	}
	if n <= 1 {
		return 0.0
	}
	if n <= 100 {
		if a[n] > 0 {
			return a[n]
		} else {
			a[n] = gammln(float64(n + 1))
			return a[n]
		}
	}
	return gammln(float64(n + 1))
}

var cof = [6]float64{
	76.18009173,
	-86.50532033,
	24.01409822,
	-1.231739516,
	0.120858003e-2,
	-0.536382e-5}

func gammln(xx float64) float64 {
	x := xx - 1.0
	tmp := x + 5.5
	tmp -= (x + 0.5) * ln(tmp)
	ser := 1.0
	for j := 0; j <= 5; j++ {
		x += 1.0
		ser += cof[j] / x
	}
	return -tmp + ln(2.50662827465*ser)
}
func Quantile(l int, g, p float64) int {
	x := 1
	q := Prob(l, g, x)
	for q <= p {
		x++
		q += Prob(l, g, x)
	}
	return x
}
func Mean(l int, g float64) float64 {
	var cp, m float64
	for x := 1; x < l; x++ {
		p := Prob(l, g, x)
		cp += p
		m += p * float64(x)
		if cp >= 1.0-math.SmallestNonzeroFloat32 {
			break
		}
	}
	return m
}
