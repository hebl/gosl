package types

//Func1 f(x)
type Func1 func(x float64) float64

//Func2 f(x1, x2)
type Func2 func(x1, x2 float64) float64

//Func3 f(x1,x2,x3)
type Func3 func(x1, x2, x3 float64) float64

//Func f(x1,x2,x3, ... xn)
type Func func(x ...float64) float64
