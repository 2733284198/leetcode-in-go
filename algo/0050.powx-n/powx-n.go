package problem0050

import "math"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	} else if n == 1 {
		return x
	} else if n < 0 {
		return 1 / myPow(x, -n)
	} else if n%2 == 0 { // n & 1 == 0
		return myPow(x*x, n/2)
	} else {
		return myPow(x*x, n/2) * x
	}
}

func myPowV2(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / pow(x, -n)
	}

	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	m := pow(x, n>>1)
	if n&1 == 0 {

		return m * m
	}
	return m * m * x
}

func PowByBuiltInLib(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}
