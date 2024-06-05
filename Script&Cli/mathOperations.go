package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"strconv"
)

func main() {
	// working with basic arthith metetic
	var a, b int = 16, 5
	fmt.Println("Addition= ", a+b)
	fmt.Println("Substraction= ", a-b)
	fmt.Println("Multiplication= ", a*b)
	fmt.Println("division= ", a/b)
	fmt.Println("Modulas= ", a%b)

	// working with floats
	var x, y float64 = 22.24, 32.64
	fmt.Println("Addition of Float values =", x+y)
	fmt.Println("Substraction of Float values =", x+y)
	fmt.Println("Multiplication of Float values =", x+y)
	fmt.Println("division of Float values =", x+y)
	fmt.Println("Modulas of Float values =", x+y)

	//type Conversion
	var q1 int = 52
	var q2 float64 = float64(q1)
	fmt.Println("Integer to float = ", q2)

	var w1 float64 = 53
	var w2 int = int(w1)
	fmt.Println("Float to integer= ", w2)

	// working with complex numbers

	var cmxl1 complex128 = complex(4, 8) // 4, 8 means 4 +8i
	var cmxl2 complex128 = complex(2, 4) // 2, 4 means 2+4i
	fmt.Println("complex addition = ", cmxl1+cmxl2)
	fmt.Println("complex Substraction = ", cmxl1-cmxl2)
	fmt.Println("complex multiplication  = ", cmxl1*cmxl2)
	fmt.Println("complex addition = ", cmxl1/cmxl2)
	fmt.Println("Complex Absolute Value:", cmplx.Abs(cmxl1))

	//Math Functions

	fmt.Println("Log:", math.Log(math.E)) //
	fmt.Println("Cosine", math.Cos(math.Pi))
	fmt.Println("Sin:", math.Sin(math.Pi/2))   // sin is pi/2
	fmt.Println("Square Root:", math.Sqrt(64)) // √64
	fmt.Println("Power:", math.Pow(8, 4))      // 8^4

	//string to Number Parsing
	numStr := "25.9876"
	parsedFloat, err := strconv.ParseFloat(numStr, 64) // 64 is float type
	if err != nil {
		fmt.Println("Error parsing string to float:", err)
	} else {
		fmt.Println("Parsed Float:", parsedFloat)
	}

}
