package ch06

import "strconv"

func Add(v1, v2 int) int {
	return v1 + v2
}

func FloatToString(f1, f2 float64) string {
	return strconv.FormatFloat(f1/f2*100, 'f', 2, 32) + "%"
}