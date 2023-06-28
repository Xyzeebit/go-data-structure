package main

import (
	"fmt"
	"math"
)

func main() {
	iMin8 := math.MinInt8
	iMax8 := math.MaxInt8
	iMin16 := math.MinInt16
	iMax16 := math.MaxInt16
	iMin32 := math.MinInt32
	iMax32 := math.MaxInt32
	iMin64 := math.MinInt64
	iMax64 := math.MaxInt64

	uMin8 := math.MinUint8
	uMax8 := math.MaxUint8
	uMin16 := math.MinUint16
	uMax16 := math.MaxUint16
	uMin32 := math.MinUint32
	uMax32 := math.MaxUint32
	uMin64 := math.MinUint64
	uMax64 := math.MaxUint64

	fMax32 := math.MaxFloat32
	fMax64 := math.MaxFloat64

	fmt.Println("Int8 min", iMin8, "max", iMax8)
	fmt.Println("Int16 min", iMin16, "max", iMax16)
	fmt.Println("Int32 min", iMin32, "max", iMax32)
	fmt.Println("Int64 min", iMin64, "max", iMax64)
	fmt.Println("Uint8 min", uMin8, "max", uMax8)
	fmt.Println("Uint16 min", uMin16, "max", uMax16)
	fmt.Println("Uint32 min", uMin32, "max", uMax32)
	fmt.Println("Uint64 min", uMin64, "max", uMax64)
	fmt.Println("float32", fMax32)
	fmt.Println("float64", fMax64)
}
