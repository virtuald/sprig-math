package sprigmath

import (
	"github.com/Masterminds/sprig"

	"math"
	"strconv"
)

func GenericFuncMap() map[string]interface{} {
	funcMap := sprig.GenericFuncMap()

	for k, v := range functions {
		funcMap[k] = v
	}

	return funcMap
}

var functions = map[string]interface{}{
	// conversions
	"atoi":    strconv.Atoi,
	"int":     toInt,
	"int64":   toInt64,
	"float64": toFloat64,

	// converts to an integer or float
	"number": toNumber,

	// convenience
	"double": toFloat64,

	// math in sprig that we're overriding
	"add1":    add1,
	"add":     add,
	"sub":     sub,
	"div":     div,
	"mod":     mod,
	"mul":     mul,
	"biggest": max,
	"max":     max,
	"min":     min,
	"ceil":    ceil,
	"floor":   floor,
	"round":   round,

	// math
	"acos":     acos,
	"acosh":    acosh,
	"asin":     asin,
	"asinh":    asinh,
	"atan":     atan,
	"atan2":    atan2,
	"atanh":    atanh,
	"cbrt":     cbrt,
	"copysign": copysign, // args are inverted to accomdate `computation | copysign -1`
	"cos":      cos,
	"cosh":     cosh,
	"erf":      erf,
	"erfc":     erfc,
	"erfinv":   erfinv,
	"exp":      exp,
	"exp2":     exp2,
	"expm1":    expm1,
	"gamma":    gamma,
	"hypot":    hypot,
	"ilogb":    ilogb,
	"inf":      inf,
	"log":      log,
	"log10":    log10,
	"log1p":    log1p,
	"log2":     log2,
	"logb":     logb,
	"pow":      pow,
	"pow10":    pow10,
	"signbit":  signbit,
	"sin":      sin,
	"sinh":     sinh,
	"sqrt":     sqrt,
	"tan":      tan,
	"tanh":     tanh,
	"trunc":    trunc,

	// these are missing from the go math stdlib, but useful anyways?
	"degrees": degrees,
	"radians": radians,

	// constants
	"pi": func() float64 { return math.Pi },
	"e":  func() float64 { return math.E },
}
