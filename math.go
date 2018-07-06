package sprigmath

import (
	"fmt"
	"math"

	"github.com/pkg/errors"
)

func initialArg(name string, a interface{}, iv int64, fv float64) (int64, float64, bool, error) {
	hasFloat := false
	an, err := toNumber(a)
	if err != nil {
		return 0, 0, false, errors.WithMessage(err, name+"[arg0]")
	}

	switch av := an.(type) {
	case int64:
		iv = av
	case float64:
		hasFloat = true
		fv = av
	default:
		return 0, 0, false, errors.New("inconceivable")
	}

	return iv, fv, hasFloat, err
}

//
// math currently present in sprig
//

func add1(a interface{}) (interface{}, error) {
	a, err := toNumber(a)
	if err != nil {
		return nil, errors.WithMessage(err, "add1")
	}

	switch av := a.(type) {
	case int64:
		return av + 1, nil
	case float64:
		return av + 1, nil
	}

	return nil, errors.New("inconceivable")
}

func add(a interface{}, args ...interface{}) (interface{}, error) {
	ival, fval, hasFloat, err := initialArg("add", a, 0, 0.0)
	if err != nil {
		return nil, err
	}

	for i, arg := range args {
		an, err := toNumber(arg)
		if err != nil {
			return nil, errors.WithMessage(err, fmt.Sprintf("add[arg%d]:", i+1))
		}
		switch av := an.(type) {
		case int64:
			ival += av
		case float64:
			hasFloat = true
			fval += av
		default:
			return nil, errors.New("inconceivable")
		}
	}

	if hasFloat {
		return float64(ival) + fval, nil
	}

	return ival, nil
}

func sub(a interface{}, b interface{}) (interface{}, error) {
	a, err := toNumber(a)
	if err != nil {
		return nil, errors.WithMessage(err, "sub[a]")
	}

	b, err = toNumber(b)
	if err != nil {
		return nil, errors.WithMessage(err, "sub[b]")
	}

	switch av := a.(type) {
	case int64:
		switch bv := b.(type) {
		case int64:
			return av - bv, nil
		case float64:
			return float64(av) - bv, nil
		}

	case float64:
		switch bv := b.(type) {
		case int64:
			return av - float64(bv), nil
		case float64:
			return av - bv, nil
		}
	}

	return nil, errors.New("inconceivable")
}

func div(a interface{}, b interface{}) (float64, error) {
	a, err := toNumber(a)
	if err != nil {
		return 0, errors.WithMessage(err, "div[a]")
	}

	b, err = toNumber(b)
	if err != nil {
		return 0, errors.WithMessage(err, "div[b]")
	}

	switch av := a.(type) {
	case int64:
		switch bv := b.(type) {
		case int64:
			return float64(av) / float64(bv), nil
		case float64:
			return float64(av) / bv, nil
		}

	case float64:
		switch bv := b.(type) {
		case int64:
			return av / float64(bv), nil
		case float64:
			return av / bv, nil
		}
	}

	return 0, errors.New("inconceivable")
}

func mod(a interface{}, b interface{}) (interface{}, error) {
	a, err := toNumber(a)
	if err != nil {
		return nil, errors.WithMessage(err, "mod[a]")
	}

	b, err = toNumber(b)
	if err != nil {
		return nil, errors.WithMessage(err, "mod[b]")
	}

	switch av := a.(type) {
	case int64:
		switch bv := b.(type) {
		case int64:
			return av % bv, nil
		case float64:
			return math.Mod(float64(av), bv), nil
		}

	case float64:
		switch bv := b.(type) {
		case int64:
			return math.Mod(av, float64(bv)), nil
		case float64:
			return math.Mod(av, bv), nil
		}
	}

	return nil, errors.New("inconceivable")
}

func mul(a interface{}, args ...interface{}) (interface{}, error) {
	ival, fval, hasFloat, err := initialArg("mul", a, 1, 1.0)
	if err != nil {
		return nil, err
	}

	for i, arg := range args {
		an, err := toNumber(arg)
		if err != nil {
			return nil, errors.WithMessage(err, fmt.Sprintf("mul[arg%d]:", i+1))
		}
		switch av := an.(type) {
		case int64:
			ival *= av
		case float64:
			hasFloat = true
			fval *= av
		default:
			return nil, errors.New("inconceivable")
		}
	}

	if hasFloat {
		return float64(ival) * fval, nil
	}

	return ival, nil
}

func max(a interface{}, args ...interface{}) (interface{}, error) {
	ival, fval, hasFloat, err := initialArg("max", a, math.MinInt64, -math.MaxFloat64)
	if err != nil {
		return nil, err
	}

	for i, arg := range args {
		an, err := toNumber(arg)
		if err != nil {
			return nil, errors.WithMessage(err, fmt.Sprintf("max[arg%d]:", i+1))
		}
		switch av := an.(type) {
		case int64:
			if av > ival {
				ival = av
			}
		case float64:
			hasFloat = true
			fval = math.Max(fval, av)
		default:
			return nil, errors.New("inconceivable")
		}
	}

	if hasFloat {
		return math.Max(fval, float64(ival)), nil
	}

	return ival, nil
}

func min(a interface{}, args ...interface{}) (interface{}, error) {
	ival, fval, hasFloat, err := initialArg("min", a, math.MaxInt64, math.MaxFloat64)
	if err != nil {
		return nil, err
	}

	for i, arg := range args {
		an, err := toNumber(arg)
		if err != nil {
			return nil, errors.WithMessage(err, fmt.Sprintf("min[arg%d]:", i+1))
		}
		switch av := an.(type) {
		case int64:
			if av < ival {
				ival = av
			}
		case float64:
			hasFloat = true
			fval = math.Min(fval, av)
		default:
			return nil, errors.New("inconceivable")
		}
	}

	if hasFloat {
		return math.Min(fval, float64(ival)), nil
	}

	return ival, nil
}

func ceil(arg interface{}) (int64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "ceil")
	}
	return int64(math.Ceil(val)), nil
}

func round(arg interface{}) (int64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "round")
	}
	return int64(math.Round(val)), nil
}

//
// 'complex' math
//

func acos(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "acos")
	}
	return math.Acos(val), nil
}

func acosh(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "acosh")
	}
	return math.Acosh(val), nil
}

func asin(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "asin")
	}
	return math.Asin(val), nil
}

func asinh(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "asinh")
	}
	return math.Asinh(val), nil
}

func atan(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "atan")
	}
	return math.Atan(val), nil
}

func atan2(y interface{}, x interface{}) (float64, error) {
	yv, err := toFloat64(y)
	if err != nil {
		return 0, errors.WithMessage(err, "atan2[y]")
	}

	xv, err := toFloat64(x)
	if err != nil {
		return 0, errors.WithMessage(err, "atan2[x]")
	}
	return math.Atan2(yv, xv), nil
}

func atanh(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "atanh")
	}
	return math.Atanh(val), nil
}

func cbrt(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "cbrt")
	}
	return math.Cbrt(val), nil
}

// args are inverted to accomdate `computation | copysign -1`
func copysign(x interface{}, y interface{}) (float64, error) {

	xv, err := toFloat64(x)
	if err != nil {
		return 0, errors.WithMessage(err, "copysign[x]")
	}

	yv, err := toFloat64(y)
	if err != nil {
		return 0, errors.WithMessage(err, "copysign[y]")
	}

	return math.Copysign(xv, yv), nil
}

func cos(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "cos")
	}
	return math.Cos(val), nil
}

func cosh(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "cosh")
	}
	return math.Cosh(val), nil
}

func erf(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "erf")
	}
	return math.Erf(val), nil
}

func erfc(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "erfc")
	}
	return math.Erfc(val), nil
}

func erfinv(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "erfinv")
	}
	return math.Erfinv(val), nil
}

func exp(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "exp")
	}
	return math.Exp(val), nil
}

func exp2(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "exp2")
	}
	return math.Exp2(val), nil
}

func expm1(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "expm1")
	}
	return math.Expm1(val), nil
}

func floor(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "floor")
	}
	return math.Floor(val), nil
}

func gamma(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "gammahypot")
	}
	return math.Gamma(val), nil
}

func hypot(p interface{}, q interface{}) (float64, error) {

	pv, err := toFloat64(p)
	if err != nil {
		return 0, errors.WithMessage(err, "copysign[p]")
	}

	qv, err := toFloat64(q)
	if err != nil {
		return 0, errors.WithMessage(err, "copysign[q]")
	}

	return math.Hypot(pv, qv), nil
}

func ilogb(arg interface{}) (int, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "ilogb")
	}
	return math.Ilogb(val), nil
}

func inf(arg interface{}) (float64, error) {
	val, err := toInt(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "inf")
	}
	return math.Inf(val), nil
}

func log(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "log")
	}
	return math.Log(val), nil
}

func log10(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "log10")
	}
	return math.Log10(val), nil
}

func log1p(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "log1p")
	}
	return math.Log1p(val), nil
}

func log2(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "log2")
	}
	return math.Log2(val), nil
}

func logb(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "logb")
	}
	return math.Logb(val), nil
}

func pow(x interface{}, y interface{}) (float64, error) {
	xv, err := toFloat64(x)
	if err != nil {
		return 0, errors.WithMessage(err, "copysign[x]")
	}

	yv, err := toFloat64(y)
	if err != nil {
		return 0, errors.WithMessage(err, "copysign[y]")
	}
	return math.Pow(xv, yv), nil
}

func pow10(arg interface{}) (float64, error) {
	val, err := toInt(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "signbit")
	}
	return math.Pow10(val), nil
}

func signbit(arg interface{}) (bool, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return false, errors.WithMessage(err, "signbit")
	}
	return math.Signbit(val), nil
}

func sin(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "sin")
	}
	return math.Sin(val), nil
}

func sinh(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "sinh")
	}
	return math.Sinh(val), nil
}

func sqrt(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "sqrt")
	}
	return math.Sqrt(val), nil
}

func tan(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "tan")
	}
	return math.Tan(val), nil
}

func tanh(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "tanh")
	}
	return math.Tanh(val), nil
}

func trunc(arg interface{}) (float64, error) {
	val, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "trunc")
	}
	return math.Trunc(val), nil
}

// extras

func degrees(arg interface{}) (float64, error) {
	rads, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "degrees")
	}
	return rads * (180.0 / math.Pi), nil
}

func radians(arg interface{}) (float64, error) {
	degs, err := toFloat64(arg)
	if err != nil {
		return 0, errors.WithMessage(err, "radians")
	}
	return degs * (math.Pi / 180.0), nil
}
